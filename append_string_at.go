package colorstrings


func (self *ColorString) AppendStringAt(raw_str string, position uint) error {
  str, err := NewColorString(raw_str)
  if err != nil { return err }

  return self.AppendAt(str, position)
}
