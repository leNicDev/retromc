package player

import "net"

type Player struct {
	Username   string
	EntityId   int
	Connection net.Conn
}
