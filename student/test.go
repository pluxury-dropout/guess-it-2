package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// var y_axis []float64
	// var x_axis []float64

	var k int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.Atoi(input)
		if err != nil {
			log.Print(err)
			return
		}
		// zxc1 := math.Round(float64(num))
		// y_axis = append(y_axis, zxc1)
		// x_axis = append(x_axis, float64(k))
		// b := GetB(x_axis, y_axis)
		// a := GetA(x_axis, y_axis)
		// sd := StandardDeviation(y_axis)
		// res1 := int((a + b*float64(k)) - sd)
		// res2 := int((a + b*float64(k)) + sd)
		k += 1
		fmt.Printf("%v %v\n", num, num*2)

	}
}

func Mean(n []float64) float64 {
	var mean float64
	for i := 0; i < len(n); i++ {
		mean += n[i]
	}
	mean = mean / float64(len(n))
	return mean
}

func StandardDeviation(n []float64) float64 {
	var dif float64
	var sum float64
	for i := 0; i < len(n); i++ {
		dif = math.Pow(n[i]-Mean(n), 2)
		sum += dif
	}
	return math.Sqrt(sum)
}

func PearsonCorrCoef(x []float64, y []float64) float64 {
	var net_sum float64
	var dem_x float64
	var dem_y float64
	var p_coef float64

	for i := 0; i < len(y); i++ {
		net_sum += ((x[i] - Mean(x)) * (y[i] - Mean(y)))
		temp1 := math.Pow(x[i]-Mean(x), 2)
		dem_x += temp1
		temp2 := math.Pow(y[i]-Mean(y), 2)
		dem_y += temp2
	}
	dem_x = math.Sqrt(dem_x)
	dem_y = math.Sqrt(dem_y)
	p_coef = net_sum / (dem_x * dem_y)
	return p_coef
}

func GetB(x, y []float64) float64 {
	// general equaation : y = a + bx
	// b = p_coef * sd(y) / sd(x)
	// a = mean(y) - b * mean(x)
	b := PearsonCorrCoef(x, y) * StandardDeviation(y) / StandardDeviation(x)
	return b
}

func GetA(x, y []float64) float64 {
	a := Mean(y) - GetB(x, y)*Mean(x)
	return a
}
