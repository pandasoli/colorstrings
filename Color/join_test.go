package color

import (
  "testing"

  . "github.com/pandasoli/colorstrings/Field"
  . "github.com/pandasoli/colorstrings/colors/TwoBitColor"

  "github.com/stretchr/testify/assert"
)


func Test_Join(t *testing.T) {
  cases := []struct {
    title string
    color Color
    expected string
  } {
    {
      "No effects, no background, no foreground",
      Color {},
      "\033[m",
    },
    {
      "Background only",
      Color {
        Background: NewTwoBitColor(NewField(47, nil)),
      },
      "\033[47m",
    },
    {
      "Foreground only",
      Color {
        Foreground: NewTwoBitColor(NewField(31, nil)),
      },
      "\033[31m",
    },
    {
      "Background and foreground",
      Color {
        Background: NewTwoBitColor(NewField(47, nil)),
        Foreground: NewTwoBitColor(NewField(31, nil)),
      },
      "\033[47;31m",
    },
    {
      "Foreground and effects",
      Color {
        Foreground: NewTwoBitColor(NewField(31, nil)),
        FontEffects: []Field {
          NewField(1, nil), // Bold
          NewField(4, nil), // Underline
        },
      },
      "\033[1;4;31m",
    },
    {
      "All",
      Color {
        Background: NewTwoBitColor(NewField(47, nil)),
        Foreground: NewTwoBitColor(NewField(31, nil)),
        FontEffects: []Field {
          NewField(1, nil), // Bold
          NewField(4, nil), // Underline
        },
      },
      "\033[1;4;47;31m",
    },
  }

  for _, tc := range cases {
    t.Run(tc.title, func(t *testing.T) {
      actual := tc.color.Join()
      assert.Equal(t, actual, tc.expected)
    })
  }
}
