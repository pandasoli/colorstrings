package colorstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_AppendAt(t *testing.T) {
  assert := assert.New(t)

  str1, err := NewColorString(" \033[41;30mHey  \033[0m")
  assert.Nil(err)

  str2, err := NewColorString("\033[1mhi\033[0;41;30m!")
  assert.Nil(err)

  err = str1.AppendAt(str2, 5)
  assert.Nil(err)

  assert.Equal(
    " \033[41;30mHey \033[1mhi\033[0;41;30m! \033[0m",
    str1.Join(),
  )
}
