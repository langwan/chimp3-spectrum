package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math"
	"spectrum/app"
)

type Led2 struct {
	Fraction       float64
	SpacingFactorX float64
	MinDecibels    float64
	MaxDecibels    float64
	LedCellHeight  float64
	SpacingFactorY float64
	Color1         color.RGBA
	Color2         color.RGBA
}

func init() {
	whiteImage.Fill(color.White)

}

func (l Led2) Update() error {
	return nil
}

var ledCache2 []float64
var gCapCache [2][]float64

func (l Led2) DrawA(screen *ebiten.Image) {

	sw, sh := screen.Size()
	channel := 0
	if len(app.FftSamples[channel]) == 0 {
		return
	}

	barNum := 128
	var bars []float64
	step := int(math.Floor(float64(len(app.FftSamples[channel])) / float64(barNum*2)))

	for i := 0; i < barNum; i++ {
		bars = append(bars, app.FftSamples[channel][i*step+channel])
	}

	if len(ledCache) == 0 {
		ledCache = make([]float64, len(bars))
		for i := range bars {
			ledCache[i] = bars[i]
		}
	}

	xLeds := float64(len(bars))
	ledCellWidth := float64(sw) / xLeds

	xPadding := ledCellWidth * l.SpacingFactorX
	ledWidth := ledCellWidth - 2.0*xPadding

	yPadding := l.LedCellHeight * l.SpacingFactorY
	ledHeight := l.LedCellHeight - 2.0*yPadding

	yLeds := math.Round(float64(sh) / l.LedCellHeight)

	//for i := range LinearGradientColors {
	//	LinearGradientColors[i] = LinearGradientColor{
	//		C1: color.RGBA{},
	//		C2: color.RGBA{},
	//	}
	//}

	if len(gCapCache[channel]) == 0 {
		gCapCache[channel] = make([]float64, len(bars))
		for i, value := range bars {
			gCapCache[channel][i] = value
		}
	}

	for i := 0.0; i < xLeds; i++ {
		for j := 0.0; j < yLeds; j++ {
			ebitenutil.DrawRect(screen, i*(ledCellWidth)+xPadding, float64(sh)-(j+1.0)*(l.LedCellHeight)+yPadding, ledWidth, ledHeight, color.RGBA{
				R: 0x11,
				G: 0x11,
				B: 0x11,
				A: 0xff,
			})
		}
	}
	dr := (float64(l.Color2.R) - float64(l.Color1.R)) / yLeds
	dg := (float64(l.Color2.G) - float64(l.Color1.B)) / yLeds
	db := (float64(l.Color2.B) - float64(l.Color1.G)) / yLeds
	da := (float64(l.Color2.A) - float64(l.Color1.A)) / yLeds
	for i := 0; i < int(xLeds); i++ {
		val := bars[int(i)]
		ledCache[i] = ledCache[i]*smoothingTime + (1-smoothingTime)*val
		yLedNumber := math.Round(ledCache[i] / 255.0 * yLeds)
		if ledCache[i] < gCapCache[channel][i] {
			n := val - gCapCache[channel][i]
			gCapCache[channel][i] = gCapCache[channel][i]*(smoothingTimeCap) + (1-smoothingTimeCap)*(math.Max(n, -1))
			yCapNumber := math.Round(gCapCache[channel][i] / 255.0 * yLeds)
			DrawBar(screen, float64(i)*(ledCellWidth)+xPadding, float64(sh)-(yCapNumber+1.0)*(l.LedCellHeight)+yPadding, ledWidth, ledHeight, color.RGBA{
				R: 0x3a,
				G: 0x84,
				B: 0xf2,
				A: 0xee,
			}, color.RGBA{
				R: 0x3a,
				G: 0x84,
				B: 0xf2,
				A: 0xee,
			})
		} else {
			gCapCache[channel][i] = val
			yCapNumber := math.Round(gCapCache[channel][i] / 255.0 * yLeds)
			DrawBar(screen, float64(i)*(ledCellWidth)+xPadding, float64(sh)-(yCapNumber+1.0)*(l.LedCellHeight)+yPadding, ledWidth, ledHeight, color.RGBA{
				R: 0x3a,
				G: 0x84,
				B: 0xf2,
				A: 0xee,
			}, color.RGBA{
				R: 0x3a,
				G: 0x84,
				B: 0xf2,
				A: 0xee,
			})
		}

		for j := 0.0; j < yLedNumber; j++ {
			DrawBar(screen, float64(i)*(ledCellWidth)+xPadding, float64(sh)-(j+1.0)*(l.LedCellHeight)+yPadding, ledWidth, ledHeight, color.RGBA{
				R: uint8(float64(l.Color2.R) - (j+1)*dr),
				G: uint8(float64(l.Color2.G) - (j+1)*dg),
				B: uint8(float64(l.Color2.B) - (j+1)*db),
				A: uint8(float64(l.Color2.A) - (j+1)*da),
			}, color.RGBA{
				R: uint8(float64(l.Color2.R) - (j)*dr),
				G: uint8(float64(l.Color2.G) - (j)*dg),
				B: uint8(float64(l.Color2.B) - (j)*db),
				A: uint8(float64(l.Color2.A) - (j)*da),
			})
		}
	}
}

