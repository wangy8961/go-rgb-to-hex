package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// 通过命令行参数，提供包含 RGB 的字符串
	rgb := flag.String("r", "", "请指定包含 Red,Green,Blue 的字符串，颜色值用逗号隔开即可，可包含其它字符，比如 'rgb(255,97,171);' 也能正常解析")
	// 提供 -h 选项，查看命令行帮助信息
	help := flag.Bool("h", false, "显示帮助信息")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *rgb == "" {
		log.Fatal("必须指定包含 RGB 的字符串")
	}

	// 正则表达式
	digitsRegexp := regexp.MustCompile(`\d+`)
	// 匹配到所有包含数字的字符子串，比如 [255 97 171]
	rgbSlice := digitsRegexp.FindAllString(*rgb, -1)
	// 转换为十六进制
	hexString := "#"
	for _, v := range rgbSlice {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		hexString += fmt.Sprintf("%X", i)
	}
	fmt.Printf("Hex conv of '%s' is: %s\n", *rgb, hexString)
}
