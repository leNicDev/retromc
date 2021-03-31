package player

import "github.com/leNicDev/retromc/packet"

type Inventory struct {
	Size  uint16
	Items []Item
}

// serialize the inventory data
// see Window items (0x68): https://wiki.vg/index.php?title=Protocol&oldid=483
func (inv *Inventory) Serialize() []byte {
	itemIds := packet.NewPacketWriter()
	countAndUses := packet.NewPacketWriter()

	for i := range inv.Items {
		itemIds.WriteShort(uint16(inv.Items[i].TypeId))    // write type id
		countAndUses.WriteByte(inv.Items[i].Count)         // write item count
		countAndUses.WriteShort(uint16(inv.Items[i].Uses)) // write item uses
	}

	// concat itemIds and countAndUses
	writer := packet.NewPacketWriter()
	writer.Write(itemIds.Bytes())
	writer.Write(countAndUses.Bytes())

	return writer.Bytes()
}

func NewInventory(size uint16) Inventory {
	inv := Inventory{
		Size:  size,
		Items: make([]Item, size),
	}

	// fill inventory with empty slots
	for i := range inv.Items {
		inv.Items[i] = NewItem(-1, 1)
	}

	return inv
}
