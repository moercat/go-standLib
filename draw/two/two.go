package main

import (
	"log"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// 创建图表
	p := plot.New()
	p.Title.Text = "Catalyst Deactivation and Regeneration"
	p.X.Label.Text = "Time (h)"
	p.Y.Label.Text = "Relative Activity (%)"

	// 准备数据
	points := plotter.XYs{
		{X: 0, Y: 100},
		{X: 200, Y: 98},
		{X: 400, Y: 96},
		{X: 600, Y: 94},
		{X: 800, Y: 93},
		{X: 1000, Y: 95},
	}

	// 创建折线
	line, err := plotter.NewLine(points)
	if err != nil {
		log.Fatal(err)
	}
	line.LineStyle.Width = vg.Points(2)
	line.LineStyle.Color = plotutil.Color(0)

	// 创建散点
	scatter, err := plotter.NewScatter(points)
	if err != nil {
		log.Fatal(err)
	}
	scatter.GlyphStyle.Color = plotutil.Color(0)
	scatter.GlyphStyle.Radius = vg.Points(3)

	// 添加到图表
	p.Add(line, scatter)

	// 添加数据点标签
	for _, pt := range points {
		// 在数据点上方添加数值标签
		l, err := plotter.NewLabels(plotter.XYLabels{
			XYs:    []plotter.XY{{X: pt.X, Y: pt.Y + 1}}, // 稍微向上偏移
			Labels: []string{strconv.FormatFloat(pt.Y, 'f', 1, 64) + "%"},
		})
		if err != nil {
			log.Fatal(err)
		}
		p.Add(l)
	}

	// 添加再生点特殊标记
	regenerationPoint := plotter.XY{X: 1000, Y: 95}
	regenerationScatter, err := plotter.NewScatter(plotter.XYs{regenerationPoint})
	if err != nil {
		log.Fatal(err)
	}
	regenerationScatter.GlyphStyle.Color = plotutil.Color(2) // 红色
	regenerationScatter.GlyphStyle.Radius = vg.Points(5)     // 更大的点
	regenerationScatter.GlyphStyle.Shape = plotutil.Shape(4) // 菱形标记
	p.Add(regenerationScatter)

	// 添加再生标注
	regenerationLabel, err := plotter.NewLabels(plotter.XYLabels{
		XYs:    []plotter.XY{{X: 1000, Y: 97}}, // 标注位置
		Labels: []string{"Regeneration\n(5% HCl wash\n200°C calcination)"},
	})
	if err != nil {
		log.Fatal(err)
	}
	p.Add(regenerationLabel)

	// 设置Y轴范围，确保所有点都可见
	p.Y.Min = 90
	p.Y.Max = 102

	// 保存图表
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "deactivation_regeneration.png"); err != nil {
		log.Fatal(err)
	}
}
