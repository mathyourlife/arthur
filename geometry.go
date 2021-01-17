package arthur

import (
	"fmt"
	"strings"
)

type Point []Expression

func NewRandPoint() *Point {
	return &Point{
		NewRandInt(),
		NewRandInt(),
		NewRandIntegerFraction(),
	}
}

func (pt Point) LaTeX() string {
	x := make([]string, 0, len(pt))
	for _, exp := range pt {
		x = append(x, exp.LaTeX())
	}
	return fmt.Sprintf("(%s)", strings.Join(x, ","))
}

type Segment struct {
	A Point
	B Point
}

func (s Segment) LaTeX() string {
  return fmt.Sprintf("\\overline{%s %s}", s.A.LaTeX(), s.B.LaTeX())
}

type Vector struct {
	A Point
	B Point
}

func (v Vector) LaTeX() string {
  return fmt.Sprintf("\\vector{%s %s}", v.A.LaTeX(), v.B.LaTeX())
}
