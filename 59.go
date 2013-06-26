package main

import (
    "code.google.com/p/go-tour/pic"
    "image/color"
    "image"
)

type Image struct{}

func (i Image) ColorModel() color.Model{
	return color.RGBAModel;
}

func (i Image) Bounds() image.Rectangle{
	return image.Rect(0,0,200,200)
} 

func (i Image) At(x, y int) color.Color{
		//might want to add another function to have different renders
    return color.RGBA{uint8(x^y),uint8(x^y),uint8(x^y),255}
}

func main() {
    m := Image{}
    pic.ShowImage(m)
}
