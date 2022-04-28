package main

import (
	"fmt"
	"io/ioutil"
	"net"
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
		//if isACK(pkt, seqno) {
		//	send(pkt, conn, )
		//}
		_, err := conn.Write(pkt.dat)
		if err != nil {
			return 3
		}
	}
	// TODO: send FIN and get FINACK
	send(make_fin_pkt(seqno), conn, nil)

	// TODO: return 0 for success, 3 for failure
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
	var header = pkt.hdr
	pkt.hdr.seqno += 1 // increment the sequence number
	return (header.flag == ACK || header.flag == FINACK) && header.ackno == expected
}
