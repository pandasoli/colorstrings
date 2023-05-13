package twobitcolor

import (
  "fmt"

  "github.com/pandasoli/colorstrings/Errors"
  . "github.com/pandasoli/colorstrings/Field"
)


func TryEightBitColor(fields []Field) errors.Error {
  // --- Testing ground ---
  ground := fields[0]

  if !valid_foreground(ground.Value) && !valid_background(ground.Value) {
    return errors.NewError(errors.WRONG_COLOR, "")
  }

  if 0 < len(ground.Props) {
    msg := fmt.Sprintf("%d doesn't expect props", ground.Value)
    return errors.NewError(errors.WRONG_PROPS, msg)
  }

  // --- Testing prefix ---
  if len(fields) < 2 {
    msg := fmt.Sprintf("%d expects one more field after it to know the color", ground.Value)
    return errors.NewError(errors.NO_PREFIX, msg)
  }

  prefix := fields[1]

  if prefix.Value != 5 {
    return errors.NewError(errors.WRONG_PREFIX, "")
  }

  if 0 < len(prefix.Props) {
    msg := fmt.Sprintf("8-bit color expects no props in the prefix field (%d)", len(prefix.Props))
    return errors.NewError(errors.WRONG_PROPS, msg)
  }

  // --- Testing color ---
  if len(fields) < 3 {
    msg := fmt.Sprintf("8-bit color expects the color field after %d;5", ground.Value)
    return errors.NewError(errors.WRONG_FIELDS, msg)
  }

  color := fields[2]

  if 0 < len(color.Props) {
    msg := fmt.Sprintf("8-bit color expects no props in the color field (%d)", len(color.Props))
    return errors.NewError(errors.WRONG_PROPS, msg)
  }

  return errors.NewError(errors.NO_ERROR, "")
}

