package rgbcolor

import (
  . "github.com/pandasoli/colorstrings/Field"
)


func NewRGBColor(fields []Field) (res RGBColor) {
  ground := fields[0].GetValue()
  colors := fields[2:5]

  res.ground = (ground - 8) / 10
  res.r = colors[0].GetValue()
  res.g = colors[1].GetValue()
  res.b = colors[2].GetValue()

  return res
}
