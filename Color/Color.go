package color

import (
  . "github.com/pandasoli/colorstrings/Field"
  . "github.com/pandasoli/colorstrings/colors"
)


type Color struct {
  FontEffects []Field

  Foreground,
  Background IColor

  Position uint
}
