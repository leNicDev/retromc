package packets

import "github.com/leNicDev/retromc/packet"

type SpawnPositionOutPacket struct {
	X int32
	Y int32
	Z int32
}

func (p *SpawnPositionOutPacket) Serialize() []byte {
	writer := packet.NewPacketWriter()

	writer.WriteByte(packet.SpawnPosition) // write packet id
	writer.WriteInt32(p.X)
	writer.WriteInt32(p.Y)
	writer.WriteInt32(p.Z)

	return writer.Bytes()
}
