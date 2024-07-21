package main

import "image/color"

type DriverDefault struct {
	// Height png height in pixel.
	Height int
	// Width Captcha png width in pixel.
	Width int
	// DefaultLen Default number of digits in captcha solution.
	Length int
	// MaxSkew max absolute skew factor of a single digit.
	MaxSkew float64
	// DotCount Number of background circles.
	DotCount int

	BgColor *color.RGBA
}

func (dd *DriverDefault) DrawCaptcha(content string) (imagetor Imagetor, err error) {
	var bgc color.RGBA
	if dd.BgColor != nil {
		bgc = *dd.BgColor
	}
	//else {
	//	bgc = RandLightColor()
	//}
	image := NewImage(dd.Width, dd.Height, bgc)

	//draw content
	err = image.DrawText(content)
	if err != nil {
		return
	}

	return image, nil
}

// GenerateIdQuestionAnswer creates rand id, content and answer
func (dd *DriverDefault) GenerateIdQuestionAnswer() (id, q, a string) {
	return id, a, a
}
