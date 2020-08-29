package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type SpawnItemPacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 int32
	Variable2 byte
	Variable3 int32
	Variable4 int32
	Variable5 int32
	Variable6 buffer.Vector2
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &SpawnItemPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *SpawnItemPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadVector2(b.Bytes(), b.Index())
}

func (packet *SpawnItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable4, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable5, b.Index())
	b.WriteVector2(b.Bytes(), packet.Variable6, b.Index())
}
