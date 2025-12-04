package main

import (
	"log"
	"math"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	// 图1：催化剂性能对比图
	createCatalystComparisonPlot()

	// 图2：DFT计算闭环能垒图
	createDFTEnergyBarrierPlot()

	// 图3：催化剂失活与再生图
	createDeactivationRegenerationPlot()

	// 图4：Pareto图（因素贡献度）
	createParetoChart()
}

// 图1：催化剂性能对比图
func createCatalystComparisonPlot() {
	// 创建图表
	p := plot.New()
	p.Title.Text = "Catalyst Performance Comparison"
	p.Title.TextStyle.Font.Size = 20
	p.Title.Padding = 10 * vg.Millimeter

	// 设置X轴和Y轴
	p.X.Label.Text = "Catalyst System"
	p.X.Label.TextStyle.Font.Size = 16
	p.X.Tick.Label.Font.Size = 14
	p.Y.Label.Text = "Yield (%)"
	p.Y.Label.TextStyle.Font.Size = 16
	p.Y.Tick.Label.Font.Size = 14

	// 准备数据
	catalysts := []string{"AlCl₃·6H₂O\nDBM", "AlCl₃·6H₂O\nMBM", "ZrO₂-SO₄²⁻\nFresh", "ZrO₂-SO₄²⁻\n5 cycles"}
	dbmYields := []float64{90.2, 0, 88.7, 0}
	mbmYields := []float64{0.7, 0, 0, 0}
	after5Cycles := []float64{0, 0, 0, 70.7} // 循环5次后活性下降18%

	// 创建柱状图
	w := vg.Points(15)

	// DBM收率柱状图
	dbmBars, err := plotter.NewBarChart(plotter.Values(dbmYields), w)
	if err != nil {
		log.Fatal(err)
	}
	dbmBars.Color = plotutil.Color(0)
	dbmBars.Offset = -w

	// MBM收率柱状图
	mbmBars, err := plotter.NewBarChart(plotter.Values(mbmYields), w)
	if err != nil {
		log.Fatal(err)
	}
	mbmBars.Color = plotutil.Color(1)

	// 循环5次后活性柱状图
	cycleBars, err := plotter.NewBarChart(plotter.Values(after5Cycles), w)
	if err != nil {
		log.Fatal(err)
	}
	cycleBars.Color = plotutil.Color(2)
	cycleBars.Offset = w

	// 添加到图表
	p.Add(dbmBars, mbmBars, cycleBars)
	p.Legend.Add("DBM Yield", dbmBars)
	p.Legend.Add("MBM Yield", mbmBars)
	p.Legend.Add("After 5 cycles", cycleBars)
	p.Legend.TextStyle.Font.Size = 14
	p.NominalX(catalysts...)

	// 添加数值标签
	for i, val := range dbmYields {
		if val > 0 {
			l, err := plotter.NewLabels(plotter.XYLabels{
				XYs:    []plotter.XY{{X: float64(i), Y: val + 1}},
				Labels: []string{formatFloat(val)},
			})
			if err == nil {
				// 设置标签样式
				for j := range l.TextStyle {
					l.TextStyle[j].Font.Size = 12
				}
				p.Add(l)
			}
		}
	}

	for i, val := range after5Cycles {
		if val > 0 {
			l, err := plotter.NewLabels(plotter.XYLabels{
				XYs:    []plotter.XY{{X: float64(i), Y: val + 1}},
				Labels: []string{formatFloat(val)},
			})
			if err == nil {
				// 设置标签样式
				for j := range l.TextStyle {
					l.TextStyle[j].Font.Size = 12
				}
				p.Add(l)
			}
		}
	}

	// 保存大尺寸图表 (10x8英寸)
	if err := p.Save(10*vg.Inch, 8*vg.Inch, "figure1_catalyst_comparison.png"); err != nil {
		log.Fatal(err)
	}
}

