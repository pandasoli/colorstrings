package utils

import "strings"


type StrPart struct {
  String string
  Pos uint
}

func Split(str, at string) (res []StrPart) {
  var start int

  for i := 0; i < len(str); i++ {
    if strings.HasPrefix(str[i:], at) {
      res = append(res, StrPart {
        String: str[start:i],
        Pos: uint(start),
      })

      start = i + len(at)
      i += len(at) - 1
    }
  }

  res = append(res, StrPart{
    String: str[start:],
    Pos: uint(start),
  })

  return res
}

