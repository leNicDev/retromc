package level

type Block struct {
	TypeId   byte
	Metadata byte
	Light    byte
	SkyLight byte
}

func NewAirBlock() Block {
	return Block{
		TypeId:   0x00,
		Metadata: 0x00,
		Light:    0x00,
		SkyLight: 0x00,
	}
}
