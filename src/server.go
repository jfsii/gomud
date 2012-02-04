package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"player"
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
