//批量替换文件中的内容。
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

var fileName *string = flag.String("f", "Null", "please input a fileName like ./omc1.csv")

type repace struct {
	line    int
	content string
}
type repaces []repace

var item repace

var repaceContents map[string]repaces

func main() {
	flag.Parse()
	if *fileName == "Null" {
		show_usage()
		return
	}
	parseReFile(fileName)
	repaceContents = map[string]repaces{}
}

func parseReFile(path string) {
	dictFile, err := os.Open(path)
	defer dictFile.Close()
	if err != nil {
		log.Fatalf("无法打开文件 \"%s\" \n", path)
	}

	reader := bufio.NewReader(dictFile)
	line := 0

	// 逐行读入分词
	for {
		line++
		dict, err := reader.ReadString('\n')
		dict = strings.TrimSpace(dict)
		arr := strings.Split("@@")

		if err != nil {
			// 文件结束
			break
		}

		item = repace{line: arr[2], content: arr[0]}
		_, ok := repaceContents[arr[1]]
		if !ok {

		} else {

		}
	}
}
