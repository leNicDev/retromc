package packets

import "github.com/leNicDev/retromc/packet"

type KeepAliveInPacket struct {
	packet.Packet
}

func ReadKeepAliveInPacket(data *[]byte) KeepAliveInPacket {
	reader := packet.PacketReader{
		Data: data,
	}

	packet := KeepAliveInPacket{}
	packet.PacketId = reader.ReadPacketId()

	return packet
}

type KeepAliveOutPacket struct {
}

func (p *KeepAliveOutPacket) Serialize() []byte {
	w := packet.NewPacketWriter()
	w.WriteByte(packet.KeepAlive) // write packet id
	return w.Bytes()
}
