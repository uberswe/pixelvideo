package main

import (
	"bytes"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/icza/mjpeg"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"image/jpeg"
)

var (
	resX = 720
	resY = 480
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Window",
		Bounds: pixel.R(0, 0, float64(resX), float64(resY)),
		// Set this to false if you want to see the window
		Invisible: true,
	}
	_, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	canvas := pixelgl.NewCanvas(pixel.R(0, 0, float64(resX), float64(resY)))

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(50, float64(resY-50)), basicAtlas)

	basicTxt.LineHeight = basicAtlas.LineHeight() * 1.5

	fmt.Fprintln(basicTxt, "This is an example text")
	fmt.Fprintln(basicTxt, "This text supports multiple lines")

	// Here i specify the total frames
	var (
		i           = 1
		totalFrames = 600
	)

	// Here the fps is set, take totalframes/fps to get the duration with my values the video duration is 10 seconds
	aw, err := mjpeg.New("test.avi", int32(resX), int32(resY), 60)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Generating frames\n")
	for i <= totalFrames {
		// Look at the pixel documentation on how to do animations
		canvas.Clear(colornames.Black)

		fmt.Fprintf(basicTxt, "%d ", i)
		if i%30 == 0 && i > 1 {
			fmt.Fprint(basicTxt, "\n")
		}
		basicTxt.Draw(canvas, pixel.IM.Scaled(basicTxt.Orig, 1))

		img := pixel.PictureDataFromPicture(canvas)
		buf := &bytes.Buffer{}
		err = jpeg.Encode(buf, img.Image(), nil)
		if err != nil {
			panic(err)
		}
		err = aw.AddFrame(buf.Bytes())
		if err != nil {
			panic(err)
		}

		i++
	}
	fmt.Printf("Saving video\n")
	// This saves the video file
	err = aw.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	pixelgl.Run(run)
}
