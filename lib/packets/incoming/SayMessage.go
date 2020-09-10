package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *SayMessagePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *SayMessagePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *SayMessagePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *SayMessagePacket) SetSend(value bool) {
	d.Send = value
}

type SayMessagePacket struct {
	ID           int64
	Send         bool
	BubbleTarget int32
	Text         string
	Text2        string
	Color        string
	Escape       bool
	Variable6    byte
	Variable7    bool
}

func (packet *SayMessagePacket) Read(b buffer.PacketBuffer) {
	packet.BubbleTarget = b.ReadInt(b.Bytes(), b.Index())
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Text2 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Escape = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *SayMessagePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.BubbleTarget, b.Index())
	b.WriteString(b.Bytes(), packet.Text, b.Index())
	b.WriteString(b.Bytes(), packet.Text2, b.Index())
	b.WriteString(b.Bytes(), packet.Color, b.Index())
	b.WriteBool(b.Bytes(), packet.Escape, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable6, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable7, b.Index())
}
