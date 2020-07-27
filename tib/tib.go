package tib

import (
	"fmt"
	"net/http"

	"github.com/mathyourlife/arthur"
)

func ArthurHandler(w http.ResponseWriter, r *http.Request) {
	var exp arthur.Expression
	exp = arthur.NewRandomIntegerFraction()
	fmt.Fprintf(w, exp.LaTeX())
}