func (l Led2) DrawB(screen *ebiten.Image) {

	sw, sh := screen.Size()
	channel := 1
	if len(app.FftSamples[channel]) == 0 {
		return
	}

	barNum := 128
	var bars []float64
	step := int(math.Floor(float64(len(app.FftSamples[channel])) / float64(float64(barNum)*1.5)))

	for i := 0; i < barNum; i++ {
		bars = append(bars, app.FftSamples[channel][i*step+channel]*0.5)
	}

	if len(ledCache2) == 0 {
		ledCache2 = make([]float64, len(bars))
		for i := range bars {
			ledCache2[i] = bars[i]
		}
	}

	xLeds := float64(len(bars))
	ledCellWidth := float64(sw) / xLeds

	xPadding := ledCellWidth * l.SpacingFactorX
	ledWidth := ledCellWidth - 2.0*xPadding

	yPadding := l.LedCellHeight * l.SpacingFactorY
	ledHeight := l.LedCellHeight - 2.0*yPadding

	yLeds := math.Round(float64(sh) / l.LedCellHeight)

	c1 := color.RGBA{
		R: 0xf0,
		G: 0x2a,
		B: 0x20,
		A: 0xff,
	}
	c2 := color.RGBA{
		R: 0xf2,
		G: 0xb8,
		B: 0x12,
		A: 0xff,
	}
	dr := (float64(c2.R) - float64(c1.R)) / yLeds
	dg := (float64(c2.G) - float64(c1.B)) / yLeds
	db := (float64(c2.B) - float64(c1.G)) / yLeds
	da := (float64(c2.A) - float64(c1.A)) / yLeds
	for i := 0; i < int(xLeds); i++ {
		val := bars[int(i)]
		ledCache2[i] = ledCache2[i]*smoothingTime + (1-smoothingTime)*val
		yLedNumber := math.Round(ledCache2[i] / 255.0 * yLeds)
		//	yLedNumber = yLeds

		for j := 0.0; j < yLedNumber; j++ {
			DrawBar(screen, float64(i)*(ledCellWidth)+xPadding, float64(sh)-(j+1.0)*(l.LedCellHeight)+yPadding, ledWidth, ledHeight, color.RGBA{
				R: uint8(float64(c2.R) - (j+1)*dr),
				G: uint8(float64(c2.G) - (j+1)*dg),
				B: uint8(float64(c2.B) - (j+1)*db),
				A: uint8(float64(c2.A) - (j+1)*da),
			}, color.RGBA{
				R: uint8(float64(c2.R) - (j)*dr),
				G: uint8(float64(c2.G) - (j)*dg),
				B: uint8(float64(c2.B) - (j)*db),
				A: uint8(float64(c2.A) - (j)*da),
			})
		}
	}
}

func (l Led2) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 0, 0
}
