package main

import (
	"net"
	"os"
)

type Character struct {
	name string
	conn net.Conn
}

func NewCharacter(name string, conn net.Conn) *Character {
	return &Character{name, conn}
}

func (ch *Character) Read(p []byte) (n int, err os.Error) {
	return ch.conn.Read(p)
}

func (ch *Character) Write(p []byte) (n int, err os.Error) {
	return ch.conn.Write(p)
}

func (ch *Character) Close() {
	ch.conn.Close()
}
