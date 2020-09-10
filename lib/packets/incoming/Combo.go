package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (d *ComboPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ComboPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ComboPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ComboPacket) SetSend(value bool) {
	d.Send = value
}

type ComboPacket struct {
	ID          int64
	Send        bool
	ComboCount  int32
	PlayerCombo []lib.PlayerComboRec
	ComboCache  []lib.ComboCacheDataRec
}

func (packet *ComboPacket) Read(b buffer.PacketBuffer) {
	packet.ComboCount = b.ReadInt(b.Bytes(), b.Index())
	packet.PlayerCombo = make([]lib.PlayerComboRec, 18) // TODO int to const
	for i, _ := range packet.PlayerCombo {
		packet.PlayerCombo[i].Num = b.ReadInt(b.Bytes(), b.Index())
	}
	packet.ComboCache = make([]lib.ComboCacheDataRec, 3) // TODO int to const
	for i, _ := range packet.ComboCache {
		packet.ComboCache[i].Num = b.ReadInt(b.Bytes(), b.Index())
	}
}

func (packet *ComboPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ComboCount, b.Index())
	for _, combo := range packet.PlayerCombo {
		b.WriteInt(b.Bytes(), combo.Num, b.Index())
	}
	for _, cache := range packet.ComboCache {
		b.WriteInt(b.Bytes(), cache.Num, b.Index())
	}
}
