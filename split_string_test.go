package colorstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_SplitString(t *testing.T) {
  assert := assert.New(t)

  str, err := NewColorString(" \033[41;30mHey \033[1mhi\033[0;41;30m! \033[0m")
  assert.Nil(err)

  parts := str.SplitString(" ")

  assert.Equal("", parts[0].Join())
  assert.Equal("\033[41;30mHey", parts[1].Join())
  assert.Equal("\033[1mhi\033[0;41;30m!", parts[2].Join())
  assert.Equal("\033[0m", parts[3].Join())
}
