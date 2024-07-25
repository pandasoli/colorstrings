package colorstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	assert := assert.New(t)

	flat_str := " | "
	clrs_str := " \033[31m|\033[m "

	str, err := NewColorString(clrs_str)
	assert.Nil(err)

	assert.Equal(len(str.String), len(flat_str))
}
