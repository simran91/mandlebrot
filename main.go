package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

const (

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////
	xmin   = -2.0 // Where the real axis stats from...
	xmax   = 1.0  // Where the real axis ends...
	ymin   = -1.0 // Where the imaginary axis stars from...
	ymax   = 1.0  // Where the imaginary axis ends...
	width  = 1024 // Width of the image we want...
	height = 1024 // Height of the image we want...

	// Interesting stuff here too...
	// xmin   = -20.0 // Where the real axis stats from...
	// xmax   = 10.0  // Where the real axis ends...
	// ymin   = -10.0 // Where the imaginary axis stars from...
	// ymax   = 10.0  // Where the imaginary axis ends...
	// width  = 1024  // Width of the image we want...
	// height = 1024  // Height of the image we want...

	// Interesting stuff here too...
	// xmin   = -1.502929
	// xmax   = -1.456054
	// ymin   = -0.023437
	// ymax   = 0.023437
	// width  = 1024
	// height = 1024

	// Interesting stuff here too...
	// xmin    = -1.5
	// xmax    = -1.4
	// ymin    = -0.10
	// ymax    = 0.10
	// width   = 1024
	// height  = 1024
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// The name of the generate image file...
	outfile = "mandlebrot-generated.png"

	// maxiters: The number of iterations for when we determine if something has diverged or not...
	// The higer this number is, the more more detailed the result... but the longer it takes to compute...
	maxiters = 1000

	// debug
	debug = false
)

func main() {

	// Open the image file for writing...
	f, err := os.Create(outfile)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new image...
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	// Iterate over each pixel position for the image and fill it in...
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			colour := getPixel(i, j)
			img.SetRGBA(i, j, colour)
		}
	}

	// Save the actual image to the file...
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated Mandlebrot set image and saved file to %v\n", outfile)
	fmt.Printf("\tReal Axis starts at                        \t: %v\n", xmin)
	fmt.Printf("\tReal Axis ends at                          \t: %v\n", xmax)
	fmt.Printf("\tImaginary Axis starts at                   \t: %v\n", ymin)
	fmt.Printf("\tImaginary Axis ends at                     \t: %v\n", ymax)
	fmt.Printf("\tNumber of iterations to check divergence is\t: %v\n", maxiters)
	fmt.Printf("\tImage Resolution is                        \t: %vx%v\n", width, height)
}

// getPixel: returns the colour we should have for the pixel position based on if the mandlebrot formula is converging or diverging at that position...
func getPixel(i, j int) color.RGBA {

	// As i, j are pixel pisitions, we do need to work out what actual axis points they represent on the real(x) and imaginary(y) axis...
	xn := normaliseX(i)
	yn := normaliseY(j)

	// C is the current point (complex constant) we start at
	c := complex(xn, yn)

	// z is always initialised to c (we could have just done z := c, but writing this looked clearer to me for some reason! :)
	z := complex(0, 0) + c

	// some debugging stuff...
	if debug {
		fmt.Printf("Getting pixel %v,%v for c:%v\n", i, j, c)
	}

	// we are going to run the formula z=z*z+c for a max of 'maxiters', if the absolute value at any of the interations gets
	// larger than 4, then we know that the position is diverging, and we take a mod 255 of it (brighten it up if we want) and
	// give it a colour...
	for i := 0; i <= maxiters; i++ {
		z = z*z + c
		if abs := cmplx.Abs(z); abs > 4 {
			mod := int(abs) % 255

			brightness := mod
			if brightness < 70 {
				brightness += 255 - 70
			}

			redValue := 0
			greenValue := 0
			blueValue := 0

			if mod%4 == 1 {
				redValue = brightness
			} else if mod%4 == 2 {
				greenValue = brightness
			} else if mod%4 == 3 {
				blueValue = brightness
			}

			return color.RGBA{uint8(redValue), uint8(greenValue), uint8(blueValue), uint8(255)}
		}
	}

	// If we reached here then we did maxiters and the absolute value did not get > 4, so we return the
	// colour black for this pixel...
	return color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)}
}

// normaliseX: given the pixel point we are at for the width, it returns the value of the point(x) on the real axis...
func normaliseX(i int) float64 {
	return (((xmax - xmin) * float64(i)) / float64(width)) + xmin
}

// normaliseY: given the pixel point we are at for the height, it returns the value of the point(y) on the imaginary axis...
func normaliseY(j int) float64 {
	return (((ymax - ymin) * float64(j)) / float64(height)) + ymin
}
