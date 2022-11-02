package scene

//import (
//	"github.com/hajimehoshi/ebiten/v2"
//	helper_graphic "github.com/langwan/langgo/helpers/graphic"
//	"image"
//	"image/color"
//	"math"
//	"spectrum/app"
//	"spectrum/octavebands"
//)
//
//const (
//	smoothingTime    float64 = 0.7
//	smoothingTimeCap float64 = 0.9
//)
//
//var ledCache []float64
//
//type Led struct {
//	Fraction       float64
//	SpacingFactorX float64
//	MinDecibels    float64
//	MaxDecibels    float64
//	LedCellHeight  float64
//	SpacingFactorY float64
//	Color1         color.RGBA
//	Color2         color.RGBA
//}
//
//var (
//	whiteImage = ebiten.NewImage(3, 3)
//)
//
//func init() {
//	whiteImage.Fill(color.White)
//
//}
//
//func (l Led) Update() error {
//	return nil
//}
//
//type LinearGradientColor struct {
//	C1 color.RGBA
//	C2 color.RGBA
//}
//
//var LinearGradientColors []LinearGradientColor
//
////var samples []float64
////func init() {
////	rand.Seed(time.Now().UnixNano())
////	samples = make([]float64, 1024)
////	for i := range samples {
////		samples[i] = float64(rand.Intn(255))
////	}
////}
//
//func (l Led) Draw(screen *ebiten.Image) {
//	sw, sh := screen.Size()
//	banks := octavebands.GenBanks(l.Fraction)
//	frequency := l.GenFrequency(banks, app.FftSamples[0])
//	if len(ledCache) == 0 {
//		ledCache = make([]float64, len(frequency))
//	}
//	xLeds := float64(len(banks))
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
//	//for i := range LinearGradientColors {
//	//	LinearGradientColors[i] = LinearGradientColor{
//	//		C1: color.RGBA{},
//	//		C2: color.RGBA{},
//	//	}
//	//}
//
//	//for i := 0.0; i < xLeds; i++ {
//	//	for j := 0.0; j < yLeds; j++ {
//	//		ebitenutil.DrawRect(screen, i*(ledCellWidth)+xPadding, float64(sh)-(j+1.0)*(l.LedCellHeight)+yPadding, ledWidth, ledHeight, color.RGBA{
//	//			R: 0x11,
//	//			G: 0x11,
//	//			B: 0x11,
//	//			A: 0xff,
//	//		})
//	//	}
//	//}
//	dr := (float64(l.Color2.R) - float64(l.Color1.R)) / yLeds
//	dg := (float64(l.Color2.G) - float64(l.Color1.B)) / yLeds
//	db := (float64(l.Color2.B) - float64(l.Color1.G)) / yLeds
//	da := (float64(l.Color2.A) - float64(l.Color1.A)) / yLeds
//	for i := 0; i < int(xLeds); i++ {
//		val := frequency[int(i)]
//		ledCache[i] = ledCache[i]*smoothingTime + (1-smoothingTime)*val
//		yLedNumber := math.Round(ledCache[i] / 255.0 * yLeds)
//		//	yLedNumber = yLeds
//
//		for j := 0.0; j < yLedNumber; j++ {
//			DrawBar(screen, float64(i)*(ledCellWidth)+xPadding, float64(sh)-(j+1.0)*(l.LedCellHeight)+yPadding, ledWidth, ledHeight, color.RGBA{
//				R: uint8(float64(l.Color2.R) - (j+1)*dr),
//				G: uint8(float64(l.Color2.G) - (j+1)*dg),
//				B: uint8(float64(l.Color2.B) - (j+1)*db),
//				A: uint8(float64(l.Color2.A) - (j+1)*da),
//			}, color.RGBA{
//				R: uint8(float64(l.Color2.R) - (j)*dr),
//				G: uint8(float64(l.Color2.G) - (j)*dg),
//				B: uint8(float64(l.Color2.B) - (j)*db),
//				A: uint8(float64(l.Color2.A) - (j)*da),
//			})
//		}
//	}
//}
//
//func (l Led) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
//	return 0, 0
//}
//
//func DrawBar(dst *ebiten.Image, x, y, width, height float64, c1 color.RGBA, c2 color.RGBA) {
//	var vertices = make([]ebiten.Vertex, 4)
//	vertices[0] = ebiten.Vertex{
//		DstX:   float32(x),
//		DstY:   float32(y),
//		SrcX:   0,
//		SrcY:   0,
//		ColorR: float32(c1.R) / float32(math.MaxUint8),
//		ColorG: float32(c1.G) / float32(math.MaxUint8),
//		ColorB: float32(c1.B) / float32(math.MaxUint8),
//		ColorA: float32(c1.A) / float32(math.MaxUint8),
//	}
//
//	vertices[1] = ebiten.Vertex{
//		DstX:   float32(x + width),
//		DstY:   float32(y),
//		SrcX:   0,
//		SrcY:   0,
//		ColorR: float32(c1.R) / float32(math.MaxUint8),
//		ColorG: float32(c1.G) / float32(math.MaxUint8),
//		ColorB: float32(c1.B) / float32(math.MaxUint8),
//		ColorA: float32(c1.A) / float32(math.MaxUint8),
//	}
//
//	vertices[2] = ebiten.Vertex{
//		DstX:   float32(x + width),
//		DstY:   float32(y + height),
//		SrcX:   0,
//		SrcY:   0,
//		ColorR: float32(c2.R) / float32(math.MaxUint8),
//		ColorG: float32(c2.G) / float32(math.MaxUint8),
//		ColorB: float32(c2.B) / float32(math.MaxUint8),
//		ColorA: float32(c2.A) / float32(math.MaxUint8),
//	}
//
//	vertices[3] = ebiten.Vertex{
//		DstX:   float32(x),
//		DstY:   float32(y + height),
//		SrcX:   0,
//		SrcY:   0,
//		ColorR: float32(c2.R) / float32(math.MaxUint8),
//		ColorG: float32(c2.G) / float32(math.MaxUint8),
//		ColorB: float32(c2.B) / float32(math.MaxUint8),
//		ColorA: float32(c2.A) / float32(math.MaxUint8),
//	}
//	op := &ebiten.DrawTrianglesOptions{}
//	op.Address = ebiten.AddressUnsafe
//	var indices []uint16
//	indices = append(indices, 0, 1, 2, 2, 3, 0)
//	dst.DrawTriangles(vertices, indices, whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)
//}
//
//func (l Led) GenFrequency(banks []octavebands.Band, samples []float64) []float64 {
//	var frequency = make([]float64, len(banks))
//	if len(samples) == 0 {
//		return frequency
//	}
//
//	sum := 0.0
//	n := 0.0
//	for i, bank := range banks {
//		for j := bank.Min; j <= bank.Max; j++ {
//			sj, _ := helper_graphic.RangeMapper(j, 0.0, float64(app.SampleRate)/2.0, 0.0, float64(len(samples)-1))
//			sj = math.Round(sj)
//			sum += samples[int(sj)] * samples[int(sj)]
//			n++
//		}
//		if sum == 0 {
//			frequency[i] = 0
//		} else {
//			frequency[i] = math.Sqrt(sum / n)
//		}
//
//	}
//	return frequency
//}
