package led

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	"math"
	"spectrum/app"
)

const (
	barNumbers       = 64
	spacingFactorX   = 0.2
	spacingFactorY   = 0.2
	ledCellHeight    = 12.0
	smoothingTime    = 0.7
	smoothingTimeCap = 0.9
)

type Scene struct {
	Fraction       float64
	SpacingFactorX float64
	MinDecibels    float64
	MaxDecibels    float64
	LedCellHeight  float64
	SpacingFactorY float64
	Color1         color.RGBA
	Color2         color.RGBA
}

func (s Scene) Update() error {
	return nil
}
func (s Scene) Draw(screen *ebiten.Image) {

	s.DrawLed(screen, 0, color.RGBA{
		R: 0xff,
		G: 0x56,
		B: 0x00,
		A: 0xff,
	}, color.RGBA{
		R: 0x22,
		G: 0xff,
		B: 0x00,
		A: 0xff,
	}, 1, 2)

	s.DrawLed(screen, 1, color.RGBA{
		R: 0xf0,
		G: 0x2a,
		B: 0x20,
		A: 0xff,
	},
		color.RGBA{
			R: 0xf2,
			G: 0xb8,
			B: 0x12,
			A: 0xff,
		}, 0.5, 1.5)
}

var spectrumCache [2][]float64
var capCache [2][]float64

