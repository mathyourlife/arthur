package arthur

import (
	"fmt"
)

type Sum struct {
	A Expression `json:"a"`
	B Expression `json:"b"`
}

func (s *Sum) LaTeX() string {
	return fmt.Sprintf("{%s}+{%s}", s.A.LaTeX(), s.B.LaTeX())
}
