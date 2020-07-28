package tib

import (
	"fmt"
	"net/http"

	"github.com/mathyourlife/arthur"
)

func ArthurHandler(w http.ResponseWriter, r *http.Request) {
	fractionType := r.URL.Query().Get("type")
	var exp arthur.Expression
	if fractionType == "improper" {
		exp = arthur.NewRandomIntegerFraction(arthur.IntegerFractionPropImproper)
	} else if fractionType == "unit" {
		exp = arthur.NewRandomIntegerFraction(arthur.IntegerFractionPropUnit)
	} else {
		exp = arthur.NewRandomIntegerFraction(arthur.IntegerFractionPropProper)
	}
	fmt.Fprintf(w, exp.LaTeX())
}
