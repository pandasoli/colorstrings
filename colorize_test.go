package colorstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_Colorize(t *testing.T) {
  assert := assert.New(t)

  var str ColorString

  cols := []string {
    "Name",
    "Age",
    "Work",
    "Country",
    "FavoriteColor",
  }

  for i, col := range cols {
    line := "\033[1m" + col + "\033[0m"
    if i < len(cols) - 1 { line += " " }

    str.AppendString(line)
  }

  assert.Equal("\033[1mName\033[0m \033[1mAge\033[0m \033[1mWork\033[0m \033[1mCountry\033[0m \033[1mFavoriteColor\033[0m", str.Join())
  assert.Equal(10, len(str.Colors))
}
