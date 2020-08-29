package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type CouponReadyPacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 byte
	Variable2 int32
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &CouponReadyPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *CouponReadyPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *CouponReadyPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
}
