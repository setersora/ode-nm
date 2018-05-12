// This source file contains implemetation of some numerical methods.

package main

import (
	"math"

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

// Tailor returns data for plot building based on Tailor method.
func Tailor(start, end, y0 float64, count int) plotter.XYs {
	step, xRange := GetRange(start, end, count)
	solution := make(plotter.XYs, count)

	for i := range solution {
		x := xRange[i]

		solution[i].X = x
		if i == 0 {
			solution[i].Y = y0
		} else {
			y := solution[i - 1].Y
			solution[i].Y = y + step * F(x, y) +
				(math.Pow(step, 2) / 2) *
				(Fx(x, y) + Fy(x, y) * F(x, y))
		}
	}

	return solution
}
