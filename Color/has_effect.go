package color

import . "github.com/pandasoli/colorstrings/Field"


func (self Color) HasEffect(effect Field) (has bool) {
  for _, had_effect := range self.FontEffects {
    // Test value
    if had_effect.Value == effect.Value {
      // Test props
      if len(had_effect.Props) == len(effect.Props) {
        for i, prop := range had_effect.Props {
          has = prop == effect.Props[i]
          if !has { break }
        }
      }
    }

    if has { break }
  }

  return has
}
