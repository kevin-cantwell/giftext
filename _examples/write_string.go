package main

import (
	"image/color"
	"image/gif"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
	"github.com/kevin-cantwell/giftext"
)

func main() {
	fontfile, err := os.Open("Arial Bold Italic.ttf")
	if err != nil {
		panic(err)
	}
	defer fontfile.Close()
	body, err := ioutil.ReadAll(fontfile)
	if err != nil {
		panic(err)
	}
	font, err := freetype.ParseFont(body)
	if err != nil {
		panic(err)
	}
	giffile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer giffile.Close()
	giff, err := gif.DecodeAll(giffile)
	if err != nil {
		panic(err)
	}
	config := giftext.Config{
		Font:     font,
		FontSize: 30,
		Color:    color.RGBA{0xff, 0xff, uint8(float32(0xff) * 0.2), 0x00},
	}
	w := giftext.NewWriter(giff, config)
	if err := w.WriteString(text, giff.Image[0].Bounds().Min); err != nil {
		panic(err)
	}
	outfile, err := os.OpenFile("output.gif", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	if err := gif.EncodeAll(outfile, giff); err != nil {
		panic(err)
	}
}

var (
	text = "0abcdefghijklmnopqrstuvwxyz\n" +
		"1abcdefghijklmnopqrstuvwxyz\n" +
		"2abcdefghijklmnopqrstuvwxyz\n" +
		"3abcdefghijklmnopqrstuvwxyz\n" +
		"4abcdefghijklmnopqrstuvwxyz\n" +
		"5abcdefghijklmnopqrstuvwxyz\n" +
		"6abcdefghijklmnopqrstuvwxyz\n" +
		"7abcdefghijklmnopqrstuvwxyz\n" +
		"8abcdefghijklmnopqrstuvwxyz\n" +
		"9abcdefghijklmnopqrstuvwxyz"
)
