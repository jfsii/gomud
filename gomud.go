package main

import (
    "log"
)

const (
	PORT = 6060
)

func main() {
	log.Printf("go Mud() init")
	s := NewServer(PORT)
	log.Println("Launching new server on port: ", s.Port)
	go s.Run()
    <-s.shutdown
}
