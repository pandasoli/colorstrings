package fourbitcolor

import "fmt"


func (self FourBitColor) Join() string {
  return fmt.Sprintf("%d%d", self.ground, self.color)
}
