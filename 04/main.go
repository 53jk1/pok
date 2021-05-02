// Napisz program, który przechowuje sumę pięciu liczb calkowitych w jednej
// zmiennej, a następnie tworzy inną zmienną ze średnią z tych pięciu liczb.
// Wydrukuj średnią na ekranie i sprawdź, czy wynik jest prawidłowy.

package main

import "fmt"

func main() {
	var a, b, c, d, e float32 = 1844, 184, 18, 123, 70

	sum := a + b + c + d + e

	average := sum / 5

	fmt.Printf("Average: %.1f\n", average) // that means precision 1
}
