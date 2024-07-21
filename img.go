package main

import (
	"errors"
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"io"
	"math/rand"
)

type Img struct {
	nrgba     *image.NRGBA
	bgColor   color.Color
	numWidth  int
	numHeight int
	doSize    int
}

func (i *Img) DrawText(text string) error {
	img := &Img{}
	c := freetype.NewContext()
	c.SetDPI(72.0)
	c.SetClip(img.nrgba.Bounds())
	c.SetDst(img.nrgba)
	c.SetHinting(font.HintingFull)

	if len(text) == 0 {
		return errors.New("text must not be empty, there is nothing to draw")
	}

	fontWidth := img.numWidth / len(text)

	for i, s := range text {
		fontSize := img.numHeight * (rand.Intn(7) + 7) / 16
		//c.SetSrc(image.NewUniform(RandDeepColor()))
		c.SetFontSize(float64(fontSize))
		//c.SetFont(randFontFrom(fonts))
		x := fontWidth*i + fontWidth/fontSize
		y := img.numHeight/2 + fontSize/2 - rand.Intn(img.numHeight/16*3)
		pt := freetype.Pt(x, y)
		if _, err := c.DrawString(string(s), pt); err != nil {
			return err
		}
	}
	return nil

}
func (i *Img) WriteTo(w io.Writer) (n int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *Img) EncodeB64string() string {
	//TODO implement me
	panic("implement me")
}

func NewImage(w, h int, bgColor color.RGBA) *Img {
	d := Img{numWidth: w, numHeight: h}
	m := image.NewNRGBA(image.Rect(0, 0, w, h))
	draw.Draw(m, m.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)
	d.nrgba = m
	return &d

}
