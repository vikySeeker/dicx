package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/oto/v2"
	"github.com/hajimehoshi/go-mp3"
)

/* func getAudio() (io.Reader, error) {
	audio, err := http.Get("http://ssl.gstatic.com/dictionary/static/sounds/20200429/hello--_gb_1.mp3")
	if err != nil {
		return nil, err
	}
	return audio.Body, nil

} */

func run() error {
	f, err := os.Open("NokiaxBetterCallSaul.mp3")
	//f, err := getAudio()
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, _, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	//defer c.Close()

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()

	fmt.Printf("Length: %d[bytes]\n", d.Length())

	/* if _, err := io.Copy(p, d); err != nil {
		return err
	} */
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}