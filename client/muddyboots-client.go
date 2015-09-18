package client

import (
	"fmt"
	"net"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("muddyboots")

//Connection muddyboots client
type Connection struct {
	conn net.Conn
}

//Dial connect to specified server
func (c *Connection) Dial(network, host string) error {
	var err error
	c.conn, err = net.Dial(network, host)
	if err != nil {
		log.Error(fmt.Sprintf("failed to connect to server: %s", err.Error()))
		return err
	}
	log.Debug("connection established")
	return nil
}
