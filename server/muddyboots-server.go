package server

import (
	"fmt"
	"net"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("muddyboots")

//Server muddyboots server
type Server struct {
	conf Config
}

//Config muddyboots server config
type Config struct {
	Listen string
}

//NewServer create new muddyboots server
func NewServer(conf Config) *Server {
	log.Debug("creating new server")
	var srv = Server{conf}
	return &srv
}

//Run execute server listener
func (s *Server) Run() {
	ln, err := net.Listen("tcp", s.conf.Listen)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to initialise server listener: %s", err.Error()))
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		log.Debug(fmt.Sprintf("init connection: %s", conn.RemoteAddr().String()))
		conn.Close()
	}
}
