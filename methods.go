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
			y := solution[i - 1].Y
			secondDerivative := Fx(x, y) + Fy(x, y) * F(x, y)
			solution[i].Y = y + step * F(x, y) +
				(math.Pow(step, 2) / 2) * secondDerivative
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
			y := solution[i - 1].Y
			secondDerivative := Fx(x, y) + Fy(x, y) * F(x, y)
			thirdDerivative := (100 * y) + (50 * (2 * x - 1.45)) +
				Fx(x, y) + Fy(x, y) + Fy(x, y) * secondDerivative
			solution[i].Y = y + step * F(x, y) +
				(math.Pow(step, 2) / 2) * secondDerivative +
				(math.Pow(step, 3) / 6) * thirdDerivative
		}
	}

	return solution
}
