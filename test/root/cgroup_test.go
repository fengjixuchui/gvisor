// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package root

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"gvisor.dev/gvisor/runsc/cgroup"
	"gvisor.dev/gvisor/runsc/dockerutil"
	"gvisor.dev/gvisor/runsc/testutil"
)

func verifyPid(pid int, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	var gots []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		got, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		if got == pid {
			return nil
		}
		gots = append(gots, got)
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return fmt.Errorf("got: %v, want: %d", gots, pid)
}

// TestCgroup sets cgroup options and checks that cgroup was properly configured.
func TestMemCGroup(t *testing.T) {
	allocMemSize := 128 << 20
	if err := dockerutil.Pull("python"); err != nil {
		t.Fatal("docker pull failed:", err)
	}
	d := dockerutil.MakeDocker("memusage-test")

	// Start a new container and allocate the specified about of memory.
	args := []string{
		"--memory=256MB",
		"python",
		"python",
		"-c",
		fmt.Sprintf("import time; s = 'a' * %d; time.sleep(100)", allocMemSize),
	}
	if err := d.Run(args...); err != nil {
		t.Fatal("docker create failed:", err)
	}
	defer d.CleanUp()

	gid, err := d.ID()
	if err != nil {
		t.Fatalf("Docker.ID() failed: %v", err)
	}
	t.Logf("cgroup ID: %s", gid)

	path := filepath.Join("/sys/fs/cgroup/memory/docker", gid, "memory.usage_in_bytes")
	memUsage := 0

	// Wait when the container will allocate memory.
	start := time.Now()
	for time.Now().Sub(start) < 30*time.Second {
		outRaw, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read %q: %v", path, err)
		}
		out := strings.TrimSpace(string(outRaw))
		memUsage, err = strconv.Atoi(out)
		if err != nil {
			t.Fatalf("Atoi(%v): %v", out, err)
		}

		if memUsage > allocMemSize {
			return
		}

		time.Sleep(100 * time.Millisecond)
	}

	t.Fatalf("%vMB is less than %vMB", memUsage>>20, allocMemSize>>20)
}

// TestCgroup sets cgroup options and checks that cgroup was properly configured.
func TestCgroup(t *testing.T) {
	if err := dockerutil.Pull("alpine"); err != nil {
		t.Fatal("docker pull failed:", err)
	}
	d := dockerutil.MakeDocker("cgroup-test")

	// This is not a comprehensive list of attributes.
	//
	// Note that we are specifically missing cpusets, which fail if specified.
	// In any case, it's unclear if cpusets can be reliably tested here: these
	// are often run on a single core virtual machine, and there is only a single
	// CPU available in our current set, and every container's set.
	attrs := []struct {
		arg            string
		ctrl           string
		file           string
		want           string
		skipIfNotFound bool
	}{
		{
			arg:  "--cpu-shares=1000",
			ctrl: "cpu",
			file: "cpu.shares",
			want: "1000",
		},
		{
			arg:  "--cpu-period=2000",
			ctrl: "cpu",
			file: "cpu.cfs_period_us",
			want: "2000",
		},
		{
			arg:  "--cpu-quota=3000",
			ctrl: "cpu",
			file: "cpu.cfs_quota_us",
			want: "3000",
		},
		{
			arg:  "--kernel-memory=100MB",
			ctrl: "memory",
			file: "memory.kmem.limit_in_bytes",
			want: "104857600",
		},
		{
			arg:  "--memory=1GB",
			ctrl: "memory",
			file: "memory.limit_in_bytes",
			want: "1073741824",
		},
		{
			arg:  "--memory-reservation=500MB",
			ctrl: "memory",
			file: "memory.soft_limit_in_bytes",
			want: "524288000",
		},
		{
			arg:            "--memory-swap=2GB",
			ctrl:           "memory",
			file:           "memory.memsw.limit_in_bytes",
			want:           "2147483648",
			skipIfNotFound: true, // swap may be disabled on the machine.
		},
		{
			arg:  "--memory-swappiness=5",
			ctrl: "memory",
			file: "memory.swappiness",
			want: "5",
		},
		{
			arg:  "--blkio-weight=750",
			ctrl: "blkio",
			file: "blkio.weight",
			want: "750",
		},
	}

	args := make([]string, 0, len(attrs))
	for _, attr := range attrs {
		args = append(args, attr.arg)
	}

	args = append(args, "alpine", "sleep", "10000")
	if err := d.Run(args...); err != nil {
		t.Fatal("docker create failed:", err)
	}
	defer d.CleanUp()

	gid, err := d.ID()
	if err != nil {
		t.Fatalf("Docker.ID() failed: %v", err)
	}
	t.Logf("cgroup ID: %s", gid)

	// Check list of attributes defined above.
	for _, attr := range attrs {
		path := filepath.Join("/sys/fs/cgroup", attr.ctrl, "docker", gid, attr.file)
		out, err := ioutil.ReadFile(path)
		if err != nil {
			if os.IsNotExist(err) && attr.skipIfNotFound {
				t.Logf("skipped %s/%s", attr.ctrl, attr.file)
				continue
			}
			t.Fatalf("failed to read %q: %v", path, err)
		}
		if got := strings.TrimSpace(string(out)); got != attr.want {
			t.Errorf("arg: %q, cgroup attribute %s/%s, got: %q, want: %q", attr.arg, attr.ctrl, attr.file, got, attr.want)
		}
	}

	// Check that sandbox is inside cgroup.
	controllers := []string{
		"blkio",
		"cpu",
		"cpuset",
		"memory",
		"net_cls",
		"net_prio",
		"devices",
		"freezer",
		"perf_event",
		"pids",
		"systemd",
	}
	pid, err := d.SandboxPid()
	if err != nil {
		t.Fatalf("SandboxPid: %v", err)
	}
	for _, ctrl := range controllers {
		path := filepath.Join("/sys/fs/cgroup", ctrl, "docker", gid, "cgroup.procs")
		if err := verifyPid(pid, path); err != nil {
			t.Errorf("cgroup control %q processes: %v", ctrl, err)
		}
	}
}

func TestCgroupParent(t *testing.T) {
	if err := dockerutil.Pull("alpine"); err != nil {
		t.Fatal("docker pull failed:", err)
	}
	d := dockerutil.MakeDocker("cgroup-test")

	parent := testutil.RandomName("runsc")
	if err := d.Run("--cgroup-parent", parent, "alpine", "sleep", "10000"); err != nil {
		t.Fatal("docker create failed:", err)
	}
	defer d.CleanUp()
	gid, err := d.ID()
	if err != nil {
		t.Fatalf("Docker.ID() failed: %v", err)
	}
	t.Logf("cgroup ID: %s", gid)

	// Check that sandbox is inside cgroup.
	pid, err := d.SandboxPid()
	if err != nil {
		t.Fatalf("SandboxPid: %v", err)
	}

	// Finds cgroup for the sandbox's parent process to check that cgroup is
	// created in the right location relative to the parent.
	cmd := fmt.Sprintf("grep PPid: /proc/%d/status | sed 's/PPid:\\s//'", pid)
	ppid, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		t.Fatalf("Executing %q: %v", cmd, err)
	}
	cgroups, err := cgroup.LoadPaths(strings.TrimSpace(string(ppid)))
	if err != nil {
		t.Fatalf("cgroup.LoadPath(%s): %v", ppid, err)
	}
	path := filepath.Join("/sys/fs/cgroup/memory", cgroups["memory"], parent, gid, "cgroup.procs")
	if err := verifyPid(pid, path); err != nil {
		t.Errorf("cgroup control %q processes: %v", "memory", err)
	}
}
