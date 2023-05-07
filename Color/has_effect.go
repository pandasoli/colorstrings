package color

import . "github.com/pandasoli/colorstrings/Field"


func (self Color) HasEffect(effect Field) (has bool) {
  for _, had_effect := range self.FontEffects {
    // Test value
    if had_effect.GetValue() == effect.GetValue() {
      // Test props
      if len(had_effect.GetProps()) == len(effect.GetProps()) {
        for i, prop := range had_effect.GetProps() {
          has = prop == effect.GetProps()[i]
          if !has { break }
        }
      }
    }

    if has { break }
  }

  return has
}
