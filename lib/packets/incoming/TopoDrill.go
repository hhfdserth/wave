package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

type TopoDrillPacket struct {
	*wave.DefaultPacket
	PlayerNum int32
	Type      int32
	DrillMod  lib.DrillModRec
	MoleHP    byte
}

func (packet *TopoDrillPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Type = b.ReadInt(b.Bytes(), b.Index())
	switch packet.Type {
	case 2, 3:
		packet.DrillMod = lib.DrillModRec{Dir: byte(b.ReadInt(b.Bytes(), b.Index()))}
	case 1:
		packet.DrillMod = lib.DrillModRec{DirStar: b.ReadInt(b.Bytes(), b.Index())}
		packet.MoleHP = b.ReadByte(b.Bytes(), b.Index())
	}
}

func (packet *TopoDrillPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Type, b.Index())
	switch packet.Type {
	case 2, 3:
		b.WriteInt(b.Bytes(), packet.DrillMod.DirStar, b.Index())
	case 1:
		b.WriteInt(b.Bytes(), int32(packet.DrillMod.Dir), b.Index())
		b.WriteByte(b.Bytes(), packet.MoleHP, b.Index())
	}
}