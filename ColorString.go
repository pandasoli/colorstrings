package colorstrings

import (
  . "github.com/pandasoli/colorstrings/Color"
)


type ColorString struct {
  string_ string
  colors []Color
}

func (self ColorString) GetString() string { return self.string_ }
func (self ColorString) GetColor() []Color { return self.colors }
