package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // 画布大小，以像素为单位
	cells         = 100                 // 网格单元数
	xyrange       = 30.0                // 坐标轴范围
	xyscale       = width / 2 / xyrange // 每 x 或 y 单位像素
	zscale        = height * 0.4        // 每 z 单位像素
	angle         = math.Pi / 6         // x, y 轴的角度（=30°）
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// 003、通过浮点计算生成图形
// go run 003_surface.go
// 输出：
// svg 图案
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	draw(w)
}

// 绘制 svg
func draw(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// 在单元 i, j 的角落里找到点 x, y
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 计算表面高度 z
	z := f(x, y)

	// 将 x, y, z 投影到二维 svg 画布 sx, sy 上
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // 与 0, 0 的距离
	return math.Sin(r) / r
}
