package main

import (
	"encoding/csv"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// 将xlsx转换成csv
func ExcelToCsv(xlsxPath string, csvPath string) error {
	// 读xlsx文件
	xlsxFile, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		return err
	}

	// 读取第一张sheet
	firstSheetName := xlsxFile.GetSheetName(1)
	rows := xlsxFile.GetRows(firstSheetName)

	// 创建csv文件
	csvFile, err := os.Create(csvPath)
	if err != nil {
		return err
	}

	defer csvFile.Close()
	csvFile.WriteString("\xEF\xBB\xBF") // 写入一个UTF-8 BOM
	csvWriter := csv.NewWriter(csvFile) //创建一个新的写入文件流
	csvWriter.WriteAll(rows)
	csvWriter.Flush()
	return nil
}
