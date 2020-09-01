package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ReceiveHourPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 int32
	Variable2 byte
	Variable3 byte
}

func (packet *ReceiveHourPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *ReceiveHourPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable2, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable3, b.Index())
}