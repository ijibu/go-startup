//批量抓取CSS文件中的图片，并与网站的目录一致。
package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	UA       = "Mozilla/5.0 (Windows NT 6.1; rv:28.0) Gecko/20100101 Firefox/28.0"
	HOST     = "http://www.fxtrip.com"
	SAVEPATH = "download"
)

var fileName *string = flag.String("f", "Null", "please input a fileName like ./a.css")

func main() {
	flag.Parse()
	if *fileName == "Null" {
		show_usage()
		return
	}
	parse(*fileName)
}

//处理html文件
func parse(path string) {
	parseHtmlFile(path)
}

func parseHtmlFile(path string) {
	//读取整个文件的内容
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	content1 := string(content)
	//fmt.Print(content1)
	//解析阅读全文地址。
	re, _ := regexp.Compile(`url\([\"|']?([\S\s]+?)[\"|']?\)`)
	ret := re.FindAllStringSubmatch(content1, -1)
	//fmt.Print(ret)

	for i := 0; i < len(ret); i++ {
		downloadUrl := ret[i][1]
		download(downloadUrl)
	}
}

func download(downUrl1 string) {
	str := strings.Split(downUrl1, ".")

	dirPath := SAVEPATH + str[0]
	fileName := SAVEPATH + downUrl1

	var downUrl string
	if downUrl1[0:1] == "/" {
		downUrl = HOST + downUrl1
	} else {
		downUrl = downUrl1
	}
	fmt.Print(downUrl)
	fmt.Print("\n")

	if !isDirExists(dirPath) { //目录不存在，则进行创建。
		err := os.MkdirAll(dirPath, 777) //递归创建目录，linux下面还要考虑目录的权限设置。
		if err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666) //其实这里的 O_RDWR应该是 O_RDWR|O_CREATE，也就是文件不存在的情况下就建一个空文件，但是因为windows下还有BUG，如果使用这个O_CREATE，就会直接清空文件，所以这里就不用了这个标志，你自己事先建立好文件。
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var req http.Request
	req.Method = "GET"
	req.Close = true
	req.URL, _ = url.Parse(downUrl)

	header := http.Header{}
	header.Set("User-Agent", UA)
	header.Set("Host", HOST)
	req.Header = header
	resp, err := http.DefaultClient.Do(&req)
	if err == nil {
		if resp.StatusCode == 200 {
			fmt.Println(":sucess")
			_, err = io.Copy(f, resp.Body)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println(strconv.Itoa(resp.StatusCode))
		}
		defer resp.Body.Close()
	} else {
		fmt.Println(":error")
	}
}

func isDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

func show_usage() {
	fmt.Fprintf(os.Stderr,
		"Usage: %s [-f=<filename>]\n"+
			"       <command> [<args>]\n\n",
		os.Args[0])
	fmt.Fprintf(os.Stderr,
		"Flags:\n")
	flag.PrintDefaults()
}
