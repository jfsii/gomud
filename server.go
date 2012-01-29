package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type Server struct {
	Port           uint16
	acceptNextConn chan bool
	shutdown       chan bool
	charMap        map[*Character]*Character
}

func NewServer(port uint16) *Server {
	return &Server{Port: port, acceptNextConn: make(chan bool, 1),
		shutdown: make(chan bool, 1), charMap: make(map[*Character]*Character)}
}

func (s *Server) Shutdown() {
	s.shutdown <- true
}

func (s *Server) getAddr() string {
	return fmt.Sprintf(":%v", s.Port)
}

func (s *Server) acceptConn(l net.Listener) error {

	conn, e := l.Accept()
	log.Printf("Accepted connection.")

	if e != nil {
		return e
	}

	ch := NewCharacter("foo", conn)
	defer ch.Close()

	s.acceptNextConn <- true

	// XXX gonna need locking
	s.charMap[ch] = ch

	fmt.Fprintln(ch, "こにちは！ Welcome to GoMUD.")

	for {
		bufr := bufio.NewReader(ch)
		buf, e := bufr.ReadString('\n')

		if e != nil {
			// XXX gonna need locking
			delete(s.charMap, ch)
			log.Printf("Connection closed: ", e)
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
	for _, ch := range s.charMap {
		_, e := fmt.Fprint(ch, str)
		if e != nil {
			delete(s.charMap, ch)
			log.Printf("Connection closed: ", e)
		}
	}
}

func (s *Server) Run() error {

	tcpAddr := s.getAddr()

	l, e := net.Listen("tcp", tcpAddr)

	if e != nil {
		log.Printf("Unable to listen: ", e)
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
