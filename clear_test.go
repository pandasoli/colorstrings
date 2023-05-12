package colorstrings

import (
	"testing"

  . "github.com/pandasoli/colorstrings/Color"

	"github.com/stretchr/testify/assert"
)


func Test_Clear(t *testing.T) {
  assert := assert.New(t)

  str, err := NewColorString(" \033[90m|\033[0m ")
  assert.Nil(err)

  str.Clear()

  assert.Equal("", str.string_)
  assert.Equal([]Color {}, str.colors)
}
