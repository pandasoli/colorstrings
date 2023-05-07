package colorstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_Connect(t *testing.T) {
  assert := assert.New(t)

  str1, err := NewColorString(" \033[90m|\033[0m          ")
  assert.Nil(err)

  str2, err := NewColorString("ab\033[31mc\033[0m")
  assert.Nil(err)

  err = str1.Connect(str2, 3)
  assert.Nil(err)

  assert.Equal(
    " \033[90m|\033[0m ab\033[31mc\033[0m      ",
    str1.Join(),
  )
}
