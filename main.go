package main

import (
	"fmt"
)

// Define the Anscombe Quartet dataset as a map
var anscombeQuartet = map[string][][]float64{
	"x": {
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, // x1 values
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, // x2 values
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, // x3 values
		{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},     // x4 values
	},
	"y": {
		{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}, // y1 values
		{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},    // y2 values
		{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}, // y3 values
		{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},  // y4 values
	},
}

// LinearRegression calculates the slope and intercept for a given x and y dataset
func LinearRegression(x, y []float64) (float64, float64) {
	var sumX, sumY, sumXY, sumX2 float64
	n := float64(len(x))

	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept := (sumY - slope*sumX) / n

	return slope, intercept
}

// CalculateAndPrintResults iterates over each dataset and prints the slope and intercept
func CalculateAndPrintResults() {
	for i := 0; i < len(anscombeQuartet["x"]); i++ {
		x := anscombeQuartet["x"][i]
		y := anscombeQuartet["y"][i]
		slope, intercept := LinearRegression(x, y)
		fmt.Printf("Set %d: Slope = %.2f, Intercept = %.2f\n", i+1, slope, intercept)
	}
}

func main() {
	CalculateAndPrintResults()
}
