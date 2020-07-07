package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/go-echarts/go-echarts/charts"
)

const (
	host          = ":8081"
	maxNum        = 50
	lineChartFile = "bar.html"
)

/*
 *
 *
 *
 */
var nameItems = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
var seed = rand.NewSource(time.Now().UnixNano())

/*
 *
 *
 *
 */
func main() {
	http.HandleFunc("/", chartHandler)
	http.HandleFunc("/line", logTracing(lineHandler))
	http.ListenAndServe(host, nil)
}

func logTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

/*
 *
 *
 *
 */
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

/*
 *
 *
 *
 */
func chartHandler(w http.ResponseWriter, _ *http.Request) {

	// if !fileExists(lineChartFile) {
	// need a *File
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(w, f) // Render 可接收多个 io.Writer 接口
	// }

}

/*
 *
 *
 *
 */
func randInt() []int {
	cnt := len(nameItems)
	r := make([]int, 0)
	for i := 0; i < cnt; i++ {
		r = append(r, int(seed.Int63())%maxNum)
	}
	return r
}

func lineMulti() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-多线"}, charts.InitOpts{Theme: "shine"})
	line.AddXAxis(nameItems).
		AddYAxis("商家 A", randInt()).
		AddYAxis("商家 B", randInt()).
		AddYAxis("商家 C", randInt()).
		AddYAxis("商家 D", randInt())
	return line
}

func lineHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("line")...)
	page.Add(
		lineMulti(),
	)
	f, err := os.Create(getRenderPath("line.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}

func getRenderPath(f string) string {
	return path.Join("html", f)
}

func orderRouters(chartType string) []charts.RouterOpts {
	for i := 0; i < len(routers); i++ {
		if routers[i].name == chartType {
			routers[i], routers[0] = routers[0], routers[i]
			break
		}
	}

	rs := make([]charts.RouterOpts, 0)
	for i := 0; i < len(routers); i++ {
		rs = append(rs, routers[i].RouterOpts)
	}
	return rs
}

type router struct {
	name string
	charts.RouterOpts
}

var (
	routers = []router{
		// {"bar", charts.RouterOpts{URL: host + "/bar", Text: "Bar-(柱状图)"}},
		// {"bar3D", charts.RouterOpts{URL: host + "/bar3D", Text: "Bar3D-(3D 柱状图)"}},
		// {"boxPlot", charts.RouterOpts{URL: host + "/boxPlot", Text: "BoxPlot-(箱线图)"}},
		// {"effectScatter", charts.RouterOpts{URL: host + "/effectScatter", Text: "EffectScatter-(动态散点图)"}},
		// {"funnel", charts.RouterOpts{URL: host + "/funnel", Text: "Funnel-(漏斗图)"}},
		// {"gauge", charts.RouterOpts{URL: host + "/gauge", Text: "Gauge-仪表盘"}},
		// {"geo", charts.RouterOpts{URL: host + "/geo", Text: "Geo-地理坐标系"}},
		// {"graph", charts.RouterOpts{URL: host + "/graph", Text: "Graph-关系图"}},
		// {"heatMap", charts.RouterOpts{URL: host + "/heatMap", Text: "HeatMap-热力图"}},
		// {"kline", charts.RouterOpts{URL: host + "/kline", Text: "Kline-K 线图"}},
		{"line", charts.RouterOpts{URL: host + "/line", Text: "Line-(折线图)"}},
		// {"line3D", charts.RouterOpts{URL: host + "/line3D", Text: "Line3D-(3D 折线图)"}},
		// {"liquid", charts.RouterOpts{URL: host + "/liquid", Text: "Liquid-(水球图)"}},
		// {"map", charts.RouterOpts{URL: host + "/map", Text: "Map-(地图)"}},
		// {"overlap", charts.RouterOpts{URL: host + "/overlap", Text: "Overlap-(重叠图)"}},
		// {"parallel", charts.RouterOpts{URL: host + "/parallel", Text: "Parallel-(平行坐标系)"}},
		// {"pie", charts.RouterOpts{URL: host + "/pie", Text: "Pie-(饼图)"}},
		// {"radar", charts.RouterOpts{URL: host + "/radar", Text: "Radar-(雷达图)"}},
		// {"sankey", charts.RouterOpts{URL: host + "/sankey", Text: "Sankey-(桑基图)"}},
		// {"scatter", charts.RouterOpts{URL: host + "/scatter", Text: "Scatter-(散点图)"}},
		// {"scatter3D", charts.RouterOpts{URL: host + "/scatter3D", Text: "Scatter-(3D 散点图)"}},
		// {"surface3D", charts.RouterOpts{URL: host + "/surface3D", Text: "Surface3D-(3D 曲面图)"}},
		// {"themeRiver", charts.RouterOpts{URL: host + "/themeRiver", Text: "ThemeRiver-(主题河流图)"}},
		// {"wordCloud", charts.RouterOpts{URL: host + "/wordCloud", Text: "WordCloud-(词云图)"}},
		// {"page", charts.RouterOpts{URL: host + "/page", Text: "Page-(顺序多图)"}},
	}
)

func lineSplitLine() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-显示分割线"})
	line.AddXAxis(nameItems).AddYAxis("商家A", randInt(), charts.LabelTextOpts{Show: true})
	line.SetGlobalOptions(charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}})
	return line
}
