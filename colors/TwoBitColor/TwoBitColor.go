package twobitcolor

import . "github.com/pandasoli/colorstrings/colors"


type TwoBitColor struct {
  /*
    3 - Foreground
    4 - Background
  */
  ground uint8

  /*
    0 - black
    1 - red
    2 - green
    ...
    7 - white
    9 - reset
  */
  color uint8
}

func (self TwoBitColor) GetKind() ColorKind {
  if self.ground == 3 {
    return FOREGROUND
  } else {
    return BACKGROUND
  }
}
