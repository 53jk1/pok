// Napisz program, który oblicza objętość V sześcianu o bokach długości
// L = 4 cm i wypisuje wynik na ekranie.
// Zarówno V , jak i L powinny być zdefiniowane jako oddzielne zmienne w
// programie.
// Uruchom program i potwierdź wydrukowanie prawidłowego wyniku.
package main

import (
	"fmt"
	"math"
)

func main() {
	var V float64
	var L float64 = 4

	V = math.Pow(L, 3)

	fmt.Println("Wynik to:", V)

}
