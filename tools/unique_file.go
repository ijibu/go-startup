//根据文件的行内容进行去重。
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	FILENAME = `./dazhe.log`
)

var (
	docMap map[string]int
)

func main() {
	docMap = map[string]int{}
	fileHander, err := os.Open(FILENAME)
	defer fileHander.Close()
	if err != nil {
		log.Fatalf("无法打开文件 \"%s\" \n", FILENAME)
	}

	reader := bufio.NewReader(fileHander)

	// 逐行读入分词
	for {
		lineContent, _ := reader.ReadString('\n') //行内容
		lineContent = strings.TrimSpace(lineContent)

		if len(lineContent) == 0 {
			// 文件结束
			break
		}
		_, ok := docMap[lineContent]
		if !ok {
			docMap[lineContent] = 0
			fmt.Println(lineContent)
		}
	}
}
