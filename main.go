// 将xlsx转换成csv
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// 解析命令行参数
	flag.Parse()
	if len(flag.Args()) != 2 {
		PrintUsage()
		return
	}

	excelPath := flag.Args()[0] // xlsx文件路径
	csvPath := flag.Args()[1]   // csv文件路径
	fmt.Println(excelPath, "->", csvPath)

	// 运行
	t1 := time.Now().UnixNano() / 1e3
	ExcelToCsv(excelPath, csvPath)
	t2 := time.Now().Unix() / 1e3
	fmt.Println("耗时：", t2-t1)
}

// 打印帮助信息
func PrintUsage() {
	fmt.Println("excelToCsv.exe [EXCEL_PATH] [CSV_PATH]")
}
