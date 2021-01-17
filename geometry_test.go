package arthur

import (
	"testing"
)

func TestPoint(t *testing.T) {

	var pt Expression
	pt = NewRandPoint()
	t.Log(pt)
	t.Log(pt.LaTeX())
	t.Fail()

}
