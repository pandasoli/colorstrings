package rgbcolor

import (
  . "github.com/pandasoli/colorstrings/Field"
)


func NewRGBColor(fields []Field) (res RGBColor) {
  ground := fields[0].Value
  colors := fields[2:5]

  res.ground = (ground - 8) / 10
  res.r = colors[0].Value
  res.g = colors[1].Value
  res.b = colors[2].Value

  return res
}
