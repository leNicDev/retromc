package packethandler

import (
	"log"
	"net"

	"github.com/leNicDev/retromc/packet/packets"
)

func handlePlayerPositionAndLookInPacket(connection net.Conn, p packets.PlayerPositionAndLookInPacket) {
	log.Printf("Player position and look: %+v", p)
}
