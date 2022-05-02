package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func sender(filename *string, conn *net.UDPConn) int {
	var seqno uint16 = 0
	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Println("Error, cannot read file.")
		return 2
	}

	var pkt *Packet

	for start := 0; start < len(data); start += 255 {
		end := start + 255
		if end > len(data) {
			end = len(data)
		}

		pkt = make_data_pkt(data[start:end], seqno)

		// TODO: send DATA and get ACK
		// Sends the data packet, checks for error
		_, err = send(pkt, conn, nil)
		if err != nil {
			return 3
		}

		// Gets the ACK packet, checks for error
		// returns a packet, the network, and bool
		receive, newErr, _ := recv(conn, 0)
		if newErr != nil {
			return 3
		}

		// Checks if it is an ACK type and if it is received, increment seqno
		if isACK(receive, seqno) {
			seqno++
		}

		//Write data packt to the connection
		_, err = conn.Write(pkt.dat)
		if err != nil {
			return 3
		}
	}

	// TODO: send FIN and get FINACK
	//TODO: return 0 for success, 3 for failure
	pkt = make_fin_pkt(seqno)

	// Sends fin packet, checks error
	_, err = send(pkt, conn, nil)
	if err != nil {
		return 3
	}

	// Receives the FINACK
	receive, newErr, _ := recv(conn, 0)
	if newErr != nil {
		return 3
	}

	// Checks if the received is FINACK
	if receive.hdr.flag == FINACK {
		err := os.WriteFile(*filename, receive.dat, 0622)
		if err != nil {
			return 3
		}
	}

	return 0
}

func make_data_pkt(data []byte, seqno uint16) *Packet {
	var pkt = &Packet{}

	pkt.hdr.flag = DATA
	pkt.hdr.seqno = seqno
	pkt.hdr.len = uint8(len(data))
	pkt.dat = data

	return pkt
}

func make_fin_pkt(seqno uint16) *Packet {
	var pkt = &Packet{}
	pkt.hdr.seqno, pkt.hdr.flag = seqno, FIN

	return pkt
}

func isACK(pkt *Packet, expected uint16) bool {
	// TODO: return true if ACK (including FINACK) and ackno is what is expected
	return (pkt.hdr.flag == ACK || pkt.hdr.flag == FINACK) && pkt.hdr.ackno == expected
}
