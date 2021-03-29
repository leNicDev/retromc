package packet

type Packet struct {
	PacketId byte
}

func ReadPacket(data *[]byte) Packet {
	packet := Packet{}

	packet.PacketId, _ = readPacketId(data)

	return packet
}

func readPacketId(data *[]byte) (byte, uint) {
	return (*data)[0], 1
}
