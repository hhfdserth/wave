package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *EnhancementPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *EnhancementPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *EnhancementPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *EnhancementPacket) SetSend(value bool) {
	d.Send = value
}

type EnhancementPacket struct {
	ID         int64
	Send       bool
	InvSlot    byte
	TargetSlot byte
}

func (packet *EnhancementPacket) Read(b buffer.PacketBuffer) {
	packet.InvSlot = b.ReadByte(b.Bytes(), b.Index())
	packet.TargetSlot = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *EnhancementPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.InvSlot, b.Index())
	b.WriteByte(b.Bytes(), packet.TargetSlot, b.Index())
}
