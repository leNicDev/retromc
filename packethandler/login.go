package packethandler

import (
	"log"
	"net"

	"github.com/leNicDev/retromc/level"
	"github.com/leNicDev/retromc/packet/packets"
	"github.com/leNicDev/retromc/player"
)

func handleLoginRequestInPacket(connection net.Conn, p packets.LoginRequestInPacket) {
	log.Printf("Login Request: %+v", p)

	sendLoginResponse(connection)
	sendPreChunk(connection)
	sendMapChunk(connection)
	sendSpawnPosition(connection)
	sendInventory(connection)
	sendPlayerPositionAndLook(connection)
}

func sendLoginResponse(connection net.Conn) {
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

func sendPreChunk(connection net.Conn) {
	// create pre chunk packet
	preChunkPacket := packets.PreChunkOutPacket{
		X:    0,
		Z:    0,
		Mode: true,
	}
	outData := preChunkPacket.Serialize()

	// write pre chunk packet
	connection.Write(outData)
}

func sendMapChunk(connection net.Conn) {
	// create a new chunk
	chunk := level.NewChunk()

	// create map chunk packet
	mapChunkPacket := packets.MapChunkOutPacket{}
	mapChunkPacket.Apply(chunk)

	outData := mapChunkPacket.Serialize()

	// write map chunk packet
	connection.Write(outData)
}

func sendSpawnPosition(connection net.Conn) {
	// create spawn position packet
	spawnPositionPacket := packets.SpawnPositionOutPacket{
		X: 0,
		Y: 0,
		Z: 0,
	}
	outData := spawnPositionPacket.Serialize()

	// write spawn position packet
	connection.Write(outData)
}

func sendInventory(connection net.Conn) {
	// create empty player inventory
	inv := player.NewInventory(player.PLAYER_INVENTORY_SIZE)

	// create new window items out packet
	windowItemsPacket := packets.WindowItemsOutPacket{
		WindowId: 0, // 0 = player inventory
		Count:    int16(inv.Size),
		Payload:  inv,
	}
	outData := windowItemsPacket.Serialize()

	// write window items out packet
	connection.Write(outData)

	// create new set slot out packet
	setSlotPacket := packets.SetSlotOutPacket{
		WindowId: 0x81,
		Slot:     -1,
		Item:     player.NewItem(-1, 1),
	}
	outData = setSlotPacket.Serialize()

	// write set slot out packet
	connection.Write(outData)
}

func sendPlayerPositionAndLook(connection net.Conn) {
	// create new player position and look out packet
	packet := packets.PlayerPositionAndLookOutPacket{
		X:        0,
		Y:        0,
		Z:        0,
		Stance:   0,
		Yaw:      0,
		Pitch:    0,
		OnGround: false,
	}
	outData := packet.Serialize()

	// write player position and look out packet
	connection.Write(outData)
}
