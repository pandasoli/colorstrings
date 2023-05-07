package twobitcolor

import (
  . "github.com/pandasoli/colorstrings/Field"
)


func NewTwoBitColor(field Field) (res TwoBitColor) {
  value := field.GetValue()

  if valid_foreground(value) {
    res.ground = 3
  } else if valid_background(value) {
    res.ground = 4
  }

  res.color = value - (res.ground * 10)

  return res
}
