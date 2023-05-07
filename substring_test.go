package colorstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_Substring(t *testing.T) {
  assert := assert.New(t)

  str, err := NewColorString(" \033[41;30mHey \033[1mhi\033[0;41;30m! \033[0m")
  assert.Nil(err)

  str, err = str.Substring(5, 3)
  assert.Nil(err)

  assert.Equal(
     "\033[1mhi\033[0;41;30m!",
    str.Join(),
  )
}
