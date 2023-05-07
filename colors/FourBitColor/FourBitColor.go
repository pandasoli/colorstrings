package fourbitcolor

import . "github.com/pandasoli/colorstrings/colors"


type FourBitColor struct {
  /*
    9 - Foreground
    10 - Background
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

func (self FourBitColor) GetKind() ColorKind {
  if self.ground == 9 {
    return FOREGROUND
  } else {
    return BACKGROUND
  }
}
