package twobitcolor

import (
  "testing"

  . "github.com/pandasoli/colorstrings/Field"
  "github.com/pandasoli/colorstrings/Errors"

  "github.com/stretchr/testify/assert"
)


func Test_Join(t *testing.T) {
  cases := []struct {
    title string
    fields []Field
    expected errors.ErrorKind
  } {
    {
      "valid color",
      []Field {
        NewField(38, nil), 
        NewField(5, nil), 
        NewField(244, nil), 
      },
      errors.NO_ERROR,
    },
    {
      "wrong prefix",
      []Field {
        NewField(38, nil),
        NewField(6, nil),
      },
      errors.WRONG_PREFIX,
    },
    {
      "wrong props",
      []Field {
        NewField(38, nil),
        NewField(5, nil),
        NewField(244, []uint8{1}),
      },
      errors.WRONG_PROPS,
    },
    {
      "wrong fields",
      []Field {
        NewField(38, nil),
        NewField(5, nil),
      },
      errors.WRONG_FIELDS,
    },
    {
      "wrong color",
      []Field {
        NewField(32, nil),
        NewField(5, nil),
        NewField(244, nil),
      },
      errors.WRONG_COLOR,
    },
  }

  for _, tc := range cases {
    t.Run(tc.title, func(t *testing.T) {
      actual := TryEightBitColor(tc.fields)
      assert.Equal(t, tc.expected, actual.Kind)
    })
  }
}
