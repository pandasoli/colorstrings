package field

import (
	"strconv"
	"strings"
)


func MakeField(str_field string) (res Field, err error) {
  for i, str_prop := range strings.Split(str_field, ":") {
    iprop, err := strconv.Atoi(str_prop)
    if err != nil { return res, err }

    if i == 0 {
      res.Value = uint8(iprop)
    } else {
      res.Props = append(res.Props, uint8(iprop))
    }
  }

  return res, err
}
