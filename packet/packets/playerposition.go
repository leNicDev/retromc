package packets

import "github.com/leNicDev/retromc/packet"

type PlayerPositionInPacket struct {
	packet.Packet
	X        float64
	Y        float64
	Stance   float64
	Z        float64
	OnGround bool
}

func ReadPlayerPositionInPacket(data *[]byte) PlayerPositionInPacket {
	reader := packet.PacketReader{
		Data: data,
	}

	packet := PlayerPositionInPacket{}
	packet.PacketId = reader.ReadPacketId()
	packet.X = reader.ReadFloat64()
	packet.Y = reader.ReadFloat64()
	packet.Stance = reader.ReadFloat64()
	packet.Z = reader.ReadFloat64()
	packet.OnGround = reader.ReadBool()

	return packet
}
