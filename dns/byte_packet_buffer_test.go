package dns

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytePackBuffer(t *testing.T) {
	var testBuffer = NewBuffer()

	t.Run("New", func(t *testing.T) {
		var buf [512]byte
		expected := &BytePacketBuffer{
			buf: &buf,
			pos: 0,
		}
		actual := NewBuffer()

		assert.Equal(t, expected, actual)
	})

	t.Run("Position", func(t *testing.T) {
		var expected uint = 0
		actual := testBuffer.Position()
		assert.Equal(t, expected, actual)
	})

	t.Run("Step", func(t *testing.T) {
		testBuffer.Step(10)
		var expected = testBuffer.Position()
		assert.Equal(t, expected, uint(10))

		testBuffer.Step(1)
		assert.Equal(t, testBuffer.Position(), uint(11))
	})

	t.Run("Seek", func(t *testing.T) {
		testBuffer.Seek(0)
		assert.Equal(t, testBuffer.Position(), uint(0))
	})

	t.Run("Read", func(t *testing.T) {
		bb, err := testBuffer.Read()
		assert.Equal(t, nil, err)
		assert.Equal(t, byte(0), bb)

		testBuffer.Seek(BUFFER_SIZE)
		bb, err = testBuffer.Read()
		assert.NotNil(t, err)
		assert.Equal(t, byte(0), bb)
	})

	t.Run("Get", func(t *testing.T) {
		bb, err := testBuffer.Get(10)
		assert.Nil(t, err)
		assert.Equal(t, byte(0), bb)

		bb, err = testBuffer.Get(1000)
		assert.NotNil(t, err)
		assert.Equal(t, byte(0), bb)
	})
}
