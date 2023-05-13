package colorstrings

import (
	"fmt"
	"strings"

	"github.com/pandasoli/colorstrings/Errors"
	. "github.com/pandasoli/colorstrings/Color"
	. "github.com/pandasoli/colorstrings/colors"
	. "github.com/pandasoli/colorstrings/Field"
	. "github.com/pandasoli/colorstrings/colors/TwoBitColor"
	. "github.com/pandasoli/colorstrings/colors/FourBitColor"
	. "github.com/pandasoli/colorstrings/colors/EightBitColor"
	. "github.com/pandasoli/colorstrings/colors/RGBColor"
)


func (self *ColorString) Colorize(str_fields string, position uint) error {
  if len(self.String) < int(position) {
    msg := fmt.Sprintf("Cannot colorize after the string's end (%d < %d)", len(self.String), position)
    return fmt.Errorf(msg)
  }

  var fields []Field

  // --- Making fields ---

  for _, str_field := range strings.Split(str_fields, ";") {
    field, err := MakeField(str_field)
    if err != nil { return err }

    fields = append(fields, field)
  }

  // --- Making color ---
  expr := Color {
    Position: position,
  }

  add := func(color IColor) {
    switch color.GetKind() {
      case FOREGROUND: expr.Foreground = color
      case BACKGROUND: expr.Background = color
    }
  }

  for i := 0; i < len(fields); i++ {
    field := fields[i]

    if 107 < field.Value {
      msg := fmt.Sprintf("A field non-RGB cannot be bigger than 107 (%d)", field.Value)
      return fmt.Errorf(msg)
    }

    if err := TryTwoBitColor(field); err.Kind != errors.NO_ERROR {
      if err.Kind == errors.WRONG_PROPS {
        return err.Err
      }

      // errors.WRONG_COLOR
    } else {
      add(NewTwoBitColor(field))
      continue
    }

    if err := TryFourBitColor(field); err.Kind != errors.NO_ERROR {
      if err.Kind == errors.WRONG_PROPS {
        return err.Err
      }

      // errors.WRONG_COLOR
    } else {
      add(NewFourBitColor(field))
      continue
    }

    if err := TryEightBitColor(fields[i:]); err.Kind != errors.NO_ERROR {
      if err.Kind == errors.WRONG_PROPS { return err.Err }
      if err.Kind == errors.NO_PREFIX { return err.Err }
      if err.Kind == errors.WRONG_FIELDS { return err.Err }

      // errors.WRONG_COLOR
    } else {
      add(NewEightBitColor(fields[i:]))
      i += 2
      continue
    }

    if err := TryRGBColor(fields[i:]); err.Kind != errors.NO_ERROR {
      if err.Kind == errors.WRONG_PROPS { return err.Err }
      if err.Kind == errors.NO_PREFIX { return err.Err }
      if err.Kind == errors.WRONG_FIELDS { return err.Err }

      // errors.WRONG_COLOR
    } else {
      add(NewRGBColor(fields[i:]))
      i += 4
      continue
    }

    expr.FontEffects = append(expr.FontEffects, field)
  }

  // --- Finding a nice position ---

  // replace the color at the same position's attributes
  // or append a new one if there's none at the same position
  var new_i int
  var found_new_i bool

  for i, old_color := range self.Colors {
    if old_color.Position == expr.Position {
      found_new_i = true
      self.Colors[i].Sum(expr)
    }

    if old_color.Position >= expr.Position {
      break
    }

    new_i++
  }

  if !found_new_i {
    self.Colors = append(
      self.Colors[:new_i],
      append(
        []Color { expr },
        self.Colors[new_i:]...,
      )...,
    )
  }

  return nil
}
