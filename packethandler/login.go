package packethandler

import (
	"log"
	"net"

	"github.com/leNicDev/retromc/level"
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

	// create spawn position packet
	spawnPositionPacket := packets.SpawnPositionOutPacket{
		X: 0,
		Y: 0,
		Z: 0,
	}
	outData = spawnPositionPacket.Serialize()

	// write spawn position packet
	connection.Write(outData)

	// create pre chunk packet
	preChunkPacket := packets.PreChunkOutPacket{
		X:    0,
		Z:    0,
		Mode: true,
	}
	outData = preChunkPacket.Serialize()

	// write pre chunk packet
	connection.Write(outData)

	// create a new chunk
	chunk := level.NewChunk()

	// create map chunk packet
	mapChunkPacket := packets.MapChunkOutPacket{}
	mapChunkPacket.Apply(chunk)

	outData = mapChunkPacket.Serialize()

	// write map chunk packet
	connection.Write(outData)
}
