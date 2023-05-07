package fourbitcolor

import (
  "testing"

  . "github.com/pandasoli/colorstrings/Field"
  "github.com/pandasoli/colorstrings/Errors"

  "github.com/stretchr/testify/assert"
)


func Test_Try(t *testing.T) {
  cases := []struct {
    title string
    field Field
    expected errors.ErrorKind
  } {
    {
      "valid foreground",
      NewField(91, nil),
      errors.NO_ERROR,
    },
    {
      "valid background",
      NewField(107, nil),
      errors.NO_ERROR,
    },
    {
      "invalid color",
      NewField(33, nil),
      errors.WRONG_COLOR,
    },
    {
      "invalid props",
      NewField(107, []uint8{2}),
      errors.WRONG_PROPS,
    },
  }

  for _, tc := range cases {
    t.Run(tc.title, func(t *testing.T) {
      actual := TryFourBitColor(tc.field)
      assert.Equal(t, tc.expected, actual.Kind)
    })
  }
}