// 图2：DFT计算闭环能垒图
func createDFTEnergyBarrierPlot() {
	p := plot.New()
	p.Title.Text = "DFT Calculation of Energy Barrier"
	p.Title.TextStyle.Font.Size = 20
	p.Title.Padding = 10 * vg.Millimeter

	p.X.Label.Text = "Reaction Coordinate"
	p.X.Label.TextStyle.Font.Size = 16
	p.X.Tick.Label.Font.Size = 14
	p.Y.Label.Text = "Energy (kJ/mol)"
	p.Y.Label.TextStyle.Font.Size = 16
	p.Y.Tick.Label.Font.Size = 14

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

	// 找到能垒最高点
	maxWithout := findMaxPoint(withoutCatalyst)
	maxWith := findMaxPoint(withCatalyst)

	// 创建线条
	line1, err := plotter.NewLine(withoutCatalyst)
	if err != nil {
		log.Fatal(err)
	}
	line1.LineStyle.Width = vg.Points(2)
	line1.LineStyle.Color = plotutil.Color(0)

	line2, err := plotter.NewLine(withCatalyst)
	if err != nil {
		log.Fatal(err)
	}
	line2.LineStyle.Width = vg.Points(2)
	line2.LineStyle.Color = plotutil.Color(1)

	// 添加散点
	scatter1, err := plotter.NewScatter(withoutCatalyst)
	if err != nil {
		log.Fatal(err)
	}
	scatter1.GlyphStyle.Color = plotutil.Color(0)
	scatter1.GlyphStyle.Radius = vg.Points(1.5)

	scatter2, err := plotter.NewScatter(withCatalyst)
	if err != nil {
		log.Fatal(err)
	}
	scatter2.GlyphStyle.Color = plotutil.Color(1)
	scatter2.GlyphStyle.Radius = vg.Points(1.5)

	// 添加到图表
	p.Add(line1, line2, scatter1, scatter2)
	p.Legend.Add("Without Catalyst", line1)
	p.Legend.Add("With AlCl₃ Catalyst", line2)
	p.Legend.TextStyle.Font.Size = 14

	// 添加能垒值标注
	label1, err := plotter.NewLabels(plotter.XYLabels{
		XYs:    []plotter.XY{{X: maxWithout.X, Y: maxWithout.Y + 2}},
		Labels: []string{"Barrier: " + formatFloat(maxWithout.Y) + " kJ/mol"},
	})
	if err == nil {
		// 设置标签样式
		for i := range label1.TextStyle {
			label1.TextStyle[i].Font.Size = 12
		}
		p.Add(label1)
	}

	label2, err := plotter.NewLabels(plotter.XYLabels{
		XYs:    []plotter.XY{{X: maxWith.X, Y: maxWith.Y + 2}},
		Labels: []string{"Barrier: " + formatFloat(maxWith.Y) + " kJ/mol"},
	})
	if err == nil {
		// 设置标签样式
		for i := range label2.TextStyle {
			label2.TextStyle[i].Font.Size = 12
		}
		p.Add(label2)
	}

	// 添加能垒降低标注
	barrierReduction := maxWithout.Y - maxWith.Y
	label3, err := plotter.NewLabels(plotter.XYLabels{
		XYs:    []plotter.XY{{X: (maxWithout.X + maxWith.X) / 2, Y: (maxWithout.Y + maxWith.Y) / 2}},
		Labels: []string{"Barrier Reduction: " + formatFloat(barrierReduction) + " kJ/mol"},
	})
	if err == nil {
		// 设置标签样式
		for i := range label3.TextStyle {
			label3.TextStyle[i].Font.Size = 14
			label3.TextStyle[i].Color = plotutil.Color(2)
		}
		p.Add(label3)
	}

	// 保存大尺寸图表
	if err := p.Save(10*vg.Inch, 8*vg.Inch, "figure2_dft_energy_barrier.png"); err != nil {
		log.Fatal(err)
	}
}

// 图3：催化剂失活与再生图
func createDeactivationRegenerationPlot() {
	p := plot.New()
	p.Title.Text = "Catalyst Deactivation and Regeneration"
	p.Title.TextStyle.Font.Size = 20
	p.Title.Padding = 10 * vg.Millimeter

	p.X.Label.Text = "Time (h)"
	p.X.Label.TextStyle.Font.Size = 16
	p.X.Tick.Label.Font.Size = 14
	p.Y.Label.Text = "Relative Activity (%)"
	p.Y.Label.TextStyle.Font.Size = 16
	p.Y.Tick.Label.Font.Size = 14

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
	line.LineStyle.Width = vg.Points(3)
	line.LineStyle.Color = plotutil.Color(0)

	// 创建散点
	scatter, err := plotter.NewScatter(points)
	if err != nil {
		log.Fatal(err)
	}
	scatter.GlyphStyle.Color = plotutil.Color(0)
	scatter.GlyphStyle.Radius = vg.Points(4)

	// 添加再生点特殊标记
	regenerationPoint := plotter.XY{X: 1000, Y: 95}
	regenerationScatter, err := plotter.NewScatter(plotter.XYs{regenerationPoint})
	if err != nil {
		log.Fatal(err)
	}
	regenerationScatter.GlyphStyle.Color = plotutil.Color(2)
	regenerationScatter.GlyphStyle.Radius = vg.Points(6)
	regenerationScatter.GlyphStyle.Shape = draw.PyramidGlyph{}

	// 添加到图表
	p.Add(line, scatter, regenerationScatter)

	// 添加数据点标签
	for _, pt := range points {
		l, err := plotter.NewLabels(plotter.XYLabels{
			XYs:    []plotter.XY{{X: pt.X, Y: pt.Y + 1}},
			Labels: []string{formatFloat(pt.Y)},
		})
		if err == nil {
			// 设置标签样式
			for i := range l.TextStyle {
				l.TextStyle[i].Font.Size = 12
			}
			p.Add(l)
		}
	}

	// 添加再生标注
	label, err := plotter.NewLabels(plotter.XYLabels{
		XYs:    []plotter.XY{{X: 1000, Y: 97}},
		Labels: []string{"Regeneration:\n5% HCl wash\n200°C calcination"},
	})
	if err == nil {
		// 设置标签样式
		for i := range label.TextStyle {
			label.TextStyle[i].Font.Size = 14
		}
		p.Add(label)
	}

	// 设置Y轴范围
	p.Y.Min = 90
	p.Y.Max = 102

	// 保存大尺寸图表
	if err := p.Save(10*vg.Inch, 8*vg.Inch, "figure3_deactivation_regeneration.png"); err != nil {
		log.Fatal(err)
	}
}

