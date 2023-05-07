package rgbcolor

import (
	"fmt"

	"github.com/pandasoli/colorstrings/Errors"
	. "github.com/pandasoli/colorstrings/Field"
)


func TryRGBColor(fields []Field) errors.Error {
  // --- Testing ground ---
  ground := fields[0]

  if !valid_foreground(ground.GetValue()) && !valid_background(ground.GetValue()) {
    return errors.NewError(errors.WRONG_COLOR, "")
  }

  if 0 < len(ground.GetProps()) {
    msg := fmt.Sprintf("%d doesn't expect props", ground.GetValue())
    return errors.NewError(errors.WRONG_PROPS, msg)
  }

  // --- Testing prefix ---
  if len(fields) < 2 {
    msg := fmt.Sprintf("%d expects one more field after it to know the color", ground.GetValue())
    return errors.NewError(errors.NO_PREFIX, msg)
  }

  prefix := fields[1]

  if prefix.GetValue() != 2 {
    return errors.NewError(errors.WRONG_PREFIX, "")
  }

  if 0 < len(prefix.GetProps()) {
    msg := fmt.Sprintf("RGB color expects no props in the prefix field (%d)", len(prefix.GetProps()))
    return errors.NewError(errors.WRONG_PROPS, msg)
  }

  // --- Testing color ---
  if len(fields) < 5 {
    msg := fmt.Sprintf("RGB color expects 3 color fields after %d;2", ground.GetValue())
    return errors.NewError(errors.WRONG_FIELDS, msg)
  }

  colors := fields[2:5]

  for i, color := range colors {
    if 0 < len(color.GetProps()) {
      field_num := map[int]string {
        0: "st",
        1: "sd",
        3: "rd",
      }

      msg := fmt.Sprintf("RGB color expects no props in the %d%s color field (%d)", i, field_num[i], len(color.GetProps()))
      return errors.NewError(errors.WRONG_PROPS, msg)
    }
  }

  return errors.NewError(errors.NO_ERROR, "")
}
