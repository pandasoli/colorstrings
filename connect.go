package colorstrings

import "fmt"


func (self *ColorString) Connect(str ColorString, position uint) error {
  if int(position) > len(self.string_) {
    return fmt.Errorf("Cannot connect at a position after the string's end")
  } else if int(position) + len(str.string_) > len(self.string_) {
    return fmt.Errorf("Cannot connect outside the string")
  }

  left, _ := self.Substring(0, position)
  right, _ := self.Substring(position, uint(len(self.string_)) - position - uint(len(str.string_)))

  self.Clear()

  self.Append(left)
  self.Append(str)
  self.Append(right)

  return nil
}
