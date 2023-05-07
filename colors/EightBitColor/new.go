package twobitcolor

import (
  . "github.com/pandasoli/colorstrings/Field"
)


func NewEightBitColor(fields []Field) (res EightBitColor) {
  ground := fields[0].GetValue()
  color := fields[2].GetValue()

  res.ground = (ground - 8) / 10
  res.color = color

  return res
}
