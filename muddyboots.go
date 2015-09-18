package main

import (
	"flag"
	"os"

	"github.com/arussellsaw/muddyboots/client"
	"github.com/arussellsaw/muddyboots/server"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("muddyboots")
var listen = flag.String("l", ":4432", "listen address for muddyboots server")
var mode = flag.String("m", "client", "application mode: [client, server, c, s]")
var host = flag.String("h", "127.0.0.1:4432", "address of muddyboots server")

func main() {
	initLogging()

	flag.Parse()

	switch *mode {
	case "client", "c":
		initClient()
	case "server", "s":
		initServer()
	}

}

func initClient() {
	var con client.Connection
	log.Debug("initializing client")
	con.Dial("tcp", *host)

}

func initServer() {
	log.Debug("initializing server")

	var cnf = server.Config{
		*listen,
	}
	srv := server.NewServer(cnf)
	srv.Run()
}

func initLogging() {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	format := logging.MustStringFormatter(
		"%{color}%{time:15:04:05.000} >> %{level:.4s} %{color:reset} %{message}",
	)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}
