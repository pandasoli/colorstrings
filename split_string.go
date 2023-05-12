package colorstrings

import . "github.com/pandasoli/colorstrings/utils"


func (self ColorString) SplitString(at string) (res []ColorString, pos []uint) {
  for _, raw_part := range Split(self.string_, at) {
    part, _ := self.Substring(
      raw_part.Pos,
      uint(len(raw_part.String)),
    )

    pos = append(pos, raw_part.Pos)
    res = append(res, part)
  }

  return res, pos
}
