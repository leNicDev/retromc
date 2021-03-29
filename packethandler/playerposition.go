package packethandler

import (
	"log"
	"net"

	"github.com/leNicDev/retromc/packet/packets"
)

func handlePlayerPositionInPacket(connection net.Conn, p packets.PlayerPositionInPacket) {
	log.Printf("Player position: %+v", p)
}
