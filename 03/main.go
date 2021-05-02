// Napisz program obliczający objętość V sześcianu o długości L : V = L^3,
// obliczoną dla trzech różnych wartości L.

// V: objętość sześcianu
// L: długość boku

package main

import (
	"fmt"
	"math"

	"github.com/pa-m/numgo"
)

var np = numgo.NumGo{}

func main() {

	declare()

	// Wykonaj ręcznie obliczenie V = L^3, gdy L jest tablicą z trzema
	// elementami. Oznacza to, że należy obliczyć V dla każdej wartości L.
	//[0]1^3 = 1
	//[1]2^3 = 8
	//[2]3^3 = 27
}

func declare() {
	var L1 float64 = 3 // L
	var L2 float64 = 4 // L
	var L3 float64 = 5 // L
	V1 := math.Pow(L1, 3)
	V2 := math.Pow(L2, 3)
	V3 := math.Pow(L3, 3)

	fmt.Printf("Objetosc szescianu 1: %.f\nObjetosc szescianu 2: %.f\nObjetosc szescianu 3: %.f\n", V1, V2, V3)
	fmt.Println("linspace:", np.Linspace(1, V1, 3, true))
	fmt.Println("linspace:", np.Linspace(1, V2, 3, true))
	fmt.Println("linspace:", np.Linspace(1, V3, 3, true))
}

// W programie użyj funkcji linspace, aby obliczyć i wydrukować trzy
// wartości L, równo rozmieszczone w przedziale [1, 3].
