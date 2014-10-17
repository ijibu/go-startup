//批量提取目录下所有文件中出现的中文字符。公司项目要做多语言，之前没考虑到，所以要把代码中的所有中文提取出来。
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var dirPath *string = flag.String("d", "Null", "please input a dirPath like ./data/")

func main() {
	flag.Parse()
	if *dirPath == "Null" {
		show_usage()
		return
	}
	parse()
}

//处理html文件
func parse() {
	filepath.Walk(*dirPath, func(path string, f os.FileInfo, e error) error {
		if f == nil {
			return e
		}
		if f.IsDir() {
			return nil
		}
		extName := filepath.Ext(path)

		if extName == ".js" || extName == ".php" || extName == ".html" {
			parseHtmlFile(path)
		}

		return nil
	})
}

func parseHtmlFile(path string) {
	dictFile, err := os.Open(path)
	defer dictFile.Close()
	if err != nil {
		log.Fatalf("无法打开文件 \"%s\" \n", path)
	}

	reader := bufio.NewReader(dictFile)
	line := 1

	// 逐行读入分词
	for {
		dict, _ := reader.ReadString('\n')
		dict = strings.TrimSpace(dict)

		if len(dict) == 0 {
			// 文件结束
			break
		}
		re, _ := regexp.Compile("[\u0391-\uFFE5]+")
		ret := re.FindAllStringSubmatch(dict, -1)
		lineConent := strings.Replace(dict, ",", "@@", -1)
		if len(ret) > 0 {
			fmt.Print(lineConent + "," + path + "," + strconv.Itoa(line) + ",")
			//str := strings.Join(ret[0], "\r\n")
			fmt.Println(ret)
		}
		line++
	}
}

func show_usage() {
	fmt.Fprintf(os.Stderr,
		"Usage: %s [-d=<dirPath>]\n"+
			"       <command> [<args>]\n\n",
		os.Args[0])
	fmt.Fprintf(os.Stderr,
		"Flags:\n")
	flag.PrintDefaults()
}
