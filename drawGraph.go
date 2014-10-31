package drawGraph

import (
	"bufio"
	"code.google.com/p/draw2d/draw2d"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

/*
Package to draw 2D graphs (vertecies and lines)
*/

type DrawObject struct {
	im *image.RGBA
	gC draw2d.GraphicContext
}

var (
	instance *DrawObject = nil
	radius               = 5.0
)

func Instance() *DrawObject {
	if instance == nil {
		instance = new(DrawObject)
		return instance
	}
	return instance
}

// Overrides current image with new image
func (i *DrawObject) NewImage(width, height float64) {
	i.im = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	i.gC = draw2d.NewGraphicContext(i.im)

	white := color.NRGBA{0xff, 0xff, 0xff, 0xff}
	i.gC.SetFillColor(white)
	i.gC.SetStrokeColor(image.Black)
	i.gC.MoveTo(0.0, 0.0)
	i.gC.LineTo(width, 0.0)
	i.gC.Stroke()
	i.gC.MoveTo(width, 0.0)
	i.gC.LineTo(width, height)
	i.gC.Stroke()
	i.gC.MoveTo(width, height)
	i.gC.LineTo(0, height)
	i.gC.Stroke()
	i.gC.MoveTo(0.0, height)
	i.gC.LineTo(0.0, 0.0)
	i.gC.Stroke()
	i.gC.Fill()
	green := color.NRGBA{0x00, 0x66, 0x66, 0xff}
	i.gC.SetFillColor(green)
}

func (i *DrawObject) SaveImage(name string) error {
	if i.im == nil {
		return errors.New("Image not available, not initialized!\n")
	}
	name = name + ".png"
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	outWriter := bufio.NewWriter(file)
	err = png.Encode(outWriter, i.im)
	if err != nil {
		log.Println(err)
		return err
	}
	err = outWriter.Flush()
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Created " + name)
	return nil
}

func (i *DrawObject) AddPoint(x, y float64) {
	draw2d.Circle(i.gC, x, y, radius)
	i.gC.Stroke()
}

func (i *DrawObject) PrintSomething() {
	i.gC.MoveTo(10.0, 10.0)
	i.gC.LineTo(100.0, 10.0)
	i.gC.Stroke()
	err := i.SaveImage("TestPath")
	if err != nil {
		log.Println(err)
	}
}
