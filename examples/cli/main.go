package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mathyourlife/arthur"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var exp arthur.Expression
	exp = arthur.NewRandIntegerFraction()
	fmt.Println(exp.LaTeX())
}
