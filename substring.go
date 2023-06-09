package colorstrings

import "fmt"


func (self *ColorString) Substring(position, length uint) (res ColorString, err error) {
  if int(position) > len(self.String) {
    return res, fmt.Errorf("Cannot substring after the string's end")
  } else if int(position + length) > len(self.String) {
    return res, fmt.Errorf("Cannot substring outside the string")
  }

  res.String = self.String[position:position + length]

  for _, color := range self.Colors {
    // Inside the range
    if position <= color.Position && color.Position <= position + length {
      color.Position -= position

      res.Colors = append(res.Colors, color)
    }
  }

  return res, nil
}
