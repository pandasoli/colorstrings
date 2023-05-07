package twobitcolor

import (
  "fmt"

  "github.com/pandasoli/colorstrings/Errors"
  . "github.com/pandasoli/colorstrings/Field"
)


func TryTwoBitColor(field Field) errors.Error {
  value := field.GetValue()
  props := field.GetProps()

  if !valid_foreground(value) && !valid_background(value) {
    msg := fmt.Sprintf("%d is not a valid 2-bit color", value)
    return errors.NewError(errors.WRONG_COLOR, msg)
  }

  if len(props) > 0 {
    msg := fmt.Sprintf("2-bit color expects no props (%v)", field.Join())
    return errors.NewError(errors.WRONG_PROPS, msg)
  }

  return errors.NewError(errors.NO_ERROR, "")
}
