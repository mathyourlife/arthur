package tib

import (
	"fmt"
	"net/http"

	"github.com/mathyourlife/arthur"
)

func ArthurFractionHandler(w http.ResponseWriter, r *http.Request) {
	fractionType := r.URL.Query().Get("type")
	var exp arthur.Expression
	if fractionType == "improper" {
		exp = arthur.NewRandIntegerFraction(arthur.IntegerFractionPropImproper)
	} else if fractionType == "unit" {
		exp = arthur.NewRandIntegerFraction(arthur.IntegerFractionPropUnit)
	} else {
		exp = arthur.NewRandIntegerFraction(arthur.IntegerFractionPropProper)
	}
	fmt.Fprintf(w, exp.LaTeX())
}

func ArthurIntegerHandler(w http.ResponseWriter, r *http.Request) {
	var exp arthur.Expression
	exp = arthur.NewRandInt()
	fmt.Fprintf(w, exp.LaTeX())
}