func (s Scene) DrawLed(screen *ebiten.Image, channel int, c1 color.RGBA, c2 color.RGBA, valueScale, frequencyScale float64) {

	spectrumLen := len(app.FrequencyFrame.Spectrum[channel])

	if spectrumLen == 0 {
		return
	}

	sw, sh := screen.Size()

	var bars []float64

	step := int(math.Floor(float64(spectrumLen) / (float64(barNumbers) * frequencyScale)))

	for i := 0; i < barNumbers; i++ {
		bars = append(bars, app.FrequencyFrame.Spectrum[channel][i*step+channel]*valueScale)
	}

	if len(spectrumCache[channel]) == 0 {
		spectrumCache[channel] = make([]float64, len(bars))
		for i := range bars {
			spectrumCache[channel][i] = bars[i]
		}
	}

	xLeds := float64(len(bars))
	ledCellWidth := float64(sw) / xLeds

	xPadding := ledCellWidth * spacingFactorX
	ledWidth := ledCellWidth - 2.0*xPadding

	yPadding := ledCellHeight * spacingFactorY
	ledHeight := ledCellHeight - 2.0*yPadding

	yLeds := math.Round(float64(sh) / ledCellHeight)

	//for i := range LinearGradientColors {
	//	LinearGradientColors[i] = LinearGradientColor{
	//		C1: color.RGBA{},
	//		C2: color.RGBA{},
	//	}
	//}

	if len(capCache[channel]) == 0 {
		capCache[channel] = make([]float64, len(bars))
		for i, value := range bars {
			capCache[channel][i] = value
		}
	}

	if channel == 0 {
		for i := 0.0; i < xLeds; i++ {
			for j := 0.0; j < yLeds; j++ {
				ebitenutil.DrawRect(screen, i*(ledCellWidth)+xPadding, float64(sh)-(j+1.0)*(ledCellHeight)+yPadding, ledWidth, ledHeight, color.RGBA{
					R: 0x11,
					G: 0x11,
					B: 0x11,
					A: 0xff,
				})
			}
		}
	}

	dr := (float64(c2.R) - float64(c1.R)) / yLeds
	dg := (float64(c2.G) - float64(c1.B)) / yLeds
	db := (float64(c2.B) - float64(c1.G)) / yLeds
	da := (float64(c2.A) - float64(c1.A)) / yLeds
	for i := 0; i < int(xLeds); i++ {
		val := bars[i]
		spectrumCache[channel][i] = spectrumCache[channel][i]*smoothingTime + (1-smoothingTime)*val
		yLedNumber := math.Round(spectrumCache[channel][i] / 255.0 * yLeds)
		if spectrumCache[channel][i] < capCache[channel][i] {
			n := val - capCache[channel][i]
			capCache[channel][i] = capCache[channel][i]*(smoothingTimeCap) + (1-smoothingTimeCap)*(math.Max(n, -1))
			yCapNumber := math.Round(capCache[channel][i] / 255.0 * yLeds)
			DrawBar(screen, float64(i)*(ledCellWidth)+xPadding, float64(sh)-(yCapNumber+1.0)*(ledCellHeight)+yPadding, ledWidth, ledHeight, color.RGBA{
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
			capCache[channel][i] = val
			//yCapNumber := math.Round(capCache[channel][i] / 255.0 * yLeds)
			//DrawBar(screen, float64(i)*(ledCellWidth)+xPadding, float64(sh)-(yCapNumber+1.0)*(ledCellHeight)+yPadding, ledWidth, ledHeight, color.RGBA{
			//	R: 0x3a,
			//	G: 0x84,
			//	B: 0xf2,
			//	A: 0xee,
			//}, color.RGBA{
			//	R: 0x3a,
			//	G: 0x84,
			//	B: 0xf2,
			//	A: 0xee,
			//})
		}

		for j := 0.0; j < yLedNumber; j++ {
			DrawBar(screen, float64(i)*(ledCellWidth)+xPadding, float64(sh)-(j+1.0)*(ledCellHeight)+yPadding, ledWidth, ledHeight, color.RGBA{
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

//func (s Scene) DrawB(screen *ebiten.Image) {
//
//	sw, sh := screen.Size()
//	channel := 1
//	if len(app.FftSamples[channel]) == 0 {
//		return
//	}
//
//	barNum := 128
//	var bars []float64
//	step := int(math.Floor(float64(len(app.FftSamples[channel])) / float64(float64(barNum)*1.5)))
//
//	for i := 0; i < barNum; i++ {
//		bars = append(bars, app.FftSamples[channel][i*step+channel]*0.5)
//	}
//
//	if len(ledCache2) == 0 {
//		ledCache2 = make([]float64, len(bars))
//		for i := range bars {
//			ledCache2[i] = bars[i]
//		}
//	}
//
//	xLeds := float64(len(bars))
//	ledCellWidth := float64(sw) / xLeds
//
//	xPadding := ledCellWidth * l.SpacingFactorX
//	ledWidth := ledCellWidth - 2.0*xPadding
//
//	yPadding := l.LedCellHeight * l.SpacingFactorY
//	ledHeight := l.LedCellHeight - 2.0*yPadding
//
//	yLeds := math.Round(float64(sh) / l.LedCellHeight)
//
//	c1 := color.RGBA{
//		R: 0xf0,
//		G: 0x2a,
//		B: 0x20,
//		A: 0xff,
//	}
//	c2 := color.RGBA{
//		R: 0xf2,
//		G: 0xb8,
//		B: 0x12,
//		A: 0xff,
//	}
//	dr := (float64(c2.R) - float64(c1.R)) / yLeds
//	dg := (float64(c2.G) - float64(c1.B)) / yLeds
//	db := (float64(c2.B) - float64(c1.G)) / yLeds
//	da := (float64(c2.A) - float64(c1.A)) / yLeds
//	for i := 0; i < int(xLeds); i++ {
//		val := bars[int(i)]
//		ledCache2[i] = ledCache2[i]*smoothingTime + (1-smoothingTime)*val
//		yLedNumber := math.Round(ledCache2[i] / 255.0 * yLeds)
//		//	yLedNumber = yLeds
//
//		for j := 0.0; j < yLedNumber; j++ {
//			DrawBar(screen, float64(i)*(ledCellWidth)+xPadding, float64(sh)-(j+1.0)*(l.LedCellHeight)+yPadding, ledWidth, ledHeight, color.RGBA{
//				R: uint8(float64(c2.R) - (j+1)*dr),
//				G: uint8(float64(c2.G) - (j+1)*dg),
//				B: uint8(float64(c2.B) - (j+1)*db),
//				A: uint8(float64(c2.A) - (j+1)*da),
//			}, color.RGBA{
//				R: uint8(float64(c2.R) - (j)*dr),
//				G: uint8(float64(c2.G) - (j)*dg),
//				B: uint8(float64(c2.B) - (j)*db),
//				A: uint8(float64(c2.A) - (j)*da),
//			})
//		}
//	}
//}

func (s Scene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 0, 0
}

func DrawBar(dst *ebiten.Image, x, y, width, height float64, c1 color.RGBA, c2 color.RGBA) {
	var vertices = make([]ebiten.Vertex, 4)
	vertices[0] = ebiten.Vertex{
		DstX:   float32(x),
		DstY:   float32(y),
		SrcX:   0,
		SrcY:   0,
		ColorR: float32(c1.R) / float32(math.MaxUint8),
		ColorG: float32(c1.G) / float32(math.MaxUint8),
		ColorB: float32(c1.B) / float32(math.MaxUint8),
		ColorA: float32(c1.A) / float32(math.MaxUint8),
	}

	vertices[1] = ebiten.Vertex{
		DstX:   float32(x + width),
		DstY:   float32(y),
		SrcX:   0,
		SrcY:   0,
		ColorR: float32(c1.R) / float32(math.MaxUint8),
		ColorG: float32(c1.G) / float32(math.MaxUint8),
		ColorB: float32(c1.B) / float32(math.MaxUint8),
		ColorA: float32(c1.A) / float32(math.MaxUint8),
	}

	vertices[2] = ebiten.Vertex{
		DstX:   float32(x + width),
		DstY:   float32(y + height),
		SrcX:   0,
		SrcY:   0,
		ColorR: float32(c2.R) / float32(math.MaxUint8),
		ColorG: float32(c2.G) / float32(math.MaxUint8),
		ColorB: float32(c2.B) / float32(math.MaxUint8),
		ColorA: float32(c2.A) / float32(math.MaxUint8),
	}

	vertices[3] = ebiten.Vertex{
		DstX:   float32(x),
		DstY:   float32(y + height),
		SrcX:   0,
		SrcY:   0,
		ColorR: float32(c2.R) / float32(math.MaxUint8),
		ColorG: float32(c2.G) / float32(math.MaxUint8),
		ColorB: float32(c2.B) / float32(math.MaxUint8),
		ColorA: float32(c2.A) / float32(math.MaxUint8),
	}
	op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe
	var indices []uint16
	indices = append(indices, 0, 1, 2, 2, 3, 0)
	dst.DrawTriangles(vertices, indices, app.WhiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)
}
