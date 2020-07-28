package arthur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerLaTeX(t *testing.T) {
	var exp Expression
	exp = &Integer{3}
	assert.Equal(t, "3", exp.LaTeX())
	exp = &Integer{-3}
	assert.Equal(t, "-3", exp.LaTeX())
}

func TestIntegerFractionLaTeX(t *testing.T) {
	var exp Expression
	exp = &IntegerFraction{int64(3), int64(4)}
	assert.Equal(t, "\\frac{3}{4}", exp.LaTeX())
	exp = &IntegerFraction{int64(-3), int64(4)}
	assert.Equal(t, "\\frac{-3}{4}", exp.LaTeX())
	exp = &IntegerFraction{int64(-3), int64(-4)}
	assert.Equal(t, "\\frac{-3}{-4}", exp.LaTeX())
}

func TestNewRandomIntegerFraction(t *testing.T) {
	f := NewRandomIntegerFraction(IntegerFractionPropProper)
	assert.Less(t, f.Numerator().(*Integer).N, f.Denominator().(*Integer).N)
	f = NewRandomIntegerFraction(IntegerFractionPropImproper)
	assert.Less(t, f.Denominator().(*Integer).N, f.Numerator().(*Integer).N)
	f = NewRandomIntegerFraction(IntegerFractionPropUnit)
	assert.Equal(t, int64(1), f.Numerator().(*Integer).N)
	assert.NotEqual(t, int64(1), f.Denominator().(*Integer).N)
}
