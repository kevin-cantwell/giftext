package giftext

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"strings"
	"sync"

	"golang.org/x/image/font"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type Writer struct {
	giff   *gif.GIF
	config Config
}

type Config struct {
	Font     *truetype.Font
	FontSize float64
	Color    color.Color
}

func NewWriter(giff *gif.GIF, config Config) *Writer {
	return &Writer{
		giff:   giff,
		config: config,
	}
}

func (w *Writer) WriteString(s string, sp image.Point) error {
	overlay := image.NewPaletted(w.giff.Image[0].Bounds(), w.giff.Image[0].Palette)
	// NOTE: This _might_ not actually draw transparency if that color is not in the palette
	draw.Draw(overlay, overlay.Bounds(), image.Transparent, overlay.Bounds().Min, draw.Src)
	ftc := freetype.NewContext()
	ftc.SetHinting(font.HintingNone)
	ftc.SetDPI(72)
	ftc.SetFont(w.config.Font)
	ftc.SetClip(overlay.Bounds())
	ftc.SetSrc(image.NewUniform(w.config.Color))
	ftc.SetFontSize(w.config.FontSize)
	ftc.SetDst(overlay)

	lines := strings.Split(s, "\n")
	fsp := freetype.Pt(sp.X, sp.Y+int(w.config.FontSize))
	fpt := fsp
	for _, line := range lines {
		var err error
		fpt, err = ftc.DrawString(line, fpt)
		if err != nil {
			return err
		}
		fpt.X = fsp.X
		fpt.Y += ftc.PointToFixed(w.config.FontSize)
	}
	var wg sync.WaitGroup
	for _, frame := range w.giff.Image {
		wg.Add(1)
		go func(frame draw.Image) {
			defer wg.Done()
			draw.Draw(frame, overlay.Bounds(), overlay, overlay.Bounds().Min, draw.Over)
		}(frame)
	}
	wg.Wait()
	return nil
}
