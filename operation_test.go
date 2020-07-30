package arthur

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	var a, b, s Expression
	a = NewRandIntegerFraction()
	b = NewRandIntegerFraction()
	s = &Sum{a, b}
	assert.Regexp(t, regexp.MustCompile(`{\\frac{[-]?[\d]+}{[-]?[\d]+}}\+{\\frac{[-]?[\d]+}{[-]?[\d]+}}`), s.LaTeX())
}
