package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"sync"

	"math"
	"math/rand"

	"image"
	"image/color"
	"image/gif"
)

// a minimal "echo" and counter server
var mu sync.Mutex // 锁
var count int     // 访问量计数
func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	lissajous(w)
	count++
	mu.Unlock()
	// 原始版
	// fmt.Fprintf(w, "URL.PATH=%q\n", r.URL.Path)
	// 加强版
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// 遍历网页header
	for k, v := range r.Header {
		// 打印网页header
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	// 打印网页host
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	// 打印网页远程地址
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	// 如果处理出错，打印错误信息
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	// 打印网页信息格式
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// 添加网页动图
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers[-size,..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
