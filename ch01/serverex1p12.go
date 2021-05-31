// Server modified to display animated gif based on settings provided to the web server. From Daniel Harms at https://github.com/jdharms/tgpl
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func getParam(rqust *http.Request, name string) int {
	param := rqust.URL.Query().Get(name)
	if len(param) == 0 {
		return 0
	}

	value, err := strconv.Atoi(param)
	if err != nil {
		fmt.Fprintf(os.Stderr, "server: %v\n", err)
		return 0
	}

	return value
}

func main() {
	http.HandleFunc("/", func(wrtr http.ResponseWriter, rqust *http.Request) {
		cycles := getParam(rqust, "cycles")
		size := getParam(rqust, "size")
		nframes := getParam(rqust, "nframes")
		delay := getParam(rqust, "delay")

		lissajous(wrtr, cycles, size, nframes, delay)
	})

	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

var palette = []color.Color{color.Black, color.RGBA{0x00, 0x60, 0x00, 0xFF}}

const (
	blackIndex = 0
	greenIndex = 1
)

func lissajous(out io.Writer, cycles int, size int, nframes int, delay int) {
	// Luckily, zero doesn't make sense as a value for any of these parameters,
	// so we'll use it as nil
	const (
		defaultCycles  = 5     // number of complete x oscillator revolutions
		res            = 0.001 // angular resolution
		defaultSize    = 100   // image canvas covers [-size..+size]
		defaultNframes = 64    // number of animation frames
		defaultDelay   = 8     // delay between frames in 10ms units
	)

	if cycles == 0 {
		cycles = defaultCycles
	}

	if size == 0 {
		size = defaultSize
	}

	if nframes == 0 {
		nframes = defaultNframes
	}

	if delay == 0 {
		delay = defaultDelay
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
