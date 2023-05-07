package colorstrings

import (
  "fmt"
  "testing"

  "github.com/stretchr/testify/assert"
)


func Test_Join(t *testing.T) {
  assert := assert.New(t)

  str, err := NewColorString("ABC")
  assert.Nil(err)

  // --- Colorizing ---
  _2bit_color := "1;31"
  err = str.Colorize(_2bit_color, 0)
  assert.Nil(err)

  _4bit_color := "2;102;91"
  err = str.Colorize(_4bit_color, 1)
  assert.Nil(err)

  _8bit_color := "3;48;5;83;38;5;196"
  err = str.Colorize(_8bit_color, 2)
  assert.Nil(err)

  rgb_color := "4;48;2;53;232;154;38;2;161;18;22"
  err = str.Colorize(rgb_color, 3)
  assert.Nil(err)

  // --- Testing results ---
  assert.Equal(
    fmt.Sprintf(
      "\033[%smA\033[%smB\033[%smC\033[%sm",
      _2bit_color,
      _4bit_color,
      _8bit_color,
      rgb_color,
    ),
    str.Join(),
  )
}
