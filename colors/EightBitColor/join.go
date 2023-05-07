package twobitcolor

import "fmt"


func (self EightBitColor) Join() string {
  return fmt.Sprintf("%d8;5;%d", self.ground, self.color)
}
