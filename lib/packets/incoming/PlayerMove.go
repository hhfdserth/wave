package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerMovePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerMovePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerMovePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerMovePacket) SetSend(value bool) {
	d.Send = value
}

type PlayerMovePacket struct {
	ID        int64
	Send      bool
	Variable0 int64
}

func (packet *PlayerMovePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *PlayerMovePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
}
