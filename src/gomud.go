package main

import (
	"logger"
)

const (
	PORT = 6060
)

var Log = logger.New()

func main() {
	Log.Std("go Mud() init")
	s := NewServer(PORT)
	Log.Std("Launching new server on port: %d", s.Port)
	go s.Run()
	<-s.shutdown
}
