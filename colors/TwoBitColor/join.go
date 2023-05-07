package twobitcolor

import "fmt"


func (self TwoBitColor) Join() string {
  return fmt.Sprintf("%d%d", self.ground, self.color)
}
