package colorstrings

import "fmt"


func (self *ColorString) Connect(str ColorString, position uint) error {
  if int(position) > len(self.String) {
    return fmt.Errorf("Cannot connect at a position after the string's end")
  } else if int(position) + len(str.String) > len(self.String) {
    return fmt.Errorf("Cannot connect outside the string")
  }

  left, _ := self.Substring(0, position)
  right, _ := self.Substring(position, uint(len(self.String)) - position - uint(len(str.String)))

  self.Clear()

  self.Append(left)
  self.Append(str)
  self.Append(right)

  return nil
}
