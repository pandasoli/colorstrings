package rgbcolor

import (
  "testing"

  . "github.com/pandasoli/colorstrings/Field"
  "github.com/pandasoli/colorstrings/Errors"

  "github.com/stretchr/testify/assert"
)


func TestTryRGBColor(t *testing.T) {
  cases := []struct {
    title string
    fields []Field
    expected errors.ErrorKind
  } {
    {
      "Valid 24-bit color",
      []Field {
        NewField(38, nil),
        NewField(2, nil),
        NewField(48, nil),
        NewField(2, nil),
        NewField(36, nil),
      },
      errors.NO_ERROR,
    },
    {
      "24-bit color with extra properties",
      []Field {
        NewField(38, nil),
        NewField(2, nil),
        NewField(48, nil),
        NewField(2, []uint8 { 0 }),
        NewField(36, nil),
      },
      errors.WRONG_PROPS,
    },
    {
      "24-bit color missing fields",
      []Field {
        NewField(38, nil),
        NewField(2, nil),
        NewField(244, nil),
        NewField(66, nil),
      },
      errors.WRONG_FIELDS,
    },
    {
      "Invalid prefix",
      []Field {
        NewField(38, nil),
        NewField(5, nil),
        NewField(244, nil),
        NewField(66, nil),
        NewField(73, nil),
      },
      errors.WRONG_PREFIX,
    },
    {
      "Invalid color value",
      []Field {
        NewField(18, nil),
        NewField(2, nil),
        NewField(244, nil),
        NewField(66, nil),
        NewField(255, nil),
      },
      errors.WRONG_COLOR,
    },
  }

  for _, tc := range cases {
    t.Run(tc.title, func(t *testing.T) {
      actual := TryRGBColor(tc.fields)
      assert.Equal(t, tc.expected, actual.Kind)
    })
  }
}
