// Napisz program, który obliczy zarówno obwód C, jak i pole A koła o
// promieniu r = 2 cm.
// Niech wyniki zostaną wydrukowane na ekranie w jednym wierszu z
// odpowiednim tekstem.
// Wszystkie zmienne C, A i r powinny być zdefiniowane jako oddzielne
// zmienne w programie.
// Uruchom program i potwierdź, że drukowane są prawidłowe wyniki.

package main

import (
	"fmt"
	"math"
)

func main() {
	circumferenceOfCircle()
	circleField()
}

const (
	r  = 2 // radius
	Pi = 3.14159265358979323846264338327950288419716939937510582097494459
)

func circumferenceOfCircle() {
	var C float64 = 2 * Pi * r
	fmt.Printf("%f\n", C)
}

func circleField() {
	var A float64 = Pi * math.Exp2(r)
	fmt.Printf("%f\n", A)
}
