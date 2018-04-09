package main

import (
	"flag"
	"net"

	"github.com/golang/glog"
	"github.com/zkirill/mimblewimble-go/handshake"
	"github.com/zkirill/mimblewimble-go/message"
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
		var h message.Header
		if err := h.Read(con); err != nil {
			glog.Errorf("could not read header: %v", err)
			break
		}
		glog.Infof("read header with magic 1 %v, magic 2 %v, msg len %v, for message type %v", h.Magic1, h.Magic2, h.Length, h.MsgType)
		switch h.MsgType {
		case message.MsgTypePing:
			// Received ping.
			var m message.Ping
			if err := m.Read(con); err != nil {
				glog.Errorf("could not read ping: %v", err)
				break
			}
			glog.Infof("read ping with difficulty %v, height %v", m.TotalDifficulty, m.Height)
			// Send pong.
			var p message.Ping
			// Mirror the sender.
			p.Height = m.Height
			p.TotalDifficulty = m.TotalDifficulty
			if err := p.Write(true, con); err != nil {
				glog.Errorf("could not send pong: %v", err)
				break
			}
			glog.Info("sent pong")
		case message.MsgTypeShake:
			// Received shake.
			var s message.Shake
			if err := s.Read(con); err != nil {
				glog.Errorf("could not read shake: %v", err)
				break
			}
			glog.Infof("read shake from user agent %v", s.UserAgent)
		default:
			// Catch all other messages and read to the end.
			b := make([]byte, h.Length)
			n, err := con.Read(b)
			if err != nil {
				glog.Errorf("could not read message: %v", err)
				break
			}
			glog.Infof("read %v bytes", n)
		}
	}
	if err := con.Close(); err != nil {
		glog.Fatalf("could not close connection: %v", err)
	}
}
