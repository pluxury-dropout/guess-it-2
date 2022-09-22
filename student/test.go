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
	var y_axis []int
	var x_axis []int

	k := 1

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		input := scanner.Text()
		num, err := strconv.Atoi(input)
		if err != nil {
			log.Print(err)
			fmt.Println(0000)
			return
		}
		if len(y_axis) > 3 || len(x_axis) > 3 {

			y_axis = append(y_axis, num)
			x_axis = append(x_axis, k)
			fmt.Printf("%v %v\n", num, num*2)
			continue
		}

		b := GetB(x_axis, y_axis)
		a := GetA(x_axis, y_axis)
		sd := StandardDeviation(y_axis)
		res1 := ((a + b*k) - sd)
		res2 := ((a + b*k) + sd)
		k++
		fmt.Printf("%v %v\n", res1, res2)

	}
}

func Mean(n []int) int {
	var mean int

	for i := 0; i < len(n); i++ {
		mean += n[i]
	}
	mean = mean / (len(n))
	return mean
}

func StandardDeviation(n []int) int {
	var dif float64
	var sum float64
	for i := 0; i < len(n); i++ {
		dif = math.Pow(float64(n[i]-Mean(n)), 2)
		sum += dif
	}

	return int(math.Sqrt(sum))
}

func PearsonCorrCoef(x []int, y []int) int {
	var net_sum int
	var dem_x int
	var dem_y int
	var p_coef int

	for i := 0; i < len(y); i++ {
		net_sum += (x[i] - Mean(x)) * (y[i] - Mean(y))
		temp1 := int(math.Pow(float64(x[i]-Mean(x)), 2))
		dem_x += temp1
		temp2 := int(math.Pow(float64(y[i]-Mean(y)), 2))
		dem_y += temp2
	}
	dem_x = int(math.Sqrt(float64(dem_x)))
	dem_y = int(math.Sqrt(float64(dem_y)))
	if dem_x == 0 || dem_y == 0 {
		return 1 / 2
	}
	p_coef = net_sum / (dem_x * dem_y)
	return p_coef
}

func GetB(x []int, y []int) int {
	b := PearsonCorrCoef(x, y) * StandardDeviation(y) / StandardDeviation(x)
	return b
}

func GetA(x []int, y []int) int {
	a := Mean(y) - GetB(x, y)*Mean(x)
	return a
}
