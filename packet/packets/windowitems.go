package packets

import (
	"github.com/leNicDev/retromc/packet"
	"github.com/leNicDev/retromc/player"
)

type WindowItemsOutPacket struct {
	WindowId byte
	Count    int16
	Payload  player.Inventory
}

func (p *WindowItemsOutPacket) Serialize() []byte {
	writer := packet.NewPacketWriter()

	writer.WriteByte(packet.WindowItems) // write packet id
	writer.WriteByte(p.WindowId)         // write window id
	writer.WriteInt16(p.Count)           // write inventory size (amount of slots)
	writer.Write(p.Payload.Serialize())  // write inventory data

	return writer.Bytes()
}
