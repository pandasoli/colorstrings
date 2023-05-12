package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_Split(t *testing.T) {
  assert := assert.New(t)

  str := " Hey hi! "
  parts := Split(str, " ")

  assert.Equal(4, len(parts))

  assert.Equal("", parts[0].String)
  assert.Equal(uint(0), parts[0].Pos)

  assert.Equal("Hey", parts[1].String)
  assert.Equal(uint(1), parts[1].Pos)

  assert.Equal("hi!", parts[2].String)
  assert.Equal(uint(5), parts[2].Pos)

  assert.Equal("", parts[3].String)
  assert.Equal(uint(9), parts[3].Pos)
}
