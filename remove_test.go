package colorstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_Remove(t *testing.T) {
  assert := assert.New(t)

  str, err := NewColorString("\033[38;5;200mAB\033[2;40;31mC\033[0m")
  assert.Nil(err)

  removed_part, err := str.Remove(2, 1)
  assert.Nil(err)

  assert.Equal(
    "\033[2;40;31mC",
    removed_part.Join(),
  )
  assert.Equal(
    "\033[38;5;200mAB\033[0m",
    str.Join(),
  )
}
