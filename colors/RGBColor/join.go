package rgbcolor

import "fmt"


func (self RGBColor) Join() string {
  return fmt.Sprintf("%d8;2;%d;%d;%d", self.ground, self.r, self.g, self.b)
}
