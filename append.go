package colorstrings

import (
	"fmt"
	"sort"
)


func (self *ColorString) Append(str ColorString, position uint) error {
  if int(position) > len(self.string_) {
    return fmt.Errorf("Cannot append after the string's end")
  }

  // Append the string
  self.string_ = self.string_[:position] + str.string_ + self.string_[position:]

  // Append the new colors if it has
  if len(str.colors) > 0 {
    // Update the colors' positions because now the string got longer
    for i := range str.colors {
      str.colors[i].Position += position
    }

    // Finding a nice position to append colors
    new_i := sort.Search(
      len(self.colors),
      func(i int) bool {
        position := &self.colors[i].Position
        valid := *position > str.colors[0].Position

        // Move the colors to the right after the position where the new string is being appended
        if valid { *position += uint(len(str.string_)) }

        return valid
      },
    )

    self.colors = append(
      self.colors[:new_i],
      append(str.colors, self.colors[new_i:]...)...,
    )
  }

  return nil
}
