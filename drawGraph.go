package drawGraph

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/fogleman/gg"
)

/*
Package to draw 2D graphs (vertecies and lines)
*/

type DrawObject struct {
	im                  *image.RGBA
	gC                  *gg.Context
	maxHeight, maxWidth float64
}

var (
	radius    = 5.0
	lineWidth = 3.0
)

func Instance() *DrawObject {
	return new(DrawObject)
}

// Overrides current image with new image
func NewImage(width, height int) *DrawObject {
	i := new(DrawObject)
	i.maxHeight = float64(height)
	i.maxWidth = float64(width)
	//i.im = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	gg.NewContext(width, height)
	//i.gC = draw2dimg.NewGraphicContext(i.im)

	white := color.NRGBA{0xff, 0xff, 0xff, 0xff}
	i.gC.SetColor(white)
	i.gC.DrawRectangle(0, 0, float64(width), float64(height))
	// i.gC.ClearRect(0, 0, width, height)
	i.gC.Fill()
	green := color.NRGBA{0x00, 0x66, 0x66, 0xff}
	i.gC.SetColor(green)
	i.gC.SetLineWidth(lineWidth)
	return i
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

func (i *DrawObject) AddPoint(x, y float64) error {
	if x < 0 || y < 0 {
		return errors.New("Coordinate < 0 !!")
	}
	if x > i.maxWidth || y > i.maxHeight {
		return errors.New("Coordinate > maximum")
	}
	//fmt.Println("Add point ", x, " ", y)
	// draw2d.Circle(i.gC, x, i.maxHeight-y, radius)
	i.gC.DrawCircle(x, i.maxHeight-y, radius)
	i.gC.Fill()
	return nil
}

func (i *DrawObject) AddLine(x1, y1, x2, y2 float64) error {
	//fmt.Println("Line from: ", x1, " ", y1, " to ", x2, " ", y2)
	if x1 < 0 || y1 < 0 || x2 < 0 || y2 < 0 {
		return errors.New("Coordinate < 0 !!")
	}
	if x1 > i.maxWidth || x2 > i.maxWidth || y1 > i.maxHeight || y2 > i.maxHeight {
		return errors.New("Coordinate > maximum")
	}
	i.gC.MoveTo(x1, i.maxHeight-y1)
	i.gC.LineTo(x2, i.maxHeight-y2)
	i.gC.Stroke()
	return nil
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
