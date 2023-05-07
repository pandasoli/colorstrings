package colorstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_Remove(t *testing.T) {
  assert := assert.New(t)

  str, err := NewColorString("\033[38;5;200mAB\033[2;40;31mC\033[0m")
  assert.Nil(err)

  err = str.Remove(2, 1)
  assert.Nil(err)

  assert.Equal(
    "\033[38;5;200mAB\033[0m",
    str.Join(),
  )
}
