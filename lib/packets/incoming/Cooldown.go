package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type CooldownPacket struct {
	*wave.DefaultPacket
	Num       byte
	Icon      int32
	SpellTime int32
	Variable4 bool
}

func (packet *CooldownPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadByte(b.Bytes(), b.Index())
	packet.Icon = b.ReadInt(b.Bytes(), b.Index())
	packet.SpellTime = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *CooldownPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Num, b.Index())
	b.WriteInt(b.Bytes(), packet.Icon, b.Index())
	b.WriteInt(b.Bytes(), packet.SpellTime, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable4, b.Index())
}