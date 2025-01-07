package colors

//package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

// home/userone/Documentos/AppReact/flaskAPI/colors/colorsGradients.py
// Works on RGB [255,255,255]
func interpolate(start []int, end []int, ratio float32) []int {
	r := int(ratio*float32(start[0]) + (1.0-ratio)*float32(end[0]))
	g := int(ratio*float32(start[1]) + (1.0-ratio)*float32(end[1]))
	b := int(ratio*float32(start[2]) + (1.0-ratio)*float32(end[2]))
	return []int{r, g, b}
}

// Works on RGB [255,255,0]-> #ffff00
func Rgb2Hex(color []int) string {

	return fmt.Sprintf("#%02x%02x%02x", color[0], color[1], color[2])
}

// "#FFFFFF" -> [255,255,255]
func Hex2Rgb(c string) []int {

	r, _ := strconv.ParseInt(c[1:3], 16, 16)
	g, _ := strconv.ParseInt(c[3:5], 16, 16)
	b, _ := strconv.ParseInt(c[5:7], 16, 16)
	return []int{int(r), int(g), int(b)}
}

// Recibe una lista Hex y devuelve una lista en RGB
// HexList2RgbList(['#ff0000', '#00ff00', '#0000ff'])
func HexList2RgbList(Hgrad []string) [][]int {

	var RgbL [][]int
	for _, h := range Hgrad {
		RgbL = append(RgbL, Hex2Rgb(h))
	}
	return RgbL

}

// Recibe una gradiente RGB y la cantidad de colores
// Devuelve una lista de colores Hex
func HexMultiColorScale(grad [][]int, steps int) []string {

	//fmt.Println("HexMultiColorScale:---------------------")
	var colors []string
	var ratios []float32

	for i := 0; i < steps; i++ {

		ratios = append(ratios, 1/float32(steps-1)*float32(i))

	}

	//fmt.Println("ratios:", ratios)

	for _, r := range ratios {
		colors = append(colors, Rgb2Hex(multiColor(grad, r)))
	}
	return colors
}

// Recibe una gradiente RGB y la cantidad de colores
// Devuelve una lista de colores RGB
func RgbMultiColorScale(grad [][]int, steps int) [][]int {

	//fmt.Println("HexMultiColorScale:---------------------")
	var colors [][]int
	var ratios []float32

	for i := 0; i < steps; i++ {

		ratios = append(ratios, 1/float32(steps-1)*float32(i))

	}

	//fmt.Println("ratios:", ratios)

	for _, r := range ratios {
		colors = append(colors, multiColor(grad, r))
	}
	return colors
}

// Recibe una gradiente y el porcentaje ej 50%
// Devuelve un color RGB interpolado dentro de la gradiente

func multiColor(colors [][]int, ratio float32) []int {
	count := len(colors)
	stepbase := float32(1.0 / float32(count-1))
	//fmt.Println("stepbase:", stepbase)

	interval := count - 1          // to fix 1!=0.99999999;
	for i := 1; i < count-1; i++ { // quito los extremos

		if ratio <= float32(i)*stepbase {
			interval = i
			break
		}
	}
	//print("interval:",interval)
	percentage := (ratio - stepbase*float32(interval-1)) / stepbase
	//fmt.Println("percentage:", percentage)

	//if (c+1==clases) percentage = 1;//la ultima clase
	return interpolate(colors[interval], colors[interval-1], percentage)
}

func GenerateBreaks(start float32, end float32, count int) []float32 {

	step := (end - start) / float32(count-1)
	var steps []float32

	// remuevo el primero y el ultimo
	for i := 1; i < count-1; i++ {

		steps = append(steps, start+step*float32(i))
	}

	return steps
}

func GenerateBinds(start float32, end float32, count int) []string {

	step := (end - start) / float32(count-1)
	var steps []string

	for i := 0; i < count; i++ {

		steps = append(steps, strconv.Itoa(int(start+step*float32(i))))
	}
	return steps
}

func CalculateHistogram(values []float32, numberOfBins int) ([]int, []float32) {

	max := slices.Max(values)
	min := slices.Min(values)

	//fmt.Println("Min:", min, " Max:", max)

	binRange := (max - min) / float32(numberOfBins)
	//fmt.Println("binRange:", binRange)

	// var numberOfBins = Math.ceil(len / binRange);

	var hist = make([]int, numberOfBins)
	var hedge = make([]float32, numberOfBins)

	for _, x := range values {

		idx := int(math.Floor(float64((x - min) / binRange)))
		//fmt.Println("idx:", idx)
		if idx == numberOfBins {
			idx = numberOfBins - 1
		}
		hist[idx]++
	}

	for i := 0; i < numberOfBins; i++ {
		hedge[i] = min + float32(i)*binRange
	}

	return hist, hedge

}

func TestColors() {
	//func main() {
	start := []int{300, 150, 100}
	end := []int{100, 50, 30}
	ratio := float32(0.5)
	color := interpolate(start, end, ratio)
	fmt.Println("Interpolate:", color)

	rgb := []int{255, 125, 100}
	hex := Rgb2Hex(rgb) // #ff7d64
	fmt.Println("Rgb2Hex:", hex)

	fmt.Println("Hex2Rgb:", Hex2Rgb("#ffaa10"))

	hexList := []string{"#ff0000", "#00ff00", "#0000ff"}
	fmt.Println("HexList2RgbList:", HexList2RgbList(hexList))

	colors := [][]int{{255, 255, 255}, {125, 125, 125}, {0, 0, 0}}
	fmt.Println("multiColor:", multiColor(colors, 0.5)) //output: 125, 125, 125

	fmt.Println("HexMultiColorScale:", HexMultiColorScale(colors, 3)) //
	fmt.Println("RgbMultiColorScale:", RgbMultiColorScale(colors, 3)) //

}

func Marrones() {
	marrones := []string{"#d2b48c", "#d2691e", "#8b4513"}
	hexColors := HexList2RgbList(marrones)
	fmt.Println("HexList2RgbList:", hexColors)

	fmt.Println("Hex2Rgb:", Hex2Rgb("#d2b48c"))
	fmt.Println("Hex2Rgb:", Hex2Rgb("#D2B48C"))

	//[[210, 180, 140], [210, 105, 30], [139, 69, 19]]
	//fmt.Println("HexMultiColorScale:", HexMultiColorScale(hexColors, 3)) //
	//fmt.Println("RgbMultiColorScale:", RgbMultiColorScale(hexColors, 3)) //

}

func TestMath() {

	//fmt.Println("GenerateBreaks:", GenerateBreaks(0, 100, 10))
	//fmt.Println("GenerateBinds:", GenerateBinds(0, 100, 6))

	breaks := GenerateBreaks(0, 100, 10)
	var hedges [100]float32
	for i := 1; i < len(hedges); i++ {
		hedges[i] = float32(i)
	}

	var colorIdx = func(val float32) int {
		i := len(breaks)
		for idx, br := range breaks {

			if val <= br {
				i = idx
				break
			}
		}
		return i
	}

	var colValues [100]int
	for idx, h := range hedges {
		colValues[idx] = colorIdx(h)
	}

	fmt.Println("Breaks:", breaks)
	fmt.Println("Hedges:", hedges)
	fmt.Println("colValues:", colValues)

}

func TestHistogram() {

	var values [100]float32
	for i := 1; i < len(values); i++ {
		values[i] = values[i] + float32(i)
	}
	//var values2 []float32
	//values2 = values

	hist, hedges := CalculateHistogram(values[:], 5)
	fmt.Println("CalculateHistogram", hist, hedges)
}

func main2() {
	//TestColors()
	Marrones()

}
