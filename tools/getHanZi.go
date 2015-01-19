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
	fmt.Println("")
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
		//dict = strings.TrimSpace(dict)
		reg1 := regexp.MustCompile("[\\n\\r]") //删除末尾的换行符，不删除空格是为了保留代码的缩进。
		dict = reg1.ReplaceAllString(dict, "")

		if err != nil {
			// 文件结束
			break
		}

		if len(dict) == 0 {
			// 空行
			continue
		}

		//行首为*的是多行注释，删除掉。	这个还有待验证。只有php时可能正确。html时就不一定正确。
		var hzRegexp = regexp.MustCompile(`^\*`)
		if hzRegexp.MatchString(dict) {
			continue
		}
		hzRegexp = regexp.MustCompile(`^\s+\*`) //行首为：多个空格+*的是多行注释的开始，删除掉。
		if hzRegexp.MatchString(dict) {
			continue
		}
		hzRegexp = regexp.MustCompile(`^\/\*`) //行首为/*的是多行注释的开始，删除掉。
		if hzRegexp.MatchString(dict) {
			continue
		}
		hzRegexp = regexp.MustCompile(`^\/\/`) //行首为//的是单行注释的开始，删除掉。
		if hzRegexp.MatchString(dict) {
			continue
		}
		hzRegexp = regexp.MustCompile(`^\/\/\s+`) //行首为//+多个空格开始的是单行注释的开始，删除掉。
		if hzRegexp.MatchString(dict) {
			continue
		}

		//匹配//后面出现的汉字，这个替换对html文档就可能出错了，比如url里面有
		reg := regexp.MustCompile("//[\u0391-\uFFE5]+.*")
		lineConent := reg.ReplaceAllString(dict, "")

		re, _ := regexp.Compile("[\u0391-\uFFE5]+")
		ret := re.FindAllStringSubmatch(lineConent, -1)

		//还应该对同一行的文字进行去重。
		lineConent = strings.Replace(dict, ",", "@@", -1)
		if len(ret) > 0 {
			//匹配所有的情况，主要是js
			if len(ret) == 1 && ret[0][0] == "：" {

			} else {
				fmt.Print(lineConent + "," + path + "," + strconv.Itoa(line) + ",")
				fmt.Println(ret)
			}

			/*
				//匹配一行出现多个连续汉字的。
				if len(ret) > 1 {
					fmt.Print(lineConent + "," + path + "," + strconv.Itoa(line) + ",")
					fmt.Println(ret[0])
				}
			*/
			/*
				每个连续出现的汉字是一行。
				for i := 0; i < len(ret); i++ {
					if i == 0 {
						fmt.Println(ret[i])
					} else {
						fmt.Print(",,,")
						fmt.Println(ret[i])
					}
				}
			*/

			/*
				//每行只出现一次连续的汉字提取出来
				if len(ret) == 1 {
					fmt.Print(lineConent + "," + path + "," + strconv.Itoa(line) + ",")
					fmt.Println(ret[0])
				}
			*/
		}
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
