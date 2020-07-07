package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
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
	http.ListenAndServe(host, nil)
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
