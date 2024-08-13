package main

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func main() {
	f, err := os.Open("../audio/NokiaxBetterCallSaul.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)

	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)

	select {}
}
