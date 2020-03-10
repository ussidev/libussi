package iso

// versi baru untuk iso server

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
)

// ExecuteHandler interface
type ExecuteHandler interface {
	Execute(msg *Message) error
}

// IsoServer server handler
type IsoServer struct {
	Handler ExecuteHandler
}

// Serve server litener @ localhost:port
func (s *IsoServer) Serve(host string, port int) error {
	// create tcp listener
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	// l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return err
	}
	defer l.Close()

	// Listen for an incoming connection.
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			continue
		}

		// Handle connections in a new goroutine.
		go s.handleRequest(conn)
	}
}

func (s *IsoServer) handleRequest(conn net.Conn) {
	defer conn.Close()

	msg := &Message{}
	msg.MTI = "2100"

	// 1. parse iso message from connection
	if status, err := s.parseMessage(conn, msg); err != nil {
		msg.WriteError(conn, status, err)
		return
	}

	// 2. execute transaction
	if err := s.Handler.Execute(msg); err != nil {
		msg.WriteError(conn, msg.ResponseCode, err)
		return
	}

	// 3. send back iso to caller
	msg.Write(conn)
}

// parse message from connection into msg (msg already created)
func (s *IsoServer) parseMessage(conn net.Conn, msg *Message) (string, error) {
	if s.Handler == nil {
		return RcFail, errors.New("Handler is empty")
	}

	// get first 4 bytes as length
	lenbuf := make([]byte, 4)
	reqLen, err := conn.Read(lenbuf)
	if err != nil || reqLen != 4 {
		return RcFail, err
	}

	dataLen, err := strconv.Atoi(string(lenbuf))
	if err != nil {
		return RcFail, err
	}

	// Make a buffer to hold incoming data.
	rawIso := make([]byte, dataLen)

	// Read the incoming connection into the buffer.
	reqLen, err = conn.Read(rawIso)
	if err != nil {
		return RcFail, err
	}

	// load rawIso into UssiIso
	if err := msg.Load(rawIso, false); err != nil {
		return RcFail, err
	}

	return RcSuccess, nil
}