package main

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// F returns the value of the original f(x, y) given in the task at point (x, y).
func F(x, y float64) float64 {
	return 50 * y * (math.Pow(x, 2) - 1.45 * x + 0.51)
}

// Fx returns the value of the partial derivative by x of the original
// f(x, y) given in the task at point (x, y).
func Fx(x, y float64) float64 {
	return 50 * y * (2 * x - 1.45)
}

// Fy returns the value of the partial derivative by x of the original
// f(x, y) given in the task at point (x, y).
func Fy(x, y float64) float64 {
	return 50 * (math.Pow(x, 2) - 1.45 * x + 0.51)
}

// Y returns the value of the manually obtained precise solution of the task
// at point with abscissa x.
func Y(x float64) float64 {
	arg := 50 * ((math.Pow(x, 3) / 3) -  1.45 * (math.Pow(x, 2) / 2) + 0.51 * x)
	return 0.1 * math.Exp(arg)
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
		"Precise", PreciseSolution(start, end, pointsCount))
	if err != nil {
		panic(err)
	}

	if err := p.Save(10*vg.Inch, 10*vg.Inch, filename); err != nil {
		panic(err)
	}
}

func main() {
	DrawPlot(0, 1, 0.1, 10, "10 steps", "10_steps.png")
	DrawPlot(0, 1, 0.1, 100, "100 steps", "100_steps.png")
	DrawPlot(0, 1, 0.1, 10000, "10000 steps", "10000_steps.png")
}
