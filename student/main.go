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
	// general equaation : y = a + bx
	// b = p_coef * sd(y) / sd(x)
	// a = mean(y) - b * mean(x)

	var x_axis []float64
	y_axis := MakeLinesArr(os.Args[1])
	for i := 0; i < len(y_axis); i++ {
		x_axis = append(x_axis, float64(i))
	}
	p_coef := PearsonCorrCoef(x_axis, y_axis)
	b := p_coef * StandardDeviation(y_axis) / StandardDeviation(x_axis)
	a := Mean(y_axis) - b*Mean(x_axis)

	// fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", b, a)
	// fmt.Printf("Pearson Correlation Coefficient: %.10f\n", p_coef)
	sd := StandardDeviation(y_axis)
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var k float64
	for scanner.Scan() {

		fmt.Printf("%v %v", (a+b*k)-sd, (a+b*k)+sd)
		k++
	}
}

func MakeLinesArr(s string) []float64 {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	var x float64

	var arr []float64
	for scanner.Scan() {
		// do something with a line
		temp, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("error1")
		}
		x = float64(temp)
		arr = append(arr, x)
	}
	return arr
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
