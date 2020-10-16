// 将xlsx转化成csv
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

	xlsxPath := flag.Args()[0] // xlsx文件路径
	csvPath := flag.Args()[1]  // csv文件路径
	fmt.Println(xlsxPath, "->", csvPath)

	t1 := time.Now().Unix()
	ExcelToCsv(xlsxPath, csvPath)
	t2 := time.Now().Unix()
	fmt.Println("耗时", (t2 - t1))

	fmt.Println(time.Now().Unix())
}

// 打印帮助信息
func PrintUsage() {
	fmt.Println("excelToCsv.exe [XLSX_PATH] [CSV_PATH]")
}
