package colorstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_SplitString(t *testing.T) {
  assert := assert.New(t)

  str, err := NewColorString(" \033[41;30mHey \033[1mhi\033[0;41;30m! \033[0m")
  assert.Nil(err)

  parts, pos := str.SplitString(" ")

  assert.Equal(4, len(parts))
  assert.Equal(4, len(pos))

  assert.Equal("", parts[0].Join())
  assert.Equal(uint(0), pos[0])

  assert.Equal("\033[41;30mHey", parts[1].Join())
  assert.Equal(uint(1), pos[1])

  assert.Equal("\033[1mhi\033[0;41;30m!", parts[2].Join())
  assert.Equal(uint(5), pos[2])

  assert.Equal("\033[0m", parts[3].Join())
  assert.Equal(uint(9), pos[3])
}
