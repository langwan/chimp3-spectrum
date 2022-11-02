package sdk

import (
	"bytes"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"github.com/langwan/langgo/core/log"
	"io"
	"path/filepath"
	"time"

	"os"
	"strings"
)

type Streamer struct {
}

type Player struct {
	otoContext *oto.Context
	otoPlayer  *oto.Player
	Current    struct {
		Name       string
		Filepath   string
		Mp3Decoder *mp3.Decoder
		FileBuffer []byte
	}
	Mode          int
	IsPlay        bool
	Change        chan string
	Buffer        []byte
	UpdateSamples func(p *Player, frame Frame)
	PlayList      *PlayList
}

func New() *Player {

	p := Player{}
	p.otoContext, _ = oto.NewContext(44100, 2, 2, 17640)
	p.otoPlayer = p.otoContext.NewPlayer()

	p.Change = make(chan string)
	go p.Update()
	return &p
}

func (p *Player) Play(filename string) (err error) {
	p.Change <- filename
	return nil
}

func (p *Player) Update() {
	for {
		select {
		case filename := <-p.Change:
			name := strings.Split(filepath.Base(filename), ".")
			p.Current.Name = name[0]
			p.Current.Filepath = filename
			p.IsPlay = false
			var err error
			p.Current.FileBuffer, err = os.ReadFile(filename)
			if err != nil {
				continue
			}
			reader := bytes.NewReader(p.Current.FileBuffer)
			p.Current.Mp3Decoder, err = mp3.NewDecoder(reader)
			if err != nil {
				continue
			}
			p.Buffer = make([]byte, 17640)
			p.Current.Mp3Decoder.Seek(0, 0)
			p.IsPlay = true
		default:
			if p.IsPlay {
				log.Logger("backend", "player.Update").Debug().Bool("isplay", p.IsPlay).Str("name", p.Current.Name).Send()
				buf := make([]byte, 17640)
				reads := 0
				isEof := false

				for {
					read, err := p.Current.Mp3Decoder.Read(p.Buffer[:len(buf)-reads])
					for i := 0; i < read; i++ {
						buf[reads+i] = p.Buffer[i]
					}
					reads += read
					if err == io.EOF {
						isEof = true
						break
					} else if reads == len(buf) {
						break
					}
				}

				samples, err := readSamples(buf[:reads])
				if err != nil {

					return
				}
				p.updateFrames(samples)
				_, err = p.otoPlayer.Write(buf[:reads])
				if err != nil {

					return
				}

				if isEof {
					p.IsPlay = false

					p.UpdateSamples(p, Frame{
						Name:     p.Current.Name,
						Filepath: p.Current.Filepath,
						IsPlay:   p.IsPlay,
						Samples:  nil,
						Mode:     p.Mode,
					})
					nextId, ok := p.PlayList.hasNext()
					if ok {
						go func() {

							p.PlayList.Play(nextId)
						}()
					}
				}
			} else {
				if p.UpdateSamples != nil {
					p.UpdateSamples(p, Frame{
						Name:     p.Current.Name,
						Filepath: p.Current.Filepath,
						IsPlay:   p.IsPlay,
						Samples:  nil,
						Mode:     p.Mode,
					})
				}

				time.Sleep(time.Second / 10)
			}
		}
	}
}

func (p *Player) updateFrames(samples *[][2]float64) {
	if samples == nil {
		return
	}

	var ware [2][]float64
	ware[0] = make([]float64, len(*samples))
	ware[1] = make([]float64, len(*samples))
	for i := 0; i < len(*samples); i++ {
		ware[0][i] = (*samples)[i][0]
		ware[1][i] = (*samples)[i][1]
	}

	switch p.Mode {
	case 0:
		Mode0(&ware)
	default:
		Mode0(&ware)
	}

	//wareReal := fft.FFTReal(ware[0])
	//var max float64 = 0
	//for i := 0; i < len(*samples); i++ {
	//	fr := real(wareReal[i])
	//	fi := imag(wareReal[i])
	//	magnitude := math.Sqrt(fr*fr + fi*fi)
	//	ware[0][i] = magnitude
	//	if magnitude > max {
	//		max = magnitude
	//	}
	//}
	//for i := 0; i < len(*samples); i++ {
	//	ware[0][i], _ = helper_graphic.RangeMapper(ware[0][i], 0, max, 0, 60)
	//}
	//if p.Mode == 0 || p.Mode == 3 {
	//	wareReal = fft.FFTReal(ware[1])
	//} else {
	//	wareReal = fft.FFTReal(ware[0])
	//}
	//
	//max = 0
	//for i := 0; i < len(*samples); i++ {
	//	fr := real(wareReal[i])
	//	fi := imag(wareReal[i])
	//	magnitude := math.Sqrt(fr*fr + fi*fi)
	//	ware[1][i] = magnitude
	//	if magnitude > max {
	//		max = magnitude
	//	}
	//}
	//for i := 0; i < len(*samples); i++ {
	//	ware[1][i], _ = helper_graphic.RangeMapper(ware[1][i], 0, max, 0, 60)
	//}
	if p.UpdateSamples != nil {
		p.UpdateSamples(p, Frame{
			Name:     p.Current.Name,
			Filepath: p.Current.Filepath,
			IsPlay:   p.IsPlay,
			Samples:  &ware,
			Mode:     p.Mode,
		})
	}

}

func readSamples(buf []byte) (*[][2]float64, error) {
	format := Format{
		SampleRate:  44100,
		NumChannels: 2,
		Precision:   2,
	}
	samples := make([][2]float64, len(buf)/(format.NumChannels*format.Precision))
	var tmp [4]byte
	reader := bytes.NewReader(buf)
	i := 0
	for {
		_, err := reader.Read(tmp[:])
		if err == io.EOF {
			break
		} else if err != nil {
			return &samples, err
		} else {
			samples[i], _ = format.DecodeSigned(tmp[:])
			i++
		}
	}
	return &samples, nil
}
