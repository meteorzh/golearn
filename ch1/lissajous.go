// lissajous产生随机利萨如图形的GIF动画
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// var palette = []color.Color{color.White, color.Black}
// var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}} // 练习：改为绿色黑底
var palette = []color.Color{color.Black, color.RGBA{0xFF, 0x00, 0x00, 0xFF}, color.RGBA{0x00, 0xFF, 0x00, 0xFF}, color.RGBA{0x00, 0x00, 0xFF, 0xFF}} // 练习：添加更多颜色

const (
	whiteIndex = 0 // 画板中的第一种颜色
	blackIndex = 1 // 画板中的下一种颜色
)

func mainlis() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整的x振荡器变化的个数
		res     = 0.001 // 角度分辨率
		size    = 100   // 图像画布包含[-size..+size]
		nframes = 64    // 动画中的帧数
		delay   = 8     // 以10ms为单位的帧间延迟
	)

	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	randIdx := rand.Intn(3) + 1  // 练习：添加更多颜色(每次调用随机使用一个颜色)
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(randIdx))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意：忽略编码错误
}
