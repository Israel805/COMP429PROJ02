package main

import (
	"fmt"
	"net"
	"os"
)

func receiver(filename *string, conn *net.UDPConn) int {
	var expected uint16 = 0
	// receive
	for {
		// TODO: receive DATA and send ACK if expected seqno arrives
		// NOTE: Don't forget to write the data
		// NOTE: You'll need the addr returned from recv in order to
		// send back to the sender.
		rcv, addr, _ := recv(conn, 0)

		var dataFile []byte
		if rcv.hdr.flag == DATA && rcv.hdr.seqno == expected {

			// Gets all the date from rcv till rcv.hdr.len
			dataFile, _ = os.ReadFile(*filename)
			expected++ // Increment the expected

			// Sends the ack packet with expected number
			_, err := send(make_ack_pkt(expected), conn, addr)
			if err != nil {
				fmt.Println("Error, cannot send pkt")
				return 3
			}
		}

		// TODO: break out of infinte loop after FINACK
		if rcv.hdr.flag == FIN {
			//Write to file, check for error
			err := os.WriteFile(*filename, dataFile, 0644)
			if err != nil {
				fmt.Println("Error, cannot write to file")
				return 3
			}

			//Sends a fin packet to the addres
			_, err = send(make_finack_pkt(expected), conn, addr)
			if err != nil {
				fmt.Println("Error, cannot send pkt")
				return 3
			}
			break
		}
	}

	return 0
}

func make_ack_pkt(ackno uint16) *Packet {
	var pkt *Packet = &Packet{}

	pkt.hdr.ackno = ackno
	pkt.hdr.flag = ACK

	return pkt
}

func make_finack_pkt(ackno uint16) *Packet {
	var pkt *Packet = &Packet{}

	pkt.hdr.flag = FINACK
	pkt.hdr.ackno = ackno

	return pkt
}
