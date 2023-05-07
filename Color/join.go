package color

import (
  "fmt"
  "strings"
)


func (self Color) Join() (result string) {
  var fields []string

  for _, effect := range self.FontEffects {
    fields = append(fields, effect.Join())
  }

  if self.Background != nil { fields = append(fields, self.Background.Join()) }
  if self.Foreground != nil { fields = append(fields, self.Foreground.Join()) }

  return fmt.Sprintf("\033[%sm", strings.Join(fields, ";"))
}
