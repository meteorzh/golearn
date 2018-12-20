package main

import (
	"fmt"
	"golearn/ch7/eval"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // 以像素表示的画布大小
	cells         = 100                 // 网格单元的个数
	xyrange       = 30.0                // 坐标轴的范围 （-xyrange..+xyrange）
	xyscale       = width / 2 / xyrange // x或y轴上每个单位长度的像素
	zscale        = height * 0.4        // z轴上每个单位长度的像素
	angle         = math.Pi / 6         // x, y轴的角度（=30°）
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y) // 与 (0, 0) 之间的距离
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}

func surface(w io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int, f func(x, y float64) float64) (float64, float64) {
	// 求出网格单元（i, j）的顶点坐标(x, y)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 计算曲面高度z
	z := f(x, y)

	// 将(x, y, z)等角投射到二维 SVG 绘图平面上, 坐标是 (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/plot", plot)
	// http3a
	// mux.HandleFunc("/list", db.list)
	// mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
