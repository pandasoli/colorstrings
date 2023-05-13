package colorstrings

import (
  "fmt"

  . "github.com/pandasoli/colorstrings/Color"
)


func (self *ColorString) Remove(position, length uint) (removed_part ColorString, err error) {
  if int(position) > len(self.String) {
    return removed_part, fmt.Errorf("Cannot remove after the string's end")
  }

  // Never delete after the string's end
  if int(position + length) > len(self.String) {
    length = uint(len(self.String)) - position
  }

  removed_part.String = self.String[position:position + length]
  self.String = self.String[:position] + self.String[position + length:]

  var new_colors []Color

  for _, color := range self.Colors {
    // Inside the range
    if position <= color.Position && color.Position < position + length {
      color.Position -= position
      removed_part.Colors = append(removed_part.Colors, color)
    } else {
      // If it's from the right side we have to decrement the removed length
      if position + length <= color.Position {
        color.Position -= length
      }

      new_colors = append(new_colors, color)
    }
  }

  self.Colors = new_colors

  return removed_part, err
}
