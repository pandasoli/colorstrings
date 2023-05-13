package colorstrings

import (
	"fmt"
	"sort"
)


func (self *ColorString) AppendAt(str ColorString, position uint) error {
  if int(position) > len(self.String) {
    return fmt.Errorf("Cannot append after the string's end")
  }

  // Append the string
  self.String = self.String[:position] + str.String + self.String[position:]

  // Append the new colors if it has
  if len(str.Colors) > 0 {
    // Update the colors' positions because now the string got longer
    for i := range str.Colors {
      str.Colors[i].Position += position
    }

    // Finding a nice position to append colors
    new_i := sort.Search(
      len(self.Colors),
      func(i int) bool {
        position := &self.Colors[i].Position
        valid := *position > str.Colors[0].Position

        // Move the colors to the right after the position where the new string is being appended
        if valid { *position += uint(len(str.String)) }

        return valid
      },
    )

    self.Colors = append(
      self.Colors[:new_i],
      append(str.Colors, self.Colors[new_i:]...)...,
    )
  }

  return nil
}
