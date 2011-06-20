package main

import (
	"fmt"
	"net"
)

func connHandler(conn net.Conn) {
	fmt.Println("Accepting connection from: ", conn)
}

func main() {

	fmt.Println("go Mud()")

	s := NewServer(6060)

	fmt.Println("Launching new server on port: ", s.Port)

	go s.Run()

	<-s.shutdown
}
