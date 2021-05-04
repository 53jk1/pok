// Napisz program obliczający objętość V sześcianu o długości L : V = L^3,
// obliczoną dla trzech różnych wartości L.

// V: objętość sześcianu
// L: długość boku

package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"

	"github.com/pa-m/numgo"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

var np = numgo.NumGo{}

func main() {

	declare()

	xys, err := readData("data.txt")
	if err != nil {
		log.Fatalf("could not read data.txt: %v", err)
	}

	_ = xys

	err = plotData("out.png", xys)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}

	// Wykonaj ręcznie obliczenie V = L^3, gdy L jest tablicą z trzema
	// elementami. Oznacza to, że należy obliczyć V dla każdej wartości L.
	//[0]1^3 = 1
	//[1]2^3 = 8
	//[2]3^3 = 27
}

func readData(path string) (plotter.XYs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var xys plotter.XYs
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}
		xys = append(xys, struct{ X, Y float64 }{x, y})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return xys, nil
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

	fmt.Printf("Objetosc szescianu 1: %.f\nObjetosc szescianu 2: %.f\nObjetosc szescianu 3: %.f\n", V1, V2, V3)
	fmt.Println("linspace:", np.Linspace(1, V1, 30, true))
	fmt.Println("linspace:", np.Linspace(1, V2, 30, true))
	fmt.Println("linspace:", np.Linspace(1, V3, 30, true))

}

func plotData(path string, xys plotter.XYs) error {

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p := plot.New()

	// V on out.png
	s, err := plotter.NewScatter(xys)
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	// L on out.png
	l, err := plotter.NewLine(plotter.XYs{{0, 0}, {3, 3}})

	if err != nil {
		return fmt.Errorf("could not create line: %v", err)
	}
	p.Add(l)

	l, err = plotter.NewLine(plotter.XYs{{0, 0}, {3, 4}})

	if err != nil {
		return fmt.Errorf("could not create line: %v", err)
	}
	p.Add(l)

	l, err = plotter.NewLine(plotter.XYs{{0, 0}, {3, 5}})

	if err != nil {
		return fmt.Errorf("could not create line: %v", err)
	}
	p.Add(l)

	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		return fmt.Errorf("could not create writer: %v", err)
	}

	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to out.png: %v", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close out.png: %v", err)
	}
	return nil
}

// W programie użyj funkcji linspace, aby obliczyć i wydrukować trzy
// wartości L, równo rozmieszczone w przedziale [1, 3].
