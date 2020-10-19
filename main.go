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

	// 运行
	t1 := time.Now().UnixNano()
	err := ExcelToCsv(excelPath, csvPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	t2 := time.Now().UnixNano()
	fmt.Println(excelPath, "->", csvPath, (t2-t1)/1e6, "ms")
}

// 打印帮助信息
func PrintUsage() {
	fmt.Println("excelToCsv.exe [EXCEL_PATH] [CSV_PATH]")
}
