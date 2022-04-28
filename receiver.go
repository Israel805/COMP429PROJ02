package main

import (
	"net"
)

func receiver(filename *string, conn *net.UDPConn) int {
	var expected uint16 = 0
	// recieve
	for {
		// TODO: receive DATA and send ACK if exepcted seqno arrives
		// NOTE: Don't forget to write the data
		// NOTE: You'll need the addr returned from recv in order to
		// send back to the sender.
		rcv, addr, _ := recv(conn, 0)
		/*if rcv.hdr.seqno == expected {
			os.WriteFile(*filename, rcv.dat, 0666)
			_, err := send(make_ack_pkt(rcv.hdr.seqno), conn, addr)
			if err != nil {
				return 0
			}
		}
		*/
		var pkt *Packet

		if rcv.hdr.seqno == expected {
			pkt = make_ack_pkt(expected)

			_, err := conn.Write(rcv.dat) // Writing the data
			if err != nil {
				return 3
			}
			pkt.dat = rcv.dat // receive DATA from rcv
			_, err = send(pkt, conn, addr)
			if err != nil {
				return 3
			}
		}

		// TODO: break out of infinte loop after FINACK
		if rcv.hdr.flag == FINACK {
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