// 图4：Pareto图（因素贡献度）
func createParetoChart() {
	// 创建图表
	p := plot.New()
	p.Title.Text = "Pareto Chart of Factor Contributions"
	p.Title.TextStyle.Font.Size = 20
	p.Title.Padding = 10 * vg.Millimeter

	p.X.Label.Text = "Factors"
	p.X.Label.TextStyle.Font.Size = 16
	p.X.Tick.Label.Font.Size = 14
	p.Y.Label.Text = "Contribution (%)"
	p.Y.Label.TextStyle.Font.Size = 16
	p.Y.Tick.Label.Font.Size = 14

	// 因素和贡献度数据
	factors := []string{"A", "B", "C", "D", "E"}
	contributions := []float64{45, 30, 25, 15, 10} // 百分比

	// 创建柱状图
	w := vg.Points(15)
	bars, err := plotter.NewBarChart(plotter.Values(contributions), w)
	if err != nil {
		log.Fatal(err)
	}
	bars.Color = plotutil.Color(0)

	// 计算累积贡献度
	cumulative := make(plotter.XYs, len(contributions))
	cumulative[0].X = 0.5
	cumulative[0].Y = contributions[0]
	for i := 1; i < len(contributions); i++ {
		cumulative[i].X = float64(i) + 0.5
		cumulative[i].Y = cumulative[i-1].Y + contributions[i]
	}

	// 创建累积贡献度折线
	line, err := plotter.NewLine(cumulative)
	if err != nil {
		log.Fatal(err)
	}
	line.LineStyle.Width = vg.Points(3)
	line.LineStyle.Color = plotutil.Color(1)

	// 添加散点
	scatter, err := plotter.NewScatter(cumulative)
	if err != nil {
		log.Fatal(err)
	}
	scatter.GlyphStyle.Color = plotutil.Color(1)
	scatter.GlyphStyle.Radius = vg.Points(4)

	// 添加到图表
	p.Add(bars, line, scatter)
	p.Legend.Add("Contribution", bars)
	p.Legend.Add("Cumulative", line)
	p.Legend.TextStyle.Font.Size = 14
	p.NominalX(factors...)

	// 添加贡献度标签
	for i, val := range contributions {
		l, err := plotter.NewLabels(plotter.XYLabels{
			XYs:    []plotter.XY{{X: float64(i) + 0.5, Y: val + 1}},
			Labels: []string{formatFloat(val)},
		})
		if err == nil {
			// 设置标签样式
			for j := range l.TextStyle {
				l.TextStyle[j].Font.Size = 12
			}
			p.Add(l)
		}
	}

	// 添加累积贡献度标签
	for _, pt := range cumulative {
		l, err := plotter.NewLabels(plotter.XYLabels{
			XYs:    []plotter.XY{{X: pt.X, Y: pt.Y + 1}},
			Labels: []string{formatFloat(pt.Y)},
		})
		if err == nil {
			// 设置标签样式
			for j := range l.TextStyle {
				l.TextStyle[j].Font.Size = 12
				l.TextStyle[j].Color = plotutil.Color(1)
			}
			p.Add(l)
		}
	}

	// 添加70%参考线
	refLine := plotter.NewFunction(func(x float64) float64 { return 70 })
	refLine.LineStyle.Width = vg.Points(1)
	refLine.LineStyle.Color = plotutil.Color(2)
	refLine.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	p.Add(refLine)

	// 添加参考线标签
	label, err := plotter.NewLabels(plotter.XYLabels{
		XYs:    []plotter.XY{{X: 0, Y: 72}},
		Labels: []string{"70% Reference Line"},
	})
	if err == nil {
		// 设置标签样式
		for i := range label.TextStyle {
			label.TextStyle[i].Font.Size = 14
			label.TextStyle[i].Color = plotutil.Color(2)
		}
		p.Add(label)
	}

	// 添加A、B、C贡献度>70%的标注
	label2, err := plotter.NewLabels(plotter.XYLabels{
		XYs:    []plotter.XY{{X: 1.5, Y: 85}},
		Labels: []string{"A, B, C contribution >70%"},
	})
	if err == nil {
		// 设置标签样式
		for i := range label2.TextStyle {
			label2.TextStyle[i].Font.Size = 14
		}
		p.Add(label2)
	}

	// 保存大尺寸图表
	if err := p.Save(10*vg.Inch, 8*vg.Inch, "figure4_pareto_chart.png"); err != nil {
		log.Fatal(err)
	}
}

// 辅助函数：找到最大值点
func findMaxPoint(points plotter.XYs) plotter.XY {
	max := points[0]
	for _, p := range points {
		if p.Y > max.Y {
			max = p
		}
	}
	return max
}

// 辅助函数：格式化浮点数
func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 1, 64)
}
