package main

import "net"

type Character struct {
	name string
	conn net.Conn
}

func NewCharacter(name string, conn net.Conn) *Character {
	return &Character{name, conn}
}

func (ch *Character) Read(p []byte) (n int, err error) {
	return ch.conn.Read(p)
}

func (ch *Character) Write(p []byte) (n int, err error) {
	return ch.conn.Write(p)
}

func (ch *Character) Close() {
	ch.conn.Close()
}
