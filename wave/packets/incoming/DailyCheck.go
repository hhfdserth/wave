package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type DailyCheckPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1
	Variable2 byte
	Variable3
}

func (packet *DailyCheckPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.Read(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable3 = b.Read(b.Bytes(), b.Index())
}

func (packet *DailyCheckPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.Write(b.Bytes(), packet.Variable1, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable2, b.Index())
	b.Write(b.Bytes(), packet.Variable3, b.Index())
}