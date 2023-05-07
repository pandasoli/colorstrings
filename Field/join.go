package field

import (
	"strconv"
	"strings"
)


func (self Field) Join() string {
  all := append([]uint8{ self.value }, self.props...)
  var str_all []string

  for _, item := range all {
    str_all = append(str_all, strconv.Itoa(int(item)))
  }

  return strings.Join(str_all, ":")
}
