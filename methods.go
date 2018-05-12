// This source file contains implemetation of some numerical methods.

package main

import (
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

