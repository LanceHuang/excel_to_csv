package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// 将xlsx转换成csv
// @param xlsxPath xlsx文件路径
// @param csvPath csv文件路径
func ExcelToCsv(xlsxPath string, csvPath string) {
	// 读xlsx文件
	xlsxFile, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读取第一张sheet
	firstSheetName := xlsxFile.GetSheetName(1)
	rows := xlsxFile.GetRows(firstSheetName)
	// for _, row := range rows {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }

	// 创建csv文件
	csvFile, err := os.Create(csvPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvFile.Close()
	csvFile.WriteString("\xEF\xBB\xBF") // 写入一个UTF-8 BOM
	csvWriter := csv.NewWriter(csvFile) //创建一个新的写入文件流
	csvWriter.WriteAll(rows)
	csvWriter.Flush()
}
