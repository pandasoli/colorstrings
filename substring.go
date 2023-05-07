package colorstrings

import "fmt"


func (self *ColorString) Substring(position, length uint) (res ColorString, err error) {
  if int(position) > len(self.string_) {
    return res, fmt.Errorf("Cannot substring after the string's end")
  } else if int(position + length) > len(self.string_) {
    return res, fmt.Errorf("Cannot substring outside the string")
  }

  res.string_ = self.string_[position:position + length]

  for _, color := range self.colors {
    // Inside the range
    if position <= color.Position && color.Position <= position + length {
      color.Position -= position

      res.colors = append(res.colors, color)
    }
  }

  return res, nil
}
