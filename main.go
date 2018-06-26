package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// F returns the value of the original f(x, y) given in the task at point (x, y).
// Needs to be changed for other equations.
func F(x, y float64) float64 {
	return 50 * y * (math.Pow(x, 2) - 1.45 * x + 0.51)
}

// Fx returns the value of the partial derivative by x of the original
// f(x, y) given in the task at point (x, y).
// Needs to be changed for other equations.
func Fx(x, y float64) float64 {
	return 50 * y * (2 * x - 1.45)
}

// Fy returns the value of the partial derivative by x of the original
// f(x, y) given in the task at point (x, y).
// Needs to be changed for other equations.
func Fy(x, y float64) float64 {
	return 50 * (math.Pow(x, 2) - 1.45 * x + 0.51)
}

// Y returns the value of the manually obtained precise solution of the task
// at point with abscissa x.
// Needs to be changed for other equations.
func Y(x float64) float64 {
	arg := 50 * ((math.Pow(x, 3) / 3) -  1.45 * (math.Pow(x, 2) / 2) + 0.51 * x)
	return 0.1 * math.Exp(arg)
}

// SecondDerivative is simply second derivative of F.
// Is used in Tailor methods.
// Needs to be changed for other equations.
func SecondDerivative(x, y float64) float64 {
	return Fx(x, y) + Fy(x, y) * F(x, y)
}

// ThirdDerivative is simpy third derivative of F.
// Is used in Tailor methods.
// Needs to be changed for other equations.
func ThirdDerivative(x, y float64) float64 {
	return (100 * y) + (50 * (2 * x - 1.45)) +
		Fx(x, y) + Fy(x, y) + Fy(x, y) * SecondDerivative(x, y)
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

type taskFunction func(float64) float64

// SimpleIteration is used as helper to solve simple equations
// in Adams and implicit Euler methods.
func SimpleIteration(x float64, f taskFunction) float64 {
	epsilon := 0.00000001
	methodError := float64(100)
	iterCount := 0
	for methodError > epsilon {
		newX := f(x)
		methodError = math.Abs(x - newX)
		x = newX
		iterCount += 1
		if iterCount > 1000 {
			fmt.Println("Simple iteration method returns",
				"inaccurate result!")
			return x
		}
	}
	return x
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

// DrawPlot draws a plot of all methods' solutions for pointsCount points
// named plotName and saves it to filename.
func DrawPlot(start, end, y0 float64, pointsCount int, plotName, filename string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = plotName
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"Precise", PreciseSolution(start, end, pointsCount),
		"Ecplicit Euler", ExplicitEuler(start, end, y0, pointsCount),
		"Modified Euler", ModifiedEuler(start, end, y0, pointsCount),
		"Cauchy", Cauchy(start, end, y0, pointsCount),
		"RungeKutta", RungeKutta(start, end, y0, pointsCount),
		"Implicit Euler", ImplicitEuler(start, end, y0, pointsCount),
		"Tailor 2nd", Tailor2nd(start, end, y0, pointsCount),
		"Tailor 3rd", Tailor3rd(start, end, y0, pointsCount),
		"Adams", Adams(start, end, y0, pointsCount))
	if err != nil {
		panic(err)
	}

	if err := p.Save(16*vg.Inch, 9*vg.Inch, filename); err != nil {
		panic(err)
	}
}

func main() {
	DrawPlot(0, 1, 0.1, 25, "25 steps",       "ALL_25_steps.png")
	DrawPlot(0, 1, 0.1, 75, "75 steps",       "ALL_75_steps.png")
	DrawPlot(0, 1, 0.1, 150, "150 steps",     "ALL_150_steps.png")
	DrawPlot(0, 1, 0.1, 750, "750 steps",     "ALL_750_steps.png")
	DrawPlot(0, 1, 0.1, 5000, "5000 steps",   "ALL_5000_steps.png")
}
