package main

import "time"

func main() {
	game = &Game{
		frequencyBands: 256,
		samples:        make(chan [2][]float64),
	}

	go func() {
		time.Sleep(time.Second * 3)
		Player("./samples/1.mp3", func(spectrum [2][]float64) {

			//game.samples <- spectrum

			gSamples[0] = spectrum[0]
			gSamples[1] = spectrum[1]

			//et := time.Now()
			//fmt.Printf("t = %s\n", et.Sub(st))
			//st = et
		})
	}()
	err := NewWin()
	if err != nil {
		panic(err)
	}
}
