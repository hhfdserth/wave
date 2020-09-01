package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdateNewspaperPacket struct {
	*wave.DefaultPacket
	Variable0 int64
}

func (packet *UpdateNewspaperPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *UpdateNewspaperPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
}