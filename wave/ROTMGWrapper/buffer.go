package ROTMGWrapper

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type Buffer struct {
	*buffer.DefaultBuffer
}

func InitBuffer(bufferInterface buffer.PacketBuffer) {
	buf := bufferInterface.(*Buffer)
	buf.DefaultBuffer = new(buffer.DefaultBuffer)
	buffer.InitBuffer(buf.DefaultBuffer)
	buf.SetInitialLength(8)
	buf.DefaultBuffer.Resize(8)
}

func (b *Buffer) WriteBool(data []byte, value bool, index uint64) {
	b.Next(1)
	b.PacketWriter.WriteBool(data, value, index)
}

func (b *Buffer) WriteByte(data []byte, value byte, index uint64) {
	b.Next(1)
	b.PacketWriter.WriteShort(data, int16(value), index)
}

func (b *Buffer) WriteFloat(data []byte, value float32, index uint64) {
	b.Next(4)
	b.PacketWriter.WriteFloat(data, value, index)
}

func (b *Buffer) WriteDouble(data []byte, value float64, index uint64) {
	b.Next(8)
	b.PacketWriter.WriteDouble(data, value, index)
}

func (b *Buffer) WriteInt(data []byte, value int32, index uint64) {
	b.Next(4)
	b.PacketWriter.WriteInt(data, value, index)
}

func (b *Buffer) WriteUint(data []byte, value uint32, index uint64) {
	b.Next(4)
	b.PacketWriter.WriteUint(data, value, index)
}

func (b *Buffer) WriteLong(data []byte, value int64, index uint64) {
	b.Next(8)
	b.PacketWriter.WriteLong(data, value, index)
}

func (b *Buffer) WriteULong(data []byte, value uint64, index uint64) {
	b.Next(8)
	b.PacketWriter.WriteULong(data, value, index)
}

func (b *Buffer) WriteShort(data []byte, value int16, index uint64) {
	b.Next(2)
	b.PacketWriter.WriteShort(data, value, index)
}

func (b *Buffer) WriteUShort(data []byte, value uint16, index uint64) {
	b.Next(2)
	b.PacketWriter.WriteUShort(data, value, index)
}

func (b *Buffer) WriteVector2(data []byte, value buffer.Vector2, index uint64) {
	b.Next(8)
	b.PacketWriter.WriteVector2(data, value, index)
}

func (b *Buffer) WriteString(data []byte, value string, index uint64) {
	bytes := 2 + uint64(len(value))
	b.Next(bytes)
	b.PacketWriter.WriteShort(data, int16(len(value)), index)
	b.PacketWriter.WriteString(data, value, index+2)
}

func (b *Buffer) WriteBytes(data []byte, value []byte, index uint64) {
	b.PacketWriter.WriteBytes(data, value, index)
}

func (b *Buffer) ReadBool(data []byte, index uint64) bool {
	b.Next(1)
	return b.PacketReader.ReadBool(data, index)
}

func (b *Buffer) ReadByte(data []byte, index uint64) byte {
	b.Next(1)
	return byte(b.PacketReader.ReadByte(data, index))
}

func (b *Buffer) ReadFloat(data []byte, index uint64) float32 {
	b.Next(4)
	return b.PacketReader.ReadFloat(data, index)
}

func (b *Buffer) ReadDouble(data []byte, index uint64) float64 {
	b.Next(8)
	return b.PacketReader.ReadDouble(data, index)
}

func (b *Buffer) ReadInt(data []byte, index uint64) int32 {
	b.Next(4)
	return b.PacketReader.ReadInt(data, index)
}

func (b *Buffer) ReadUint(data []byte, index uint64) uint32 {
	b.Next(4)
	return b.PacketReader.ReadUint(data, index)
}

func (b *Buffer) ReadLong(data []byte, index uint64) int64 {
	b.Next(8)
	return b.PacketReader.ReadLong(data, index)
}

func (b *Buffer) ReadULong(data []byte, index uint64) uint64 {
	b.Next(8)
	return b.PacketReader.ReadULong(data, index)
}

func (b *Buffer) ReadShort(data []byte, index uint64) int16 {
	b.Next(2)
	return b.PacketReader.ReadShort(data, index)
}

func (b *Buffer) ReadUShort(data []byte, index uint64) uint16 {
	b.Next(2)
	return b.PacketReader.ReadUShort(data, index)
}

func (b *Buffer) ReadVector2(data []byte, index uint64) buffer.Vector2 {
	b.Next(8)
	return b.PacketReader.ReadVector2(data, index)
}

func (b *Buffer) ReadString(data []byte, index uint64, length uint64) string {
	b.Next(2)
	length = uint64(b.PacketReader.ReadShort(data, index))
	b.Next(length)
	return b.PacketReader.ReadString(data, index+2, length)
}

func (b *Buffer) ReadBytes(data []byte, index uint64, length uint64) []byte {
	b.Next(length)
	return b.PacketReader.ReadBytes(data, index, length)
}

func (b *Buffer) SetReader(reader buffer.PacketReader) {
	b.PacketReader = reader
}

func (b *Buffer) SetWriter(writer buffer.PacketWriter) {
	b.PacketWriter = writer
}
