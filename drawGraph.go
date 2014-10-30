package drawGraph

import (
	"bufio"
	"code.google.com/p/draw2d/draw2d"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func NewImage(width, height int) (image.Image, draw2d.GraphicContext) {
	im := image.NewRGBA(image.Rect(0, 0, width, height))
	gC := draw2d.NewGraphicContext(im)

	gC.SetFillColor(image.White)
	gC.SetStrokeColor(image.Black)

	return im, gC
}

func SaveImage(name string, im image.Image) {
	name = name + ".png"
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	outWriter := bufio.NewWriter(file)
	err = png.Encode(outWriter, im)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = outWriter.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Created " + name)
}
