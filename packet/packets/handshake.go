package packets

import "github.com/leNicDev/retromc/packet"

type HandshakeInPacket struct {
	packet.Packet
	Username string // string16
}

func ReadHandshakeInPacket(data *[]byte) HandshakeInPacket {
	reader := packet.PacketReader{
		Data: data,
	}

	packet := HandshakeInPacket{}
	packet.PacketId = reader.ReadPacketId()
	packet.Username = reader.ReadString16()

	return packet
}

type HandshakeOutPacket struct {
	ConnectionHash string // string16
}

func (p *HandshakeOutPacket) Serialize() []byte {
	w := packet.NewPacketWriter()
	w.WriteByte(packet.Handshake)     // write packet id
	w.WriteString16(p.ConnectionHash) // write connection hash (no name authentication)
	return w.Bytes()
}
