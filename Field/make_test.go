package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_Make(t *testing.T) {
  cases := []struct {
    title string
    input string
    err bool
  } {
    {
      "Wavy font effect",
      "4:3",
      false,
    },
    {
      "Bold font effect",
      "1",
      false,
    },
    {
      "Bold with non-valid effect",
      "1;a",
      true,
    },
  }

  for _, tc := range cases {
    t.Run(tc.title, func(t *testing.T) {
      actual, err := MakeField(tc.input)

      if tc.err {
        assert.NotNil(t, err)
      } else {
        assert.Equal(t, tc.input, actual.Join())
      }
    })
  }
}
