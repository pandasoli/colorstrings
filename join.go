package colorstrings


func (self ColorString) Join() (result string) {
  result = self.String

  for i := len(self.Colors) - 1; i > -1; i-- {
    color := self.Colors[i]
    code := color.Join()
    position := int(color.Position)

    result = result[:position] + code + result[position:]
  }

  return result
}
