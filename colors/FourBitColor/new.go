package fourbitcolor

import (
  . "github.com/pandasoli/colorstrings/Field"
)


func NewFourBitColor(field Field) (res FourBitColor) {
  value := field.GetValue()

  if valid_foreground(value) {
    res.ground = 9
  } else if valid_background(value) {
    res.ground = 10
  }

  res.color = value - (res.ground * 10)

  return res
}
