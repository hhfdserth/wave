package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

// GetID returns packet ID.
func (d *AnimationPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *AnimationPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *AnimationPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *AnimationPacket) SetSend(value bool) {
	d.Send = value
}

type AnimationPacket struct {
	ID        int64
	Send      bool
	Animation int32
	Duration  int32
	Position  objects.Vector2
	Prefix    string
	Type      byte
	Target    int32
	Sprite    int32
}

func (packet *AnimationPacket) Read(b buffer.PacketBuffer) {
	packet.Animation = b.ReadInt(b.Bytes(), b.Index())
	packet.Duration = b.ReadInt(b.Bytes(), b.Index())
	packet.Position = objects.Vector2{
		X: b.ReadInt(b.Bytes(), b.Index()),
		Y: b.ReadInt(b.Bytes(), b.Index()),
	}
	packet.Prefix = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	packet.Target = b.ReadInt(b.Bytes(), b.Index())
	packet.Sprite = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *AnimationPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Animation, b.Index())
	b.WriteInt(b.Bytes(), packet.Duration, b.Index())
	b.WriteInt(b.Bytes(), packet.Position.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Position.Y, b.Index())
	b.WriteString(b.Bytes(), packet.Prefix, b.Index())
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	b.WriteInt(b.Bytes(), packet.Target, b.Index())
	b.WriteInt(b.Bytes(), packet.Sprite, b.Index())
}
