// Copyright 2019 The gVisor Authors.
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

// Binary proctor runs the test for a particular runtime. It is meant to be
// included in Docker images for all runtime tests.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"syscall"
)

// TestRunner is an interface that must be implemented for each runtime
// integrated with proctor.
type TestRunner interface {
	// ListTests returns a string slice of tests available to run.
	ListTests() ([]string, error)

	// TestCmd returns an *exec.Cmd that will run the given test.
	TestCmd(test string) *exec.Cmd
}

var (
	runtime  = flag.String("runtime", "", "name of runtime")
	list     = flag.Bool("list", false, "list all available tests")
	testName = flag.String("test", "", "run a single test from the list of available tests")
	pause    = flag.Bool("pause", false, "cause container to pause indefinitely, reaping any zombie children")
)

func main() {
	flag.Parse()

	if *pause {
		pauseAndReap()
		panic("pauseAndReap should never return")
	}

	if *runtime == "" {
		log.Fatalf("runtime flag must be provided")
	}

	tr, err := testRunnerForRuntime(*runtime)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// List tests.
	if *list {
		tests, err := tr.ListTests()
		if err != nil {
			log.Fatalf("failed to list tests: %v", err)
		}
		for _, test := range tests {
			fmt.Println(test)
		}
		return
	}

	var tests []string
	if *testName == "" {
		// Run every test.
		tests, err = tr.ListTests()
		if err != nil {
			log.Fatalf("failed to get all tests: %v", err)
		}
	} else {
		// Run a single test.
		tests = []string{*testName}
	}
	for _, test := range tests {
		cmd := tr.TestCmd(test)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("FAIL: %v", err)
		}
	}
}

// testRunnerForRuntime returns a new TestRunner for the given runtime.
func testRunnerForRuntime(runtime string) (TestRunner, error) {
	switch runtime {
	case "go":
		return goRunner{}, nil
	case "java":
		return javaRunner{}, nil
	case "nodejs":
		return nodejsRunner{}, nil
	case "php":
		return phpRunner{}, nil
	case "python":
		return pythonRunner{}, nil
	}
	return nil, fmt.Errorf("invalid runtime %q", runtime)
}

// pauseAndReap is like init. It runs forever and reaps any children.
func pauseAndReap() {
	// Get notified of any new children.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGCHLD)

	for {
		if _, ok := <-ch; !ok {
			// Channel closed. This should not happen.
			panic("signal channel closed")
		}

		// Reap the child.
		for {
			if cpid, _ := syscall.Wait4(-1, nil, syscall.WNOHANG, nil); cpid < 1 {
				break
			}
		}
	}
}

// search is a helper function to find tests in the given directory that match
// the regex.
func search(root string, testFilter *regexp.Regexp) ([]string, error) {
	var testSlice []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		name := filepath.Base(path)

		if info.IsDir() || !testFilter.MatchString(name) {
			return nil
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		testSlice = append(testSlice, relPath)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walking %q: %v", root, err)
	}

	return testSlice, nil
}
