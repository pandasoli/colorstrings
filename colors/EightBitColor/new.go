package twobitcolor

import (
  . "github.com/pandasoli/colorstrings/Field"
)


func NewEightBitColor(fields []Field) (res EightBitColor) {
  ground := fields[0].Value
  color := fields[2].Value

  res.ground = (ground - 8) / 10
  res.color = color

  return res
}
