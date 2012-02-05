package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"player"
)

const (
	// Telnet protocol definitions.
	IAC   byte = 255
	DONT  byte = 254
	DO    byte = 253
	WONT  byte = 252
	WILL  byte = 251
	SB    byte = 250
	GA    byte = 249
	EL    byte = 248
	EC    byte = 247
	ATT   byte = 246
	AO    byte = 245
	IP    byte = 244
	BREAK byte = 243
	DK    byte = 242
	NOP   byte = 241
	SE    byte = 240
	EOR   byte = 239
	ABORT byte = 238
	SUSP  byte = 237
	EOF   byte = 236

	// Telnet options.
	BINARY byte = 0
	ECHO   byte = 1
)

type Server struct {
	Port           uint16
	acceptNextConn chan bool
	shutdown       chan bool
	charMap        map[*player.Player]bool
}

func NewServer(port uint16) *Server {
	return &Server{Port: port, acceptNextConn: make(chan bool, 1),
		shutdown: make(chan bool, 1), charMap: make(map[*player.Player]bool)}
}

func (s *Server) Shutdown() {
	s.shutdown <- true
}

func (s *Server) getAddr() string {
	return fmt.Sprintf(":%v", s.Port)
}

func sendMessage(ch *player.Player, data ...byte) error {
    slice := make([]byte, len(data))
    for _, d := range data {
        slice = append(slice, d)
    }

    return binary.Write(ch, binary.BigEndian, slice)
}

func negotiate(ch *player.Player) {
    sendMessage(ch, IAC, WILL, ECHO)
}

func (s *Server) acceptConn(l net.Listener) error {

	conn, e := l.Accept()
	Log.Std("Accepted connection.")

	if e != nil {
		return e
	}

	ch := player.NewPlayer("foo", conn)
	defer ch.Close()

	s.acceptNextConn <- true

	// XXX gonna need locking
	s.charMap[ch] = true
	fmt.Fprintln(ch, "こにちは！ Welcome to GoMUD.")

	for {
		bufr := bufio.NewReader(ch)
		buf, e := bufr.ReadString('\n')

		if e != nil {
			// XXX gonna need locking
			delete(s.charMap, ch)
			Log.Std("Connection closed: ", e)
			return e
		}

		if buf[0] == 'x' {
			s.Shutdown()
		}

		s.SendToAllConnections(buf)
		fmt.Print(buf)
	}

	return nil
}

func (s *Server) SendToAllConnections(str string) {
	for ch := range s.charMap {
		_, e := fmt.Fprint(ch, str)
		if e != nil {
			delete(s.charMap, ch)
			Log.Std("Connection closed: ", e)
		}
	}
}

func (s *Server) Run() error {

	tcpAddr := s.getAddr()

	l, e := net.Listen("tcp", tcpAddr)

	if e != nil {
		Log.Std("Unable to listen: ", e)
		return e
	}

	s.acceptNextConn <- true

	for {
		select {
		case <-s.acceptNextConn:
			go s.acceptConn(l)

		case <-s.shutdown:
			os.Exit(0)
		}
	}

	return nil
}
