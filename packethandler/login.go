package packethandler

import (
	"log"
	"net"

	"github.com/leNicDev/retromc/packet/packets"
)

func handleLoginRequestInPacket(connection net.Conn, p packets.LoginRequestInPacket) {
	log.Printf("Login Request: %+v", p)

	// create login response packet
	outPacket := packets.LoginResponseOutPacket{
		EntityId:  0,
		MapSeed:   0,
		Dimension: 0,
	}
	outData := outPacket.Serialize()

	// write login response packet
	connection.Write(outData)
}
