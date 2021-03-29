package packets

import "github.com/leNicDev/retromc/packet"

type PlayerPositionAndLookInPacket struct {
	packet.Packet
	X        float64
	Y        float64
	Stance   float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func ReadPlayerPositionAndLookInPacket(data *[]byte) PlayerPositionAndLookInPacket {
	reader := packet.PacketReader{
		Data: data,
	}

	packet := PlayerPositionAndLookInPacket{}
	packet.PacketId = reader.ReadPacketId()
	packet.X = reader.ReadFloat64()
	packet.Y = reader.ReadFloat64()
	packet.Stance = reader.ReadFloat64()
	packet.Z = reader.ReadFloat64()
	packet.Yaw = reader.ReadFloat32()
	packet.Pitch = reader.ReadFloat32()
	packet.OnGround = reader.ReadBool()

	return packet
}

type PlayerPositionAndLookOutPacket struct {
	X        float64
	Y        float64
	Stance   float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (p *PlayerPositionAndLookOutPacket) Serialize() []byte {
	w := packet.NewPacketWriter()
	w.WriteByte(packet.PlayerPosition) // write packet id
	w.WriteFloat64(p.X)                // write x position
	w.WriteFloat64(p.Stance)           // write stance
	w.WriteFloat64(p.Y)                // write y position
	w.WriteFloat64(p.Z)                // write z position
	w.WriteFloat32(p.Yaw)              // write yaw
	w.WriteFloat32(p.Pitch)            // write pitch
	w.WriteBool(p.OnGround)            // write on ground
	return w.Bytes()
}
