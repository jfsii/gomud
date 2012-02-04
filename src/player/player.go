package player

import "net"

type Player struct {
	name string
	conn net.Conn
}

func NewPlayer(name string, conn net.Conn) *Player {
	return &Player{name, conn}
}

func (ch *Player) Read(p []byte) (n int, err error) {
	return ch.conn.Read(p)
}

func (ch *Player) Write(p []byte) (n int, err error) {
	return ch.conn.Write(p)
}

func (ch *Player) Close() {
	ch.conn.Close()
}
