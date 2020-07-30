package tib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mathyourlife/arthur"
)

func ArthurFractionHandler(w http.ResponseWriter, r *http.Request) {
	format := r.URL.Query().Get("format")
	fractionType := r.URL.Query().Get("type")
	var exp arthur.Expression
	if fractionType == "improper" {
		exp = arthur.NewRandIntegerFraction(arthur.IntegerFractionPropImproper)
	} else if fractionType == "unit" {
		exp = arthur.NewRandIntegerFraction(arthur.IntegerFractionPropUnit)
	} else {
		exp = arthur.NewRandIntegerFraction(arthur.IntegerFractionPropProper)
	}
	if format == "json" {
		content, _ := json.Marshal(exp)
		fmt.Fprint(w, string(content))
	} else if format == "latex" {
		fmt.Fprintf(w, exp.LaTeX())
	} else {
		fmt.Fprintf(w, exp.LaTeX())
	}
}

func ArthurFractionSumHandler(w http.ResponseWriter, r *http.Request) {
	var a, b, exp arthur.Expression
	a = arthur.NewRandIntegerFraction()
	b = arthur.NewRandIntegerFraction()
	exp = &arthur.Sum{a, b}
	fmt.Fprintf(w, exp.LaTeX())
}

func ArthurIntegerHandler(w http.ResponseWriter, r *http.Request) {
	format := r.URL.Query().Get("format")

	intSize := r.URL.Query().Get("size")
	var exp arthur.Expression
	if intSize == "small" {
		exp = arthur.NewRandInt(arthur.IntegerPropertySmall)
	} else if intSize == "medium" {
		exp = arthur.NewRandInt(arthur.IntegerPropertyMedium)
	} else if intSize == "large" {
		exp = arthur.NewRandInt(arthur.IntegerPropertyLarge)
	} else {
		exp = arthur.NewRandInt(arthur.IntegerPropertySmall)
	}

	if format == "json" {
		content, _ := json.Marshal(exp)
		fmt.Fprint(w, string(content))
	} else if format == "latex" {
		fmt.Fprintf(w, exp.LaTeX())
	} else {
		fmt.Fprintf(w, exp.LaTeX())
	}
}

func SetupMux(prefix string, mux *http.ServeMux) error {
	base, err := url.Parse(prefix)
	if err != nil {
		return err
	}

	u, _ := base.Parse("integer")
	mux.HandleFunc(u.String(), ArthurIntegerHandler)
	u, _ = base.Parse("fraction")
	mux.HandleFunc(u.String(), ArthurFractionHandler)
	u, _ = base.Parse("sum")
	mux.HandleFunc(u.String(), ArthurFractionSumHandler)
	return nil
}
