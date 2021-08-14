package dns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytePackBuffer(t *testing.T) {
	var buf [BUFFER_SIZE]byte
	var emptyBuf = NewBuffer(buf)

	for i := 0; i < 512; i++ {
		buf[i] = byte(i)
	}
	var seqBuf = NewBuffer(buf)

	t.Run("New", func(t *testing.T) {
		var buf [BUFFER_SIZE]byte
		expected := &BytePacketBuffer{
			buf: &buf,
			pos: 0,
		}
		actual := NewBuffer(buf)

		assert.Equal(t, expected, actual)
	})

	t.Run("Position", func(t *testing.T) {
		var expected uint = 0
		actual := emptyBuf.Position()
		assert.Equal(t, expected, actual)
	})

	t.Run("Step", func(t *testing.T) {
		emptyBuf.Step(10)
		var expected = emptyBuf.Position()
		assert.Equal(t, expected, uint(10))

		emptyBuf.Step(1)
		assert.Equal(t, emptyBuf.Position(), uint(11))
	})

	t.Run("Seek", func(t *testing.T) {
		emptyBuf.Seek(0)
		assert.Equal(t, emptyBuf.Position(), uint(0))
	})

	t.Run("Read", func(t *testing.T) {
		bb, err := emptyBuf.Read()
		assert.Equal(t, nil, err)
		assert.Equal(t, byte(0), bb)

		emptyBuf.Seek(BUFFER_SIZE)
		bb, err = emptyBuf.Read()
		assert.NotNil(t, err)
		assert.Equal(t, byte(0), bb)
	})

	t.Run("Get", func(t *testing.T) {
		bb, err := emptyBuf.Get(10)
		assert.Nil(t, err)
		assert.Equal(t, byte(0), bb)

		bb, err = emptyBuf.Get(1000)
		assert.NotNil(t, err)
		assert.Equal(t, byte(0), bb)
	})

	t.Run("GetRange", func(t *testing.T) {
		bbs, err := seqBuf.GetRange(10, 2)
		assert.Nil(t, err)
		assert.Equal(t, []byte{10, 11}, bbs)
	})

	t.Run("Read2Bytes", func(t *testing.T) {
		bbs, err := seqBuf.Read2Bytes()
		assert.Nil(t, err)
		assert.Equal(t, []byte{0, 1}, bbs)
	})

	t.Run("Read4Bytes", func(t *testing.T) {
		bbs, err := seqBuf.Read4Bytes()
		assert.Nil(t, err)
		assert.Equal(t, []byte{2, 3, 4, 5}, bbs)
	})
}
