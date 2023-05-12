package colorstrings


func (self *ColorString) Append(strs ...ColorString) error {
  for _, str := range strs {
    self.AppendAt(str, uint(len(self.string_)))
  }

  return nil
}
