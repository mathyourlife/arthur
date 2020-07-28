package arthur

import (
	"fmt"
	"math/rand"
)

type IntegerProperty int

const (
	IntegerPropertyPositive IntegerProperty = 1
	IntegerPropertySmall    IntegerProperty = 2
	IntegerPropertyMedium   IntegerProperty = 2
	IntegerPropertyLarge    IntegerProperty = 2
)

type Integer struct {
	N int64
}

func NewRandInt(props ...IntegerProperty) *Integer {

	size := int64(20)
	for _, prop := range props {
		if prop == IntegerPropertyMedium {
			size = int64(1000)
		} else if prop == IntegerPropertyLarge {
			size = int64(100000)
		}
	}
	return &Integer{rand.Int63n(size) + 1}
}

func (z *Integer) LaTeX() string {
	return fmt.Sprintf("%d", z.N)
}

func (z *Integer) Equal(n *Integer) bool {
	return z.N == n.N
}

type FractionProperty int

const (
	FractionPropertyProper FractionProperty = 1
)

type Fraction interface {
	Numerator() Expression
	Denominator() Expression
}

type IntegerFractionProp int

const (
	IntegerFractionPropProper   IntegerFractionProp = 1
	IntegerFractionPropImproper IntegerFractionProp = 2
	IntegerFractionPropUnit     IntegerFractionProp = 3
)

type IntegerFraction struct {
	n int64
	d int64
}

func NewRandIntegerFraction(props ...IntegerFractionProp) *IntegerFraction {

	a := NewRandInt()
	b := NewRandInt()

	for _, prop := range props {
		if prop == IntegerFractionPropProper {
			for {
				if !a.Equal(b) {
					break
				}
				b = NewRandInt()
			}
			if a.N > b.N {
				a, b = b, a
			}
		} else if prop == IntegerFractionPropImproper {
			for {
				if !a.Equal(b) {
					break
				}
				b = NewRandInt()
			}
			if a.N < b.N {
				a, b = b, a
			}
		} else if prop == IntegerFractionPropUnit {
			a.N = 1
			for {
				if !a.Equal(b) {
					break
				}
				b = NewRandInt()
			}
		}
	}

	return &IntegerFraction{a.N, b.N}
}

func (f *IntegerFraction) Numerator() Expression {
	return &Integer{f.n}
}

func (f *IntegerFraction) Denominator() Expression {
	return &Integer{f.d}
}

func (f *IntegerFraction) LaTeX() string {
	n := &Integer{f.n}
	d := &Integer{f.d}
	return fmt.Sprintf("\\frac{%s}{%s}", n.LaTeX(), d.LaTeX())
}
