package colorstrings

import (
  "fmt"

  . "github.com/pandasoli/colorstrings/Color"
)


func (self *ColorString) Remove(position, length uint) (removed_part ColorString, err error) {
  if int(position) > len(self.string_) {
    return removed_part, fmt.Errorf("Cannot remove after the string's end")
  }

  // Never delete after the string's end
  if int(position + length) > len(self.string_) {
    length = uint(len(self.string_)) - position
  }

  removed_part.string_ = self.string_[position:position + length]
  self.string_ = self.string_[:position] + self.string_[position + length:]

  var new_colors []Color

  for _, color := range self.colors {
    // Inside the range
    if position <= color.Position && color.Position < position + length {
      color.Position -= position
      removed_part.colors = append(removed_part.colors, color)
    } else {
      // If it's from the right side we have to decrement the removed length
      if position + length <= color.Position {
        color.Position -= length
      }

      new_colors = append(new_colors, color)
    }
  }

  self.colors = new_colors

  return removed_part, err
}
