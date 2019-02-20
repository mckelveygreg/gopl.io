// HTTP Server that echoes the HTTP Header
package main

import (
	"strconv"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		cs, ok := q["cycles"]
		if ok {
			s = css[0]
			strconv.Atoi
		}
		lissajous(w, 5)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// ex1.5
var green = color.RGBA{0, 255, 0, 255} // is there a better way to do this?
var blue = color.RGBA{0, 0, 255, 255}
var red = color.RGBA{255, 0, 0, 0}
var mystery = color.RGBA{200, 50, 0, 255}

var palette = []color.Color{color.Black, green, color.White, blue, red, mystery}

const (
	background = 0 // first color in palette
	greenIndex = 1 // next color in palette
)

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.0001 // angular resolution
		size    = 300    // image canvas covers [-size..+size]
		nframes = 64     // number of animation frames
		delay   = 8      // delay between frames in 10ms units
	)
	if cycles <= 0 {
		cycles = 5
	}
	// number of complete x oscillator revolutions
	freq := rand.Float64() * 2.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(rand.Intn(5)))
		}
		phase += 0.05
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
