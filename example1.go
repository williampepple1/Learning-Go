package main

import (
	fm "fmt"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fm.Printf("boiling point = %g°F or %g°C\n", f, c)
}
