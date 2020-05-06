// Copyright 2020 The gVisor Authors.
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

package tcp_zero_window_probe_retransmit_test

import (
	"testing"
	"time"

	"golang.org/x/sys/unix"
	"gvisor.dev/gvisor/pkg/tcpip/header"
	tb "gvisor.dev/gvisor/test/packetimpact/testbench"
)

// TestZeroWindowProbeRetransmit tests retransmits of zero window probes
// to be sent at exponentially inreasing time intervals.
func TestZeroWindowProbeRetransmit(t *testing.T) {
	dut := tb.NewDUT(t)
	defer dut.TearDown()
	listenFd, remotePort := dut.CreateListener(unix.SOCK_STREAM, unix.IPPROTO_TCP, 1)
	defer dut.Close(listenFd)
	conn := tb.NewTCPIPv4(t, tb.TCP{DstPort: &remotePort}, tb.TCP{SrcPort: &remotePort})
	defer conn.Close()

	conn.Handshake()
	acceptFd, _ := dut.Accept(listenFd)
	defer dut.Close(acceptFd)

	dut.SetSockOptInt(acceptFd, unix.IPPROTO_TCP, unix.TCP_NODELAY, 1)

	sampleData := []byte("Sample Data")
	samplePayload := &tb.Payload{Bytes: sampleData}

	// Send and receive sample data to the dut.
	dut.Send(acceptFd, sampleData, 0)
	if _, err := conn.ExpectData(&tb.TCP{}, samplePayload, time.Second); err != nil {
		t.Fatalf("expected a packet with payload %v: %s", samplePayload, err)
	}
	conn.Send(tb.TCP{Flags: tb.Uint8(header.TCPFlagAck | header.TCPFlagPsh)}, samplePayload)
	if _, err := conn.ExpectData(&tb.TCP{Flags: tb.Uint8(header.TCPFlagAck)}, nil, time.Second); err != nil {
		t.Fatalf("expected a packet with sequence number %s", err)
	}

	// Check for the dut to keep the connection alive as long as the zero window
	// probes are acknowledged. Check if the zero window probes are sent at
	// exponentially increasing intervals. The timeout intervals are function
	// of the recorded first zero probe transmission duration.
	//
	// Advertize zero receive window again.
	conn.Send(tb.TCP{Flags: tb.Uint8(header.TCPFlagAck), WindowSize: tb.Uint16(0)})
	probeSeq := tb.Uint32(uint32(*conn.RemoteSeqNum() - 1))
	ackProbe := tb.Uint32(uint32(*conn.RemoteSeqNum()))

	startProbeDuration := time.Second
	current := startProbeDuration
	first := time.Now()
	// Ask the dut to send out data.
	dut.Send(acceptFd, sampleData, 0)
	// Expect the dut to keep the connection alive as long as the remote is
	// acknowledging the zero-window probes.
	for i := 0; i < 5; i++ {
		start := time.Now()
		// Expect zero-window probe with a timeout which is a function of the typical
		// first retransmission time. The retransmission times is supposed to
		// exponentially increase.
		if _, err := conn.ExpectData(&tb.TCP{SeqNum: probeSeq}, nil, 2*current); err != nil {
			t.Fatalf("expected a probe with sequence number %v: loop %d", probeSeq, i)
		}
		if i == 0 {
			startProbeDuration = time.Now().Sub(first)
			current = 2 * startProbeDuration
			continue
		}
		// Check if the probes came at exponentially increasing intervals.
		if p := time.Since(start); p < current-startProbeDuration {
			t.Fatalf("zero probe came sooner interval %d probe %d\n", p, i)
		}
		// Acknowledge the zero-window probes from the dut.
		conn.Send(tb.TCP{AckNum: ackProbe, Flags: tb.Uint8(header.TCPFlagAck), WindowSize: tb.Uint16(0)})
		current *= 2
	}
	// Advertize non-zero window.
	conn.Send(tb.TCP{AckNum: ackProbe, Flags: tb.Uint8(header.TCPFlagAck)})
	// Expect the dut to recover and transmit data.
	if _, err := conn.ExpectData(&tb.TCP{SeqNum: ackProbe}, samplePayload, time.Second); err != nil {
		t.Fatalf("expected a packet with payload %v: %s", samplePayload, err)
	}
}
