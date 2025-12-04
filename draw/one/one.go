package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// 创建图表
	p := plot.New()
	p.Title.Text = "Catalyst Performance Comparison"
	p.X.Label.Text = "Catalyst System"
	p.Y.Label.Text = "Yield (%)"

	// 准备数据
	catalysts := []string{"AlCl₃·6H₂O", "ZrO₂-SO₄²⁻ (fresh)", "ZrO₂-SO₄²⁻ (5 cycles)"}
	dbmYields := plotter.Values{90.2, 88.7, 70.7}
	mbmYields := plotter.Values{0.7, 0, 0}

	// 设置柱状图宽度
	w := vg.Points(20)

	// 创建DBM收率柱状图
	dbmBars, err := plotter.NewBarChart(dbmYields, w)
	if err != nil {
		log.Fatal(err)
	}
	dbmBars.LineStyle.Width = vg.Length(0)
	dbmBars.Color = plotutil.Color(0)
	dbmBars.Offset = -w / 2

	// 创建MBM收率柱状图
	mbmBars, err := plotter.NewBarChart(mbmYields, w)
	if err != nil {
		log.Fatal(err)
	}
	mbmBars.LineStyle.Width = vg.Length(0)
	mbmBars.Color = plotutil.Color(1)
	mbmBars.Offset = w / 2

	// 添加到图表
	p.Add(dbmBars, mbmBars)
	p.Legend.Add("DBM Yield", dbmBars)
	p.Legend.Add("MBM Yield", mbmBars)
	p.NominalX(catalysts...)

	// 保存图表
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "catalyst_comparison.png"); err != nil {
		log.Fatal(err)
	}
}
