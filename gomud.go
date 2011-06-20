package main

import (
	"fmt"
	"net"
	"os"
)

type Server struct {
	Port           uint16
	acceptNextConn chan bool
	shutdown       chan bool
}

func NewServer(port uint16) *Server {
	return &Server{Port: port, acceptNextConn: make(chan bool, 1), shutdown: make(chan bool, 1)}
}

func (s *Server) getAddr() string {
	return fmt.Sprintf(":%v", s.Port)
}

func (s *Server) acceptConn(l net.Listener) os.Error {

	conn, e := l.Accept()

	if e != nil {
		return e
	}

	defer conn.Close()

	s.acceptNextConn <- true

	conn.Write([]byte("Hello World!\n"))

	buf := make([]byte, 256)
	for {

		readlen, e := conn.Read(buf)

		if e != nil {
			fmt.Println("Connection closed: ", e)
			return e
		}

		if buf[0] == 'x' {
			os.Exit(1)
		}

		// read callback
		fmt.Println(readlen, string(buf[:readlen]))
	}

	return nil
}

func (s *Server) Run() os.Error {

	tcpAddr := s.getAddr()

	l, e := net.Listen("tcp", tcpAddr)

	if e != nil {
		fmt.Println("Unable to listen: ", e)
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

func main() {

	fmt.Println("go Mud()")

	server := NewServer(6060)

	fmt.Println("Launching new server on port: ", server.Port)

	go server.Run()

	<-server.shutdown
}
