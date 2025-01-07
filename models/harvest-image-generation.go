package models

import (
	"fmt"
	"gisapi/dto"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"math"
)

func HarvestOperationsImage(body dto.HarvestOperationsImg) image.Image {

	fmt.Println("HarvestOperationsImage:")
	//----------------
	// Bbox in 3857
	//----------------

	bbox, _ := HarvestOperationsBounds3857(body.Hoids)

	Xmin := bbox["xmin"]
	Ymin := bbox["ymin"]
	Xmax := bbox["xmax"]
	Ymax := bbox["ymax"]

	//fmt.Printf("Ymax: %T\n", Ymax)

	//fmt.Println("Bounding Box ---------------------:")
	//fmt.Printf("W: %f, \n", Xmax-Xmin)
	//fmt.Printf("H: %f, \n", Ymax-Ymin)

	//------------------
	// Retrieve points stamps  {x,y,values}
	//------------------

	stamps, _ := HarvestOperationStamps(body.Variable, body.Hoids)

	//fmt.Println("Stamp: ", stamps[0])

	//--------------------------
	// Assign Colors to stamps
	//--------------------------
	coloredStamps := HarvestStampsColorsAssignation(body.Pallete, body.Scale, stamps)

	//fmt.Println(coloredStamps[1])

	//--------------------------
	// Draw Image
	//--------------------------

	// Define pixel_size and NoData value of new raster

	pixel_size := float64(1) // n unit are one pixel
	// Create the target data source
	target_Width := int(math.Abs(Xmax-Xmin) / pixel_size)
	target_Height := int(math.Abs(Ymax-Ymin) / pixel_size)
	//fmt.Println("target_Width:", target_Width)
	//fmt.Println("target_Height:", target_Height)

	img := image.NewRGBA(image.Rect(0, 0, target_Width, target_Height))

	for _, s := range coloredStamps {

		// get the x,y coordinates for the point
		xx := float64(s.X)
		yy := float64(s.Y)
		//print("xx,yy",xx,yy)
		//print( int((xx-Xmin) / pixel_size))

		x := int((xx - Xmin) / pixel_size)
		y := int((Ymax - yy) / pixel_size)

		drawCircle(img, x, y, 3, color.RGBA{uint8(s.Color[0]), uint8(s.Color[1]), uint8(s.Color[2]), 255})
		drawCircle(img, x, y, 2, color.RGBA{uint8(s.Color[0]), uint8(s.Color[1]), uint8(s.Color[2]), 255})

	}

	return img

}

func TestImageCreation() image.Image {

	img := image.NewRGBA(image.Rect(0, 0, 240, 240)) // set image size
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	drawCircle2(img, 120, 120, 20, color.RGBA{255, 0, 0, 255})

	//var img image.Image = m
	return img

}

func drawCircle(img draw.Image, x0, y0, r int, c color.Color) {
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)

	for x > y {
		img.Set(x0+x, y0+y, c)
		img.Set(x0+y, y0+x, c)
		img.Set(x0-y, y0+x, c)
		img.Set(x0-x, y0+y, c)
		img.Set(x0-x, y0-y, c)
		img.Set(x0-y, y0-x, c)
		img.Set(x0+y, y0-x, c)
		img.Set(x0+x, y0-y, c)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
}

func drawCircle2(img draw.Image, x, y int, radius int, fill color.Color) error {
	// Algorithm taken from
	// http://en.wikipedia.org/wiki/Midpoint_circle_algorithm
	// No need to check the radius is in bounds because you can only
	// create circles using NewCircle() which guarantees it is within
	// bounds. But the x, y might be outside the image so we check.
	//if err := checkBounds(img, x, y); err != nil {
	//	return err
	//}

	x0, y0 := x, y
	f := 1 - radius
	ddF_x, ddF_y := 1, -2*radius
	x, y = 0, radius

	img.Set(x0, y0+radius, fill)
	img.Set(x0, y0-radius, fill)
	img.Set(x0+radius, y0, fill)
	img.Set(x0-radius, y0, fill)

	for x < y {
		if f >= 0 {
			y--
			ddF_y += 2
			f += ddF_y
		}
		x++
		ddF_x += 2
		f += ddF_x
		img.Set(x0+x, y0+y, fill)
		img.Set(x0-x, y0+y, fill)
		img.Set(x0+x, y0-y, fill)
		img.Set(x0-x, y0-y, fill)
		img.Set(x0+y, y0+x, fill)
		img.Set(x0-y, y0+x, fill)
		img.Set(x0+y, y0-x, fill)
		img.Set(x0-y, y0-x, fill)
	}
	return nil
}

/*
func checkBounds(img image.Image, x, y int) error {
	if !image.Rect(x, y, x, y).In(img.Bounds()) {

		fmt.Errorf("Error: point (%d, %d) is outside the image\n", x, y)
		return nil
	}
	return nil
}
*/
