package sdk

import (
	"errors"
	"strings"
)

type PlayList struct {
	Files        []string
	Player       *Player
	CurrentIndex int
}

func (p *PlayList) PlayList(files []string) error {
	//speaker.Lock()
	//defer speaker.Unlock()
	//if p.Current.Streamer != nil {
	//	p.Current.Streamer.Close()
	//}
	//if p.Current.FileHandle != nil {
	//	p.Current.FileHandle.Close()
	//}
	p.Files = nil
	for _, f := range files {
		if !strings.HasSuffix(f, ".mp3") {
			continue
		}
		p.Files = append(p.Files, f)
	}

	if len(p.Files) < 1 {
		return errors.New("playlist is empty")
	}
	return nil
}

func (p *PlayList) Next() (err error) {
	if len(p.Files) == 0 {
		return errors.New("playlist is empty")
	}
	if len(p.Files) == 1 {
		p.CurrentIndex = 0
	} else if p.CurrentIndex > len(p.Files)-1 {
		p.CurrentIndex = len(p.Files) - 1
	} else {
		p.CurrentIndex = p.CurrentIndex + 1
	}
	return p.Play(p.CurrentIndex)
}

func (p *PlayList) Prev() (err error) {
	if len(p.Files) == 0 {
		return errors.New("playlist is empty")
	}
	if p.CurrentIndex < 1 {
		p.CurrentIndex = 0
	} else {
		p.CurrentIndex = p.CurrentIndex - 1
	}
	return p.Play(p.CurrentIndex)
}

func (p *PlayList) Play(index int) (err error) {
	p.CurrentIndex = index
	p.Player.Play(p.Files[index])
	return nil
}

func (p *PlayList) Playing(isPlay bool) {
	p.Player.IsPlay = isPlay
}

func (p *PlayList) hasNext() (int, bool) {
	if p.CurrentIndex < len(p.Files)-1 {
		return p.CurrentIndex + 1, true
	} else {
		return -1, false
	}
}
