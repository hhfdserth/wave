package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

type NpcDirPacket struct {
	*wave.DefaultPacket
	NpcNum     int64
	Dir        byte
	IsPosition bool
	Position   objects.Vector2
}

func (n *NpcDirPacket) Read(b buffer.PacketBuffer) {
	n.NpcNum = b.ReadLong(b.Bytes(), b.Index())
	n.Dir = b.ReadByte(b.Bytes(), b.Index())
	if n.IsPosition = b.ReadBool(b.Bytes(), b.Index()); n.IsPosition {
		n.Position = objects.Vector2{
			X: b.ReadInt(b.Bytes(), b.Index()),
			Y: b.ReadInt(b.Bytes(), b.Index()),
		}
	}
}

func (n *NpcDirPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), n.NpcNum, b.Index())
	b.WriteByte(b.Bytes(), n.Dir, b.Index())
	b.WriteBool(b.Bytes(), n.IsPosition, b.Index())
	if n.IsPosition {
		b.WriteInt(b.Bytes(), n.Position.X, b.Index())
		b.WriteInt(b.Bytes(), n.Position.Y, b.Index())
	}
}