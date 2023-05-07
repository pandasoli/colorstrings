package twobitcolor

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
      NewField(30, nil),
      errors.NO_ERROR,
    },
    {
      "valid background",
      NewField(40, nil),
      errors.NO_ERROR,
    },
    {
      "invalid color",
      NewField(50, nil),
      errors.WRONG_COLOR,
    },
    {
      "invalid props",
      NewField(30, []uint8{1}),
      errors.WRONG_PROPS,
    },
  }

  for _, tc := range cases {
    t.Run(tc.title, func(t *testing.T) {
      actual := TryTwoBitColor(tc.field)
      assert.Equal(t, tc.expected, actual.Kind)
    })
  }
}
