package colorstrings

import . "github.com/pandasoli/colorstrings/utils"


func (self ColorString) SplitString(at string) (res []ColorString) {
  for _, raw_part := range Split(self.string_, at) {
    part, _ := self.Substring(
      raw_part.Pos,
      uint(len(raw_part.String)),
    )

    res = append(res, part)
  }

  return res
}
