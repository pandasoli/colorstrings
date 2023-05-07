package rgbcolor

import (
  . "github.com/pandasoli/colorstrings/colors"
  . "github.com/pandasoli/colorstrings/utils"
)


type RGBColor struct {
  /*
    3 - Foreground
    4 - Background
    5 - Font effect foreground
  */
  ground uint8

  /*
    0-255
  */
  r, g, b uint8
}

func (self RGBColor) GetKind() ColorKind {
  if UIntIn(uint(self.ground), []uint { 3, 5 }) {
    return FOREGROUND
  } else {
    return BACKGROUND
  }
}
