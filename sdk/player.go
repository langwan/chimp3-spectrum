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
	Mode                 int
	IsPlay               bool
	Change               chan string
	Buffer               []byte
	UpdateFrequencyFrame func(p *Player, frame FrequencyFrame)
	PlayList             *PlayList
}

func New() *Player {
	p := Player{}
	p.otoContext, _ = oto.NewContext(44100, 2, 2, FftSize*channels*bitDepth)
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
			//p.Buffer = make([]byte, 17640)
			p.Current.Mp3Decoder.Seek(0, 0)
			p.IsPlay = true
		default:
			if p.IsPlay {
				log.Logger("backend", "player.Update").Debug().Bool("isplay", p.IsPlay).Str("name", p.Current.Name).Send()
				buf := make([]byte, FftSize*channels*bitDepth)

				isEof := false

				reads, err := p.Current.Mp3Decoder.Read(buf)
				//for i := 0; i < read; i++ {
				//	buf[reads+i] = p.Buffer[i]
				//}

				if err == io.EOF {
					isEof = true
					
				}

				spectrum, err := frequencySpectrum(buf[:reads], p.Current.Mp3Decoder.SampleRate(), FftSize)
				if err != nil {

					return
				}
				sampleRate := 44100
				if p.Current.Mp3Decoder != nil {
					sampleRate = p.Current.Mp3Decoder.SampleRate()
				}
				p.sendFrequencyFrame(FrequencyFrame{
					SampleRate: sampleRate,
					Name:       p.Current.Name,
					Filepath:   p.Current.Filepath,
					IsPlay:     p.IsPlay,
					Spectrum:   spectrum,
				})
				_, err = p.otoPlayer.Write(buf[:reads])
				if err != nil {

					return
				}

				if isEof {
					p.IsPlay = false
					sampleRate := 44100
					if p.Current.Mp3Decoder != nil {
						sampleRate = p.Current.Mp3Decoder.SampleRate()
					}
					p.sendFrequencyFrame(FrequencyFrame{
						SampleRate: sampleRate,
						Name:       p.Current.Name,
						Filepath:   p.Current.Filepath,
						IsPlay:     p.IsPlay,
						Spectrum:   [2][]float64{},
					})
					nextId, ok := p.PlayList.hasNext()
					if ok {
						go func() {

							p.PlayList.Play(nextId)
						}()
					}
				}
			} else {
				sampleRate := 44100
				if p.Current.Mp3Decoder != nil {
					sampleRate = p.Current.Mp3Decoder.SampleRate()
				}
				p.sendFrequencyFrame(FrequencyFrame{
					SampleRate: sampleRate,
					Name:       p.Current.Name,
					Filepath:   p.Current.Filepath,
					IsPlay:     p.IsPlay,
					Spectrum:   [2][]float64{},
				})
				time.Sleep(time.Second / 10)
			}
		}
	}
}

func (p *Player) sendFrequencyFrame(frame FrequencyFrame) {
	if p.UpdateFrequencyFrame != nil {
		p.UpdateFrequencyFrame(p, frame)
	}
}
