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

/*
 *
 *
 *
 */
const (
	host          = ":8081"
	maxNum        = 50
	lineChartFile = "line.html"
	barChartFile  = "bar.html"
)

/*
 *
 *
 *
 */
type router struct {
	name string
	charts.RouterOpts
}

/*
 *
 *
 *
 */
var (
	routers = []router{
		{"line", charts.RouterOpts{URL: host + "/line", Text: "Line"}},
	}
	nameItems = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	seed      = rand.NewSource(time.Now().UnixNano())
)

/*
 *
 *
 *
 */
func main() {
	http.HandleFunc("/", logTracing(singleLineChartHandler))
	http.HandleFunc("/line", logTracing(multiLineChart))
	http.ListenAndServe(host, nil)
}

/*
 *
 *
 *
 */
func logTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

func binaryLogData() {

}

/*	note: charts.ToolboxOpts{Show: true})
 *
 *
 *
 */
func singleLineChartHandler(w http.ResponseWriter, _ *http.Request) {

	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Logarithm Chart"}, charts.ToolboxOpts{Show: false})

	nameItems := []string{"1", "2", "3", "4", "5", "6", "7", "8"}

	// line.AddXAxis(nameItems).AddYAxis("商家A", randInt()).AddYAxis("商家B", randInt())

	values := []float64{0, 1, 1.5849, 2, 2.3219, 2.5849, 2.8073, 3}

	// nameItems := []string{"1", "2", "4", "8"}
	// values := []float64{0, 1, 2, 3}

	line.AddXAxis(nameItems).AddYAxis("Binary Logarithm, y = log₂(x)", values)

	// smooth the line
	line.SetSeriesOptions(
	// charts.LineOpts{Smooth: true},
	)

	f, err := os.Create("line.html")
	if err != nil {
		log.Println(err)
	}
	line.Render(w, f)
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

/*
 *
 *
 *
 */
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

/*
 *
 *
 *
 */
func multiLineChart(w http.ResponseWriter, _ *http.Request) {
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

/*
 *
 *
 *
 */
func getRenderPath(f string) string {
	return path.Join("html", f)
}

/*	Parameter:	string
 *	Return:		[]charts.RouterOpts
 *
 *
 */
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

/*
 *
 *
 *
 */
func lineSplitLine() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-显示分割线"})
	line.AddXAxis(nameItems).AddYAxis("商家A", randInt(), charts.LabelTextOpts{Show: true})
	line.SetGlobalOptions(charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}})
	return line
}

/*
 *
 *
 *
 */
// func fileExists(filename string) bool {
// 	info, err := os.Stat(filename)
// 	if os.IsNotExist(err) {
// 		return false
// 	}
// 	return !info.IsDir()
// }
