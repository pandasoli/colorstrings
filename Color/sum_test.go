package color

import (
	"testing"

	. "github.com/pandasoli/colorstrings/Field"
	. "github.com/pandasoli/colorstrings/colors/TwoBitColor"

	"github.com/stretchr/testify/assert"
)


func Test_Sum(t *testing.T) {
  assert := assert.New(t)

  color1 := Color {
    FontEffects: []Field { NewField(3, nil) },
    Foreground: NewTwoBitColor(NewField(37, nil)),
    Background: NewTwoBitColor(NewField(42, nil)),
  }

  color2 := Color {
    FontEffects: []Field { NewField(2, nil) },
    Foreground: NewTwoBitColor(NewField(31, nil)),
  }

  color1.Sum(color2)

  assert.Equal(
    Color {
      FontEffects: []Field {
        NewField(3, nil),
        NewField(2, nil),
      },
      Foreground: NewTwoBitColor(NewField(31, nil)),
      Background: NewTwoBitColor(NewField(42, nil)),
    },
    color1,
  )
}
