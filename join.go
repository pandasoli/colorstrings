package colorstrings


func (self ColorString) Join() (result string) {
  result = self.string_

  for i := len(self.colors) - 1; i > -1; i-- {
    color := self.colors[i]
    code := color.Join()
    position := int(color.Position)

    result = result[:position] + code + result[position:]
  }

  return result
}
