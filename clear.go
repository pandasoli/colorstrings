package colorstrings

import (
  . "github.com/pandasoli/colorstrings/Color"
)


func (self *ColorString) Clear() {
  self.string_ = ""
  self.colors = []Color {}
}
