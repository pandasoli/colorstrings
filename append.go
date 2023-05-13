package colorstrings


func (self *ColorString) Append(strs ...ColorString) {
  for _, str := range strs {
    self.AppendAt(str, uint(len(self.String)))
  }
}
