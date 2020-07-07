package main

import (
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/go-echarts/go-echarts/charts"
)

/*
 *
 *
 *
 */
const (
	host          = ":8081"
	lineChartFile = "line.html"
	maxXAxisValue = 100
	cacheChart    = false
)

/*
 *
 *
 *
 */
func main() {
	http.HandleFunc("/", logTracing(logarithmHandler))
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

func binaryLogData() []float64 {
	r := make([]float64, maxXAxisValue)
	for i := 1; i <= maxXAxisValue; i++ {
		r[i-1] = math.Log2(float64(i))
	}

	return r
}

func xAxisData() []string {
	s := make([]string, maxXAxisValue)
	for i := 1; i <= maxXAxisValue; i++ {
		s[i-1] = strconv.Itoa(i)
	}
	return s
}

func commonLogData() []float64 {
	r := make([]float64, maxXAxisValue)
	for i := 1; i <= maxXAxisValue; i++ {
		r[i-1] = math.Log10(float64(i))
	}

	return r
}

func naturalLogData() []float64 {
	r := make([]float64, maxXAxisValue)
	for i := 1; i <= maxXAxisValue; i++ {
		r[i-1] = math.Log(float64(i))
	}
	return r
}

/*	note: charts.ToolboxOpts{Show: true})
 *
 *
 *
 */
func logarithmHandler(w http.ResponseWriter, _ *http.Request) {

	var f *os.File
	var line *charts.Line

	if cacheChart {
		_, err := os.Stat(lineChartFile)
		if os.IsNotExist(err) {
			f, line = createChart()
			line.Render(w, f)
		} else {
			t, _ := template.ParseFiles(lineChartFile)
			log.Printf("Executing the rendered template\n")
			t.Execute(w, nil)
		}
	} else {
		f, line = createChart()
		line.Render(w, f)
	}
}

func createChart() (*os.File, *charts.Line) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.TitleOpts{Title: "Logarithm Chart"}, charts.ToolboxOpts{Show: false},
		charts.InitOpts{
			PageTitle:       "Logarithm Charts",
			BackgroundColor: "white",
			// Theme:           charts.ThemeType.Vintage,
		},
	)

	xAxisNames := xAxisData()
	binaryLogVals := binaryLogData()
	commonLogVals := commonLogData()
	naturalLogVals := naturalLogData()

	line.AddXAxis(xAxisNames).
		AddYAxis("Binary Logarithm, y = log₂(x)", binaryLogVals).
		AddYAxis("Common Logarithm, y = log\u2081\u2080(x)", commonLogVals).
		AddYAxis("Natural Logarithm, y = log\u2091(x)", naturalLogVals)

	f, err := os.Create(lineChartFile)
	if err != nil {
		log.Println(err)
	}
	return f, line
}
