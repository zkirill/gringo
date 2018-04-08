package main

import (
	"flag"
	"io"
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
	response := make([]byte, 1024)
	// Read the second "shake" part of the handshake.
	for {
		n, err = con.Read(response)
		if err != nil {
			if err != io.EOF {
				glog.Errorf("read error:", err)
			}
			glog.Info("EOF")
			break
		}
		glog.Infof("got %v bytes", n)
	}
	if err := con.Close(); err != nil {
		glog.Fatalf("could not close connection: %v", err)
	}
}
