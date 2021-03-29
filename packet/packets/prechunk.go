package packets

import "github.com/leNicDev/retromc/packet"

type PreChunkOutPacket struct {
	X    int32
	Z    int32
	Mode bool // True = 1 (Initialize chunk); False = 0 (Unload chunk)
}

func (p *PreChunkOutPacket) Serialize() []byte {
	writer := packet.NewPacketWriter()

	writer.WriteByte(packet.PreChunk) // write packet id
	writer.WriteInt32(p.X)            // write chunk x position
	writer.WriteInt32(p.Z)            // write chunk z position
	writer.WriteBool(p.Mode)          // write pre chunk mode

	return writer.Bytes()
}
