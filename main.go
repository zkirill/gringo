package main

import (
	"flag"
	"net"

	"github.com/golang/glog"
	"github.com/zkirill/mimblewimble-go/handshake"
	"github.com/zkirill/mimblewimble-go/seeds"
)

// port is the port on which we connect to the seed.
const port = 13414

func main() {
	flag.Parse()
	seed := seeds.Seeds()[0]
	raddr := net.TCPAddr{
		IP:   net.ParseIP(seed),
		Port: port,
	}
	con, err := net.DialTCP("tcp", nil, &raddr)
	if err != nil {
		glog.Errorf("could not connect: %v", err)
		return
	}
	glog.Infof("connected")
	// Send the "hand" part of the handshake.
	h, err := handshake.NewHandshake()
	if err != nil {
		glog.Errorf("could not compose hand part of the handshake: %v", err)
		return
	}
	n, err := con.Write(h)
	if err != nil {
		glog.Errorf("could not write to connection: %v", err)
		return
	}
	glog.Infof("wrote %v bytes", n)
	// Wait for and read the second "shake" part of the handshake.
	for {
		var h handshake.Header
		if err := h.Read(con); err != nil {
			glog.Errorf("could not read header: %v", err)
			break
		}
		glog.Infof("read header with magic 1 %v, magic 2 %v, msg len %v, for message type %v", h.Magic1, h.Magic2, h.Length, h.MsgType)
		if h.MsgType == handshake.MsgTypeShake {
			var s handshake.Shake
			if err := s.Read(con); err != nil {
				glog.Errorf("could not read shake: %v", err)
				break
			}
			glog.Infof("read shake from user agent %v", s.UserAgent)
			break
		}
	}
	if err := con.Close(); err != nil {
		glog.Fatalf("could not close connection: %v", err)
	}
}
