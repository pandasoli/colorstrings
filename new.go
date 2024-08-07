package colorstrings

import "regexp"


func NewColorString(str string) (result ColorString, err error) {
	re := regexp.MustCompile(`\033\[[0-9;:]*m`)
  matches := re.FindAllStringIndex(str, -1)

  result.String = re.ReplaceAllString(str, "")

  var colors_w int
  for _, pos := range matches {
    start := pos[0]
    end := pos[1]

    code := str[start + 2:end - 1] // Ignoring `\033[` and `m`
		if code == "" { code = "0" }   // Support for `\033[m`

    err = result.Colorize(code, uint(start - colors_w))
    if err != nil { break }

    colors_w += end - start
  }

  return result, err
}
