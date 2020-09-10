package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

// GetID returns packet ID.
func (d *TakeCapturePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *TakeCapturePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *TakeCapturePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *TakeCapturePacket) SetSend(value bool) {
	d.Send = value
}

type TakeCapturePacket struct {
	ID                   int64
	Send                 bool
	Pos                  objects.Vector2
	NpcNum               int32
	InvestigationInvSlot byte
}

func (packet *TakeCapturePacket) Read(b buffer.PacketBuffer) {
	packet.Pos = objects.Vector2{X: b.ReadInt(b.Bytes(), b.Index()), Y: b.ReadInt(b.Bytes(), b.Index())}
	packet.NpcNum = b.ReadInt(b.Bytes(), b.Index())
	packet.InvestigationInvSlot = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *TakeCapturePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Pos.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Pos.Y, b.Index())
	b.WriteInt(b.Bytes(), packet.NpcNum, b.Index())
	b.WriteByte(b.Bytes(), packet.InvestigationInvSlot, b.Index())
}
