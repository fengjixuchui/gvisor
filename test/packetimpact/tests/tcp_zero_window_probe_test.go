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

package tcp_zero_window_probe_test

import (
	"flag"
	"testing"
	"time"

	"golang.org/x/sys/unix"
	"gvisor.dev/gvisor/pkg/tcpip/header"
	tb "gvisor.dev/gvisor/test/packetimpact/testbench"
)

func init() {
	tb.RegisterFlags(flag.CommandLine)
}

// TestZeroWindowProbe tests few cases of zero window probing over the
// same connection.
func TestZeroWindowProbe(t *testing.T) {
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

	start := time.Now()
	// Send and receive sample data to the dut.
	dut.Send(acceptFd, sampleData, 0)
	if _, err := conn.ExpectData(&tb.TCP{}, samplePayload, time.Second); err != nil {
		t.Fatalf("expected a packet with payload %v: %s", samplePayload, err)
	}
	sendTime := time.Now().Sub(start)
	conn.Send(tb.TCP{Flags: tb.Uint8(header.TCPFlagAck | header.TCPFlagPsh)}, samplePayload)
	if _, err := conn.ExpectData(&tb.TCP{Flags: tb.Uint8(header.TCPFlagAck)}, nil, time.Second); err != nil {
		t.Fatalf("expected a packet with sequence number %s", err)
	}

	// Test 1: Check for receive of a zero window probe, record the duration for
	//         probe to be sent.
	//
	// Advertize zero window to the dut.
	conn.Send(tb.TCP{Flags: tb.Uint8(header.TCPFlagAck), WindowSize: tb.Uint16(0)})

	// Expected sequence number of the zero window probe.
	probeSeq := tb.Uint32(uint32(*conn.RemoteSeqNum() - 1))
	// Expected ack number of the ACK for the probe.
	ackProbe := tb.Uint32(uint32(*conn.RemoteSeqNum()))

	// Expect there are no zero-window probes sent until there is data to be sent out
	// from the dut.
	if _, err := conn.ExpectData(&tb.TCP{SeqNum: probeSeq}, nil, 2*time.Second); err == nil {
		t.Fatalf("unexpected a packet with sequence number %v: %s", probeSeq, err)
	}

	start = time.Now()
	// Ask the dut to send out data.
	dut.Send(acceptFd, sampleData, 0)
	// Expect zero-window probe from the dut.
	if _, err := conn.ExpectData(&tb.TCP{SeqNum: probeSeq}, nil, time.Second); err != nil {
		t.Fatalf("expected a packet with sequence number %v: %s", probeSeq, err)
	}
	// Expect the probe to be sent after some time. Compare against the previous
	// time recorded when the dut immediately sends out data on receiving the
	// send command.
	if startProbeDuration := time.Now().Sub(start); startProbeDuration <= sendTime {
		t.Fatalf("expected the first probe to be sent out after retransmission interval, got %v want > %v\n", startProbeDuration, sendTime)
	}

	// Test 2: Check if the dut recovers on advertizing non-zero receive window.
	//         and sends out the sample payload after the send window opens.
	//
	// Advertize non-zero window to the dut and ack the zero window probe.
	conn.Send(tb.TCP{AckNum: ackProbe, Flags: tb.Uint8(header.TCPFlagAck)})
	// Expect the dut to recover and transmit data.
	if _, err := conn.ExpectData(&tb.TCP{SeqNum: ackProbe}, samplePayload, time.Second); err != nil {
		t.Fatalf("expected a packet with payload %v: %s", samplePayload, err)
	}

	// Test 3: Sanity check for dut's processing of a similar probe it sent.
	//         Check if the dut responds as we do for a similar probe sent to it.
	//         Basically with sequence number to one byte behind the unacknowledged
	//         sequence number.
	p := tb.Uint32(uint32(*conn.LocalSeqNum()))
	conn.Send(tb.TCP{Flags: tb.Uint8(header.TCPFlagAck), SeqNum: tb.Uint32(uint32(*conn.LocalSeqNum() - 1))})
	if _, err := conn.ExpectData(&tb.TCP{Flags: tb.Uint8(header.TCPFlagAck), AckNum: p}, nil, time.Second); err != nil {
		t.Fatalf("expected a packet with ack number: %d: %s", p, err)
	}
}
