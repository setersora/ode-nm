// This source file contains implemetation of some numerical methods.

package main

import (
	"fmt"
	"math"
	"os"

	"gonum.org/v1/plot/plotter"
)

// ExplicitEuler returns data for plot building based on explicit Euler method.
func ExplicitEuler(start, end, y0 float64, count int) plotter.XYs {
	step, xRange := GetRange(start, end, count)
	solution := make(plotter.XYs, count)

	for i := range solution {
		x := xRange[i]

		solution[i].X = x
		if i == 0 {
			solution[i].Y = y0
		} else {
			y := solution[i - 1].Y
			solution[i].Y = y + step * F(x, y)
		}
	}

	return solution
}

// ModifiedEuler returns data for plot building based on modified Euler method.
func ModifiedEuler(start, end, y0 float64, count int) plotter.XYs {
	step, xRange := GetRange(start, end, count)
	solution := make(plotter.XYs, count)

	for i := range solution {
		x := xRange[i]

		solution[i].X = x
		if i == 0 {
			solution[i].Y = y0
		} else {
			x = solution[i - 1].X
			y := solution[i - 1].Y
			predict := y + step * F(x, y)
			solution[i].Y = y + (step / 2) * (F(x, y) + F(x + step, predict))
		}
	}

	return solution
}

// Cauchy returns data for plot building based on Cauchy method.
func Cauchy(start, end, y0 float64, count int) plotter.XYs {
	step, xRange := GetRange(start, end, count)
	solution := make(plotter.XYs, count)

	for i := range solution {
		x := xRange[i]

		solution[i].X = x
		if i == 0 {
			solution[i].Y = y0
		} else {
			x = solution[i - 1].X
			y := solution[i - 1].Y
			predict := y + (step / 2) * F(x, y)
			solution[i].Y = y + step * F(x + step / 2, predict)
		}
	}

	return solution
}

// RungeKutta returns data for plot building based on RungeKutta 4th-Order method.
func RungeKutta(start, end, y0 float64, count int) plotter.XYs {
	step, xRange := GetRange(start, end, count)
	solution := make(plotter.XYs, count)

	for i := range solution {
		x := xRange[i]

		solution[i].X = x
		if i == 0 {
			solution[i].Y = y0
		} else {
			x = solution[i - 1].X
			y := solution[i - 1].Y

			k1 := step * F(x, y)
			k2 := step * F(x + step / 2, y + k1 / 2)
			k3 := step * F(x + step / 2, y + k2 / 2)
			k4 := step * F(x + step, y + k3)

			solution[i].Y = y + (k1 / 6) + 2 * (k2 / 6) +
				2 * (k3 / 6) + (k4 / 6)
		}
	}

	return solution
}

func ImplicitEuler(start, end, y0 float64, count int) plotter.XYs {
	step, xRange := GetRange(start, end, count)
	solution := make(plotter.XYs, count)

	for i := range solution {
		if i == len(solution) - 1 {
			break
		}
		x := xRange[i + 1]

		solution[i].X = x
		if i == 0 {
			solution[i].Y = y0
		} else {
			prevY := solution[i - 1].Y
			g := func(y float64) float64 {
				return prevY + step * F(x, y)
			}
			yn1 := SimpleIteration(prevY, g)
			solution[i].Y = g(yn1)
		}
	}

	return solution
}

// Tailor2th returns data for plot building based on Tailor 2th-Order method.
func Tailor2th(start, end, y0 float64, count int) plotter.XYs {
	step, xRange := GetRange(start, end, count)
	solution := make(plotter.XYs, count)

	for i := range solution {
		x := xRange[i]

		solution[i].X = x
		if i == 0 {
			solution[i].Y = y0
		} else {
			x = solution[i - 1].X
			y := solution[i - 1].Y
			solution[i].Y = y + step * F(x, y) +
				(math.Pow(step, 2) / 2) * SecondDerivative(x, y)
		}
	}

	return solution
}

// Tailor3th returns data for plot building based on Tailor 3th-Order method.
func Tailor3th(start, end, y0 float64, count int) plotter.XYs {
	step, xRange := GetRange(start, end, count)
	solution := make(plotter.XYs, count)

	for i := range solution {
		x := xRange[i]

		solution[i].X = x
		if i == 0 {
			solution[i].Y = y0
		} else {
			x = solution[i - 1].X
			y := solution[i - 1].Y
			solution[i].Y = y + step * F(x, y) +
				(math.Pow(step, 2) / 2) * SecondDerivative(x, y) +
				(math.Pow(step, 3) / 6) * ThirdDerivative(x, y)
		}
	}

	return solution
}

// Adams returns data for plot building based on Adams 2th-Order method.
func Adams(start, end, y0 float64, count int) plotter.XYs {
	step, xRange := GetRange(start, end, count)
	solution := make(plotter.XYs, count)

	for i := range solution {
		x := xRange[i]

		solution[i].X = x
		if i == 0 {
			solution[i].Y = y0
		} else {
			xn1 := x
			xn := solution[i - 1].X
			yn := solution[i - 1].Y
			fn := F(xn, yn)
			g := func(y float64) float64 {
				return yn + step / 2 * (fn + F(xn1, y))
			}
			yn1 := SimpleIteration(yn, g)
			solution[i].Y = g(yn1)
		}
	}

	return solution
}

func Dichotomy(a, b float64, f func(float64) float64) float64 {
	epsilon := 0.00000001

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

		fmt.Println("Dichotomy error: fc * fb < 0")
		os.Exit(1)
	}

	return (a + b) / 2
}
