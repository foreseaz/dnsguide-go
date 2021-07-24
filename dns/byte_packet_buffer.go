package dns

import "errors"

const BUFFER_SIZE = 512

// BytePacketBuffer defines a buffer for DNS packet.
type BytePacketBuffer struct {
	buf *[BUFFER_SIZE]byte
	pos uint
}

// NewBuffer creates a new BytePacketBuffer.
func NewBuffer() *BytePacketBuffer {
	var buffer [BUFFER_SIZE]byte
	return &BytePacketBuffer{
		buf: &buffer,
		pos: 0,
	}
}

// Position returns current position within buffer.
func (b *BytePacketBuffer) Position() uint {
	return b.pos
}

// Step step the buffer position forward a specific number of steps.
func (b *BytePacketBuffer) Step(size uint) {
	b.pos += size
}

// Seek changes buffer positions.
func (b *BytePacketBuffer) Seek(pos uint) {
	b.pos = pos
}

// Read reads a single byte and move the position one step forward.
func (b *BytePacketBuffer) Read() (byte, error) {
	if (b.pos >= BUFFER_SIZE) {
		return 0, errors.New("End of buffer")
	}

	val := b.buf[b.pos]
	b.pos += 1

	return val, nil
}

// Get reads gets a single byte on the position without change the position.
func (b *BytePacketBuffer) Get(pos uint) (byte, error) {
	if (pos >= BUFFER_SIZE) {
		return 0, errors.New("End of buffer")
	}

	val := b.buf[pos]

	return val, nil
}


