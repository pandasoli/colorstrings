package colorstrings


func (self *ColorString) AppendString(strs ...string) error {
  for _, raw_str := range strs {
    str, err := NewColorString(raw_str)
    if err != nil { return err }

    self.Append(str)
  }

  return nil
}
