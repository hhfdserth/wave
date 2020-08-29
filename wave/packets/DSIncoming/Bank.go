package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type BankPacket struct {
	*packets.DefaultPacket
	Variable1 int32
	BankSlot  byte
}

type BackSlot struct {
	Value1 int32
	Value2 int64
	Value3 byte
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &BankPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *BankPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.BankSlot = b.ReadByte(b.Bytes(), b.Index())
	if packet.BankSlot != 0 {
		// TODO Read bank data
	} else {

	}
}

func (packet *BankPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteByte(b.Bytes(), packet.BankSlot, b.Index())
}
