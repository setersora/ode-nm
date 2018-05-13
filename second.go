package main

import (
	"fmt"
	"math"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const N = 10
const A = 2 + 0.1 * N

func F(x, y, dy float64) float64 {
	return y + A*x*(1-x) + 2*A + 2
}

func F1(dy float64) float64 {
	return dy - math.Exp(1.0) - math.Exp(-1.0) - A
}

func Y(x float64) float64 {
	return -2 + math.Exp(-x) + math.Exp(x) + A * (x - 1) * x
}

type method func(float64, float64, float64, float64, int, task) (plotter.XYs, []float64)

func Fire(m float64, n int, curmeth method) float64 {
	_, dys := curmeth(0, 1, m, -A, n, F)
	return F1(dys[n - 2])
}

func Dichotomy(a, b float64, f func(float64) float64) float64 {
	epsilon := 0.0000000001

	fa := f(a)
	fb := f(b)

	if fa == 0 {
		return a
	}
	if fb == 0 {
		return b
	}

	if fa * fb > 0 {
		fmt.Println("Dichotomy error: fa * fb > 0")
		os.Exit(1)
	}

	for math.Abs(a - b) > epsilon {
		c := (a + b) / 2
		fc := f(c)

		if fc == 0 {
			return c
		}

		if fc * fa < 0 {
		    b = c
		    fb = fc
		    continue
		}

		if fc * fb < 0 {
		    a = c
		    fa = fc
		    continue
		}

		fmt.Println("Dichotomy error: fc")
		os.Exit(1)
	}

	return (a + b) / 2
}

// DrawPlot draws a plot of all methods' solutions for pointsCount points
// named plotName and saves it to filename.
func DrawPlot(start, end float64, pointsCount int, plotName, filename string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = plotName
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	curmeth := ExplicitEuler
	M := Dichotomy(-10, 10,
		func(m float64) float64 { return Fire(m, pointsCount, curmeth) })
	eulers, _ := ExplicitEuler(start, end, M, -A, pointsCount, F)

	curmeth = RungeKutta
	M = Dichotomy(-10, 10,
		func(m float64) float64 { return Fire(m, pointsCount, curmeth) })
	runges, _ := RungeKutta(start, end, M, -A, pointsCount, F)

	err = plotutil.AddLinePoints(p,
		"Precise", PreciseSolution(start, end, pointsCount),
		"Explicit Euler", eulers,
		"Runge-Kutta", runges)
	if err != nil {
		panic(err)
	}

	if err := p.Save(16*vg.Inch, 9*vg.Inch, filename); err != nil {
		panic(err)
	}
}

func main() {
	DrawPlot(0, 1, 25, "25 steps",       "SECOND_25_steps.png")
	DrawPlot(0, 1, 150, "150 steps",     "SECOND_150_steps.png")
}

// GetRange returns a range of *count* values between *start* and *end*.
func GetRange(start, end float64, count int) (float64, []float64) {
	var step float64
	var resultRange []float64

	step = (end - start) / float64(count)
	value := start
	for i := 0; i < count; i++ {
		resultRange = append(resultRange, value)
		value += step
	}

	return step, resultRange
}

// PreciseSolution returns data for plot building based on manually obtained
// solution implemented as Y(x).
func PreciseSolution(start, end float64, count int) plotter.XYs {
	_, xRange := GetRange(start, end, count)
	solution := make(plotter.XYs, count)

	for i := range solution {
		x := xRange[i]

		solution[i].X = x
		solution[i].Y = Y(x)
	}

	return solution
}

type task func(float64, float64, float64) float64

// ExplicitEuler returns data for plot building based on explicit Euler method.
func ExplicitEuler(start, end, y0, dy0 float64, count int, f task) (plotter.XYs, []float64) {
	h, xRange := GetRange(start, end, count)
	solution := make(plotter.XYs, count)
	dys := []float64{}

	for i := range solution {
		x := xRange[i]

		solution[i].X = x
		if i == 0 {
			solution[i].Y = y0
			dys = append(dys, dy0)
		} else {
			x = solution[i - 1].X
			y := solution[i - 1].Y
			dy := dys[i - 1]
			dyNext := dy + h * f(x, y, dy)
			dys = append(dys, dyNext)
			solution[i].Y = y + h * dy
		}
	}

	return solution, dys
}

// RungeKutta returns data for plot building based on RungeKutta 4th-Order method.
func RungeKutta(start, end, y0, dy0 float64, count int, f task) (plotter.XYs, []float64) {
	h, xRange := GetRange(start, end, count)
	solution := make(plotter.XYs, count)
	dys := []float64{}

	for i := range solution {
		x := xRange[i]

		solution[i].X = x
		if i == 0 {
			solution[i].Y = y0
			dys = append(dys, dy0)
		} else {
			x = solution[i - 1].X
			y := solution[i - 1].Y
			dy := dys[i - 1]

			k1y := h * dy
			k1dy := h * f(x, y, dy)

			k2y := h * (dy + k1dy/2)
			k2dy := h * f(x + h/2, y + k1y/2, dy + k1dy/2)

			k3y := h * (dy + k2dy/2)
			k3dy := h * f(x + h/2, y + k2y/2, dy + k2dy/2)

			k4y := h * (dy + k3dy)
			k4dy := h * f(x + h, y + k3y, dy + k3dy)

			solution[i].Y = y + (k1y / 6) + 2 * (k2y / 6) +
				2 * (k3y / 6) + (k4y / 6)
			dys = append(dys, dy + (k1dy + 2*(k2dy + k3dy) + k4dy) / 6)
		}
	}

	return solution, dys
}
