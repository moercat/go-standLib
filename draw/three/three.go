package main

import (
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// 创建图表
	p := plot.New()
	p.Title.Text = "DFT Calculation of Energy Barrier"
	p.X.Label.Text = "Reaction Coordinate"
	p.Y.Label.Text = "Energy (kJ/mol)"

	// 生成模拟数据
	withoutCatalyst := make(plotter.XYs, 100)
	withCatalyst := make(plotter.XYs, 100)

	for i := 0; i < 100; i++ {
		x := float64(i) / 10
		withoutCatalyst[i].X = x
		withoutCatalyst[i].Y = 50 + 20*math.Sin(x) + 5*math.Cos(3*x)

		withCatalyst[i].X = x
		withCatalyst[i].Y = 40 + 15*math.Sin(x) + 5*math.Cos(3*x)
	}

	// 创建线条和散点
	line1, points1, err := plotter.NewLinePoints(withoutCatalyst)
	if err != nil {
		log.Fatal(err)
	}
	line1.LineStyle.Width = vg.Points(2)
	line1.LineStyle.Color = plotutil.Color(0)
	points1.Shape = plotutil.DefaultGlyphShapes[0]
	points1.Color = plotutil.Color(0)

	line2, points2, err := plotter.NewLinePoints(withCatalyst)
	if err != nil {
		log.Fatal(err)
	}
	line2.LineStyle.Width = vg.Points(2)
	line2.LineStyle.Color = plotutil.Color(1)
	points2.Shape = plotutil.DefaultGlyphShapes[1]
	points2.Color = plotutil.Color(1)

	// 添加到图表
	p.Add(line1, points1, line2, points2)
	p.Legend.Add("Without Catalyst", line1)
	p.Legend.Add("With AlCl₃ Catalyst", line2)

	// 保存图表
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "dft_energy_barrier.png"); err != nil {
		log.Fatal(err)
	}
}
