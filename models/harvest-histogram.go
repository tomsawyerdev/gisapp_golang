package models

import (
	//"fmt"
	grd "gisapi/colors"
	"gonum.org/v1/gonum/stat"
	"math"
	"slices"
	"strconv"
)

//from colors import colorsGradients as grd

// hexpallete: is a hex string list ['#FF0000', '#FFD700', '#006455']
// scale:CLO, I2,I3,I4,I5, Q2,Q3,Q4, K2,K3,K4
// argvalues: float array
// Return:  {"status":"success","bins": bins, "colors": colValues, "freq" : hist.tolist(), "pallete": colorscale, "labels": labels}

//home/userone/Documentos/AppReact/flaskAPI/harvest/harvestHistogram.py

func HarvestGenerateHistogram(hexPallete []string, scale string, values []float32) map[string]any {

	//fmt.Println("GenerateHistogram:", hexPallete, scale, len(values))

	histogram, hedges := grd.CalculateHistogram(values, 100)

	vmax := slices.Max(values)
	vmin := slices.Min(values)

	//fmt.Println("Min:", vmin, " Max:", vmax)

	var colorscale []string
	var breaks []float32

	// pallete is a hex string list ['#FF0000', '#FFD700', '#006455']
	rgbPallete := grd.HexList2RgbList(hexPallete)

	//fmt.Println("rgbPallete:", rgbPallete)

	//CLO, I2,I3,I4,I5, Q3,Q4,Q5,IQR,STD
	if scale == "CLO" {
		breaks = grd.GenerateBreaks(vmin, vmax, 10)
		colorscale = grd.HexMultiColorScale(rgbPallete, 9)
	}
	if scale == "I2" {
		breaks = grd.GenerateBreaks(vmin, vmax, 3)
		colorscale = grd.HexMultiColorScale(rgbPallete, 2)
	}

	if scale == "I3" {
		breaks = grd.GenerateBreaks(vmin, vmax, 4)
		colorscale = grd.HexMultiColorScale(rgbPallete, 3)
	}

	if scale == "I4" {
		breaks = grd.GenerateBreaks(vmin, vmax, 5)
		colorscale = grd.HexMultiColorScale(rgbPallete, 4)
	}
	if scale == "I5" {
		breaks = grd.GenerateBreaks(vmin, vmax, 6)
		colorscale = grd.HexMultiColorScale(rgbPallete, 5)
	}
	//Q3,Q4,Q5,IQR,STD

	slices.Sort(values)
	values64 := make([]float64, len(values))
	for i, b := range values {
		values64[i] = float64(b)

	}

	if scale == "Q3" {

		Q33 := float32(stat.Quantile(0.33, stat.Empirical, values64, nil))
		Q66 := float32(stat.Quantile(0.66, stat.Empirical, values64, nil))
		breaks = ([]float32{Q33, Q66})
		colorscale = grd.HexMultiColorScale(rgbPallete, 3)
	}
	if scale == "Q4" {
		Q25 := float32(stat.Quantile(0.25, stat.Empirical, values64, nil))
		Q50 := float32(stat.Quantile(0.50, stat.Empirical, values64, nil))
		Q75 := float32(stat.Quantile(0.75, stat.Empirical, values64, nil))
		breaks = []float32{Q25, Q50, Q75}
		colorscale = grd.HexMultiColorScale(rgbPallete, 4)
	}
	if scale == "Q5" {
		Q20 := float32(stat.Quantile(0.20, stat.Empirical, values64, nil))
		Q40 := float32(stat.Quantile(0.40, stat.Empirical, values64, nil))
		Q60 := float32(stat.Quantile(0.60, stat.Empirical, values64, nil))
		Q80 := float32(stat.Quantile(0.80, stat.Empirical, values64, nil))
		breaks = []float32{Q20, Q40, Q60, Q80}

		colorscale = grd.HexMultiColorScale(rgbPallete, 5)
	}
	//IQR,STD
	if scale == "IQR" {
		Q1 := float32(stat.Quantile(0.25, stat.Empirical, values64, nil))
		//Q50 := stat.Quantile(0.50, stat.Empirical, values, nil)
		Q3 := float32(stat.Quantile(0.75, stat.Empirical, values64, nil))
		IQR := Q3 - Q1
		breaks = []float32{Q1 - 1.5*IQR, Q1, Q3, Q3 + 1.5*IQR}
		colorscale = grd.HexMultiColorScale(rgbPallete, 5)
	}

	if scale == "STD" {
		m := float32(stat.Mean(values64, nil))
		variance := stat.Variance(values64, nil)
		sd := float32(math.Sqrt(variance))
		breaks = []float32{m - sd, m, m + sd}
		colorscale = grd.HexMultiColorScale(rgbPallete, 4)
	}

	//fmt.Println("Breaks:", breaks)

	// Bins generation --------------------------------
	// 6 bins between min and max
	bins := grd.GenerateBinds(vmin, vmax, 6) //[]string

	//fmt.Println("Bins:", bins)

	// Labels generation --------------------------------
	var labels []string
	for _, v := range breaks {
		labels = append(labels, strconv.Itoa(int(v)))
	}

	//fmt.Println("Labels:", labels)

	//-----------------------------------
	// Color assignation
	// find the hedge colors from the hedge value
	//-----------------------------------

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

	var colValues [100]string
	for idx, h := range hedges {
		colValues[idx] = colorscale[colorIdx(h)]
	}
	//fmt.Println("histogram:", histogram)
	//fmt.Println("colValues:", colValues)

	//-----------------------------------------------
	//var result map[string]any // devuelve nil esta declarado pero no inicializado
	// Alternativa conmake

	//result := make(map[string]interface{})
	//result["status"] = "success"
	result := map[string]any{"status": "success"}
	result["bins"] = bins
	result["colors"] = colValues
	result["freq"] = histogram
	result["pallete"] = colorscale
	result["labels"] = labels
	//fmt.Println("result:", result)

	return result

}
