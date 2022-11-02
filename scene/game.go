package scene

//import (
//	"github.com/hajimehoshi/ebiten/v2"
//	helper_graphic "github.com/langwan/langgo/helpers/graphic"
//	"image"
//	"image/color"
//	"math"
//	"spectrum/app"
//	"sync"
//)
//
//const (
//	windowWidth      float64 = 1080
//	barPadding       float64 = 2
//	capHeight        float64 = 2
//	smoothingTime    float64 = 0.7
//	smoothingTimeCap float64 = 0.95
//)
//
//var (
//	windowHeight = math.Round(windowWidth / 16.0 * 9)
//	game         *Game
//)
//
//type Band struct {
//	lower float64
//	upper float64
//}
//
//var gSamplesCache [2][]float64
//var gCapCache [2][]float64
//
//type GameState int
//
//const (
//	gameStateStart   = 0
//	gameStatePlaying = 1
//)
//
//type Game struct {
//	state          ebiten.Game
//	frequencyBands float64
//	startFrequency float64
//	endFrequency   float64
//	samples        chan [2][]float64
//	bands          []Band
//	mu             sync.RWMutex
//}
//
//func (g Game) Update() error {
//
//	return nil
//}
//
//var samplesCache *[2][]float64
//var samplesBuffer []float64
//
//func (g Game) Draw(screen *ebiten.Image) {
//	led.DrawA(screen)
//	led.DrawB(screen)
//
//	return
//	game.mu.Lock()
//	defer game.mu.Unlock()
//	//msg := fmt.Sprintf("TPS: %0.2f FPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS())
//	//ebitenutil.DebugPrint(screen, msg)
//	//samples := <-g.samples
//	DrawBars(screen, app.FftSamples, 0)
//	DrawBars(screen, app.FftSamples, 1)
//}
//
//func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
//	s := ebiten.DeviceScaleFactor()
//	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
//}
//
//func NewWin() error {
//
//	ebiten.SetWindowSize(int(windowWidth), int(windowHeight))
//	ebiten.SetWindowTitle("chihuo-mp3-spectrum")
//
//	if err := ebiten.RunGame(game); err != nil {
//		return err
//	}
//	return nil
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
//func DrawCap(dst *ebiten.Image, x, y, width, height float64) {
//	var vertices = make([]ebiten.Vertex, 4)
//	vertices[0] = ebiten.Vertex{
//		DstX:   float32(x),
//		DstY:   float32(y),
//		SrcX:   0,
//		SrcY:   0,
//		ColorR: float32(1),
//		ColorG: float32(1),
//		ColorB: float32(1),
//		ColorA: float32(1),
//	}
//
//	vertices[1] = ebiten.Vertex{
//		DstX:   float32(x + width),
//		DstY:   float32(y),
//		SrcX:   0,
//		SrcY:   0,
//		ColorR: float32(1),
//		ColorG: float32(1),
//		ColorB: float32(1),
//		ColorA: float32(1),
//	}
//
//	vertices[2] = ebiten.Vertex{
//		DstX:   float32(x + width),
//		DstY:   float32(y + height),
//		SrcX:   0,
//		SrcY:   0,
//		ColorR: float32(1),
//		ColorG: float32(1),
//		ColorB: float32(1),
//		ColorA: float32(1),
//	}
//
//	vertices[3] = ebiten.Vertex{
//		DstX:   float32(x),
//		DstY:   float32(y + height),
//		SrcX:   0,
//		SrcY:   0,
//		ColorR: float32(1),
//		ColorG: float32(1),
//		ColorB: float32(1),
//		ColorA: float32(1),
//	}
//	op := &ebiten.DrawTrianglesOptions{}
//	op.Address = ebiten.AddressUnsafe
//	var indices []uint16
//	indices = append(indices, 0, 1, 2, 2, 3, 0)
//	dst.DrawTriangles(vertices, indices, whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)
//}
//
//func DrawBars(screen *ebiten.Image, samples [2][]float64, channel int) {
//	if len(samples[channel]) == 0 {
//		return
//	}
//
//	sw, sh := screen.Size()
//	var spectrumHeight float64 = float64(sh) - 2
//	barWidth := (float64(sw)-barPadding)/game.frequencyBands - barPadding
//
//	barNum := int(math.Floor(float64(sw)/barWidth + barPadding))
//	var bars []float64
//	step := int(math.Floor(float64(len(samples[channel])) / float64(barNum)))
//
//	for i := 0; i < barNum; i++ {
//		bars = append(bars, samples[channel][i*step+channel])
//	}
//
//	if len(gCapCache[channel]) == 0 {
//		gCapCache[channel] = make([]float64, len(bars))
//		for i, value := range bars {
//			gCapCache[channel][i] = value
//		}
//	}
//
//	if len(gSamplesCache[channel]) == 0 {
//		gSamplesCache[channel] = make([]float64, len(bars))
//		for i, value := range bars {
//
//			gSamplesCache[channel][i] = value
//		}
//	}
//	start := barPadding
//	//if channel == 1 {
//	//	start = barPadding + float64(len(bars))/2*(barWidth+barPadding)
//	//}
//	index := 0
//
//	for i, _ := range bars {
//		if channel == 0 {
//			index = i
//		} else {
//			index = len(bars) - 1 - i
//		}
//		v, _ := helper_graphic.RangeMapper(bars[index], 0, math.MaxUint8, 0, spectrumHeight)
//		gSamplesCache[channel][index] = gSamplesCache[channel][index]*smoothingTime + (1-smoothingTime)*v
//		vv := gSamplesCache[channel][index]
//		if vv < gCapCache[channel][index] {
//			n := vv - gCapCache[channel][index]
//			gCapCache[channel][index] = gCapCache[channel][index]*(smoothingTimeCap) + (1-smoothingTimeCap)*(math.Max(n, -1))
//			DrawCap(screen, start+float64(i)*(barWidth+barPadding), spectrumHeight-gCapCache[channel][index], barWidth, capHeight)
//		} else {
//			gCapCache[channel][index] = vv
//			DrawCap(screen, start+float64(i)*(barWidth+barPadding), spectrumHeight-gCapCache[channel][index], barWidth, capHeight)
//		}
//		if channel == 0 {
//			DrawBar(screen, start+float64(i)*(barWidth+barPadding), spectrumHeight-vv, barWidth, vv, color.RGBA{
//				R: 99,
//				G: 151,
//				B: 189,
//				A: 150,
//			}, color.RGBA{
//				R: 51,
//				G: 59,
//				B: 64,
//				A: 150,
//			})
//		} else {
//			DrawBar(screen, start+float64(i)*(barWidth+barPadding), spectrumHeight-vv, barWidth, vv, color.RGBA{
//				R: 209,
//				G: 120,
//				B: 39,
//				A: 150,
//			}, color.RGBA{
//				R: 118,
//				G: 77,
//				B: 43,
//				A: 150,
//			})
//		}
//
//	}
//
//}
