package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var data []float64
	for reader.Scan() {
		num, err := strconv.ParseFloat(reader.Text(), 64)
		if err != nil {
			fmt.Println("Error: number is not valid")
			return
		}
		data = append(data, num)
		min, max := RangeGuess(data)
		fmt.Println(int(math.Round(min)), int(math.Round(max)))
	}
}

func RangeGuess(data []float64) (float64, float64) {
	var min, max float64
	slope, intercept := LinearRegression(data)
	cov := PCC(data)
	prediction := slope*float64(len(data)-1) + intercept
	max = prediction + prediction*(1-cov)
	min = prediction - prediction*(1-cov)
	return min, max
}

func LinearRegression(data []float64) (float64, float64) {
	sumX := 0.00
	sumX2 := 0.00
	sumY := 0.00
	sumXY := 0.00
	sumSQR := 0.00
	lngth := float64(len(data))
	for i := 0; i < len(data); i++ {
		sumX += float64(i + 1)
		sumX2 += float64(i)
		sumY += data[i]
		sumSQR += float64((i + 1) * (i + 1))
		sumXY += (float64(i)) * data[i]
	}
	slope := ((lngth * sumXY) - (sumX2 * sumY)) / ((lngth * sumSQR) - (sumX * sumX))
	if math.IsNaN(slope) {
		slope = 0
	}
	intercept := data[len(data)-1] - (slope * float64(len(data)))
	return slope, intercept
}

func PCC(data []float64) float64 {
	sumX := 0.00
	sumY := 0.00
	sumXY := 0.00
	sumSQRX := 0.00
	sumSQRY := 0.00
	lngth := float64(len(data))
	for i := 0; i < len(data); i++ {
		sumX += float64(i + 1)
		sumY += data[i]
		sumSQRX += float64((i + 1) * (i + 1))
		sumSQRY += data[i] * data[i]
		sumXY += (float64(i + 1)) * data[i]
	}
	covariance := ((lngth * sumXY) - (sumX * sumY)) / math.Sqrt(((lngth*sumSQRX)-(sumX*sumX))*((lngth*sumSQRY)-(sumY*sumY)))
	return covariance
}
