package main

import (
	"fmt"
	"net/http"
	"os"

	"projects/charcountv2/browser"
	"projects/charcountv2/counter"
	"projects/charcountv2/template"
	"projects/office"
)

func main() {
	// 打开主页
	http.HandleFunc("/", indexPage)

	// 上传数据
	http.HandleFunc("/post", postPage)

	// 返回结果
	http.HandleFunc("/download", downloadPage)

	// 打开浏览器
	port := "10005"
	url := "http://localhost:" + port
	go openBrowser(url)

	// 运行
	host := fmt.Sprintf(":%s", port)
	http.ListenAndServe(host, nil)
}

func openBrowser(url string) {
	err := browser.RetardOpen(url)
	if err != nil {
		panic(err)
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, template.GetIndex())
}

// 计数结果
var result [][]string

func postPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	direcSli := r.Form["direc"]
	var direc string

	var alert string
	if len(direcSli) == 0 {
		alert = "没有路径被输入"
	} else if len(direcSli) > 1 {
		alert = "输入路径有误"
	} else if len(direcSli) == 1 {
		direc = direcSli[0]
		fInfo, err := os.Stat(direc)
		if err != nil {
			alert = "输入路径不存在"
		} else if !fInfo.IsDir() {
			alert = "输入路径不是文件夹"
		}
	}

	if alert != "" {
		fmt.Fprintf(w, template.GetAlert(alert))
	} else {
		result = counter.CountDir(direc)
		fmt.Fprintf(w, template.GetReply())
	}
}

func downloadPage(w http.ResponseWriter, r *http.Request) {
	columns := []string{"file_name", "ext_name", "file_path", "word_number", "Remarks"}
	w.Header().Add("Content-Disposition", "attachment; filename=result.xlsx")
	office.WriteExcelToWriter(w, result, columns)
}
