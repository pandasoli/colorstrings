package field

import (
  "testing"

  "github.com/stretchr/testify/assert"
)


func Test_Join(t *testing.T) {
  cases := []struct {
    title string
    field Field
    expected string
  } {
    {
      "Wavy font effect",
      NewField(4, []uint8 { 3 }),
      "4:3",
    },
    {
      "Bold font effect",
      NewField(1, nil),
      "1",
    },
  }

  for _, tc := range cases {
    t.Run(tc.title, func(t *testing.T) {
      actual := tc.field.Join()
      assert.Equal(t, tc.expected, actual)
    })
  }
}
