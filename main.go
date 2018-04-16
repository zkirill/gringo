package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/golang/glog"
	"github.com/zkirill/gringo/handshake"
	"github.com/zkirill/gringo/message"
	"github.com/zkirill/gringo/seeds"
)

// port is the port on which we connect to the seed.
const port = 13414

func main() {
	flag.Parse()
	seed := seeds.Seeds()[1]
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
			// Request peer addresses.
			// if err := RequestPeerAddrs(con); err != nil {
			// 	glog.Errorf("could not request peer addrs: %v", err)
			// 	break
			// }
			// Request block headers.
			if err := RequestBlockHeaders(con); err != nil {
				glog.Errorf("could not request block headers: %v", err)
				break
			}
		case message.MsgTypePeerAddrs:
			glog.Infof("msg peer addrs")
			var v message.PeerAddrs
			if err := v.Read(con); err != nil {
				glog.Errorf("could not read peers: %v", err)
				break
			}
			glog.Infof("read %v peer addrs", len(v.Peers))
			if len(v.Peers) > 0 {
				glog.Infof("first peer: %v", v.Peers[0])
			}
		case message.MsgTypeHeaders:
			glog.Infof("msg headers")
			var v message.BlockHeaders
			if err := v.Read(con); err != nil {
				glog.Errorf("could not read headers: %v", err)
				break
			}
			glog.Infof("read %v headers", len(v.Headers))
			if len(v.Headers) > 0 {
				glog.Infof("first header difficulty: %v", v.Headers[0].TotalDifficulty)
				glog.Infof("first header nonce: %v", v.Headers[0].Nonce)
				glog.Infof("first header pow: %v", v.Headers[0].ProofOfWork)
			}
		case message.MsgTypeBlock:
			glog.Infof("msg block")
			var v message.Block
			if err := v.Read(con); err != nil {
				glog.Errorf("could not read headers: %v", err)
				break
			}
			glog.Infof("read %v headers", len(v.Headers))
			if len(v.Headers) > 0 {
				glog.Infof("first header difficulty: %v", v.Headers[0].TotalDifficulty)
				glog.Infof("first header nonce: %v", v.Headers[0].Nonce)
				glog.Infof("first header pow: %v", v.Headers[0].ProofOfWork)
			}
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

// RequestPeerAddrs requests peer addresses.
func RequestPeerAddrs(con *net.TCPConn) error {
	var r message.GetPeerAddrs
	err := r.Write(con)
	if err != nil {
		return fmt.Errorf("could not write to connection: %v", err)
	}
	glog.Info("requested peer addresses")
	return nil
}

// RequestBlockHeaders requests block headers.
func RequestBlockHeaders(con *net.TCPConn) error {
	var r message.GetHeaders
	err := r.Write(con)
	if err != nil {
		return fmt.Errorf("could not write to connection: %v", err)
	}
	glog.Info("requested headers")
	return nil
}

// RequestBlock requests block headers.
func RequestBlock(hash message.Hash, con *net.TCPConn) error {
	if err := message.GetBlock(hash, con); err != nil {
		return fmt.Errorf("could not write to connection: %v", err)
	}
	glog.Info("requested block")
	return nil
}
