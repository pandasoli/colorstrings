package twobitcolor

import (
  . "github.com/pandasoli/colorstrings/colors"
  . "github.com/pandasoli/colorstrings/utils"
)


type EightBitColor struct {
  /*
    3 - Foreground
    4 - Background
    5 - Font effect foreground
  */
  ground uint8

  /*
    0-255
  */
  color uint8
}

func (self EightBitColor) GetKind() ColorKind {
  if UIntIn(uint(self.ground), []uint { 3, 5 }) {
    return FOREGROUND
  } else {
    return BACKGROUND
  }
}
