package color


func (self *Color) Sum(color Color) {
  for _, effect := range color.FontEffects {
    if !self.HasEffect(effect) {
      self.FontEffects = append(self.FontEffects, effect)
    }
  }

  if color.Background != nil { self.Background = color.Background }
  if color.Foreground != nil { self.Foreground = color.Foreground }
}
