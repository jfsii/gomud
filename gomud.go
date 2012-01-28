package main

import (
	"fmt"
)

const (
	PORT = 6060
)

func main() {
	fmt.Println("go Mud()")
	s := NewServer(PORT)
	fmt.Println("Launching new server on port: ", s.Port)
	go s.Run()
	<-s.shutdown
}
