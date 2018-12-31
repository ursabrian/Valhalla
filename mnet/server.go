package mnet

import (
	"net"

	"github.com/Hucaru/Valhalla/constant"
	"github.com/Hucaru/Valhalla/mpacket"
)

type MConnServer interface {
	MConn
}

type server struct {
	baseConn
}

func NewServer(conn net.Conn, eRecv chan *Event, queueSize int) *server {
	s := &server{}
	s.Conn = conn

	s.eSend = make(chan mpacket.Packet, queueSize)
	s.eRecv = eRecv

	s.reader = func() {
		serverReader(s, s.eRecv, constant.ClientHeaderSize)
	}

	return s
}
