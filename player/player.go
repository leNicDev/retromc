package player

import "net"

const (
	PLAYER_INVENTORY_SIZE = 45
)

type Player struct {
	Username   string
	EntityId   int
	Connection net.Conn
	Inventory  Inventory
}
