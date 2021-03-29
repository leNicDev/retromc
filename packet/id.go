package packet

const (
	KeepAlive             byte = 0x00
	LoginRequest          byte = 0x01
	Handshake             byte = 0x02
	ChatMessage           byte = 0x03
	TimeUpdate            byte = 0x04
	EntityEquipment       byte = 0x05
	SpawnPosition         byte = 0x06
	PlayerPosition        byte = 0x0b
	PlayerPositionAndLook byte = 0x0d
	PreChunk              byte = 0x32
	MapChunk              byte = 0x33
)
