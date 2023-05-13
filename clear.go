package colorstrings

import (
  . "github.com/pandasoli/colorstrings/Color"
)


func (self *ColorString) Clear() {
  self.String = ""
  self.Colors = []Color {}
}
