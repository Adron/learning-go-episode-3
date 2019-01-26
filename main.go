package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

const (
	width, height = 600, 600
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {

	var outputResult string

	outputResult = "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: blue; fill: white; stroke-width: 0.7' " +
		"width='" + strconv.Itoa(width) + "' height='" + strconv.Itoa(height) + "'>"

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			outputResult = outputResult + "<polygon points='" +
				ax + "," + ay +
				bx + "," + by +
				cx + "," + cy +
				dx + "," + dy +
				"'/>\n"
		}
	}
	outputResult = outputResult + "</svg>"
	fmt.Printf(outputResult)
	err := ioutil.WriteFile("surface-plot.svg", []byte(outputResult), 0644)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}

func corner(i, j int) (string, string) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	xResult := strconv.FormatFloat(sx, 'f', -1, 64)
	yResult := strconv.FormatFloat(sy, 'f', -1, 64)
	return xResult, yResult
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)
}
