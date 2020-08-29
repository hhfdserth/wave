package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type CharacterConfirmPacket struct {
	*packets.DefaultPacket
	Variable0 int64
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &CharacterConfirmPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *CharacterConfirmPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *CharacterConfirmPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
}
