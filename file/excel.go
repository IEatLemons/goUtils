package file

import "github.com/xuri/excelize/v2"

type Excel struct {
	Path string
	File *excelize.File
	Data []*ExcelSheet
}

type ExcelSheet struct {
	Name  string
	Cells []*ExcelCell
}

type ExcelCell struct {
	Name  string
	Value interface{}
}

func CreateExcel(Data *Excel) (err error) {
	Data.File = excelize.NewFile()
	defer func() {
		if err := Data.File.Close(); err != nil {
			return
		}
	}()
	for _, Sheet := range Data.Data {
		index, err := Data.File.NewSheet(Sheet.Name)
		if err != nil {
			return err
		}
		for _, Cell := range Sheet.Cells {
			Data.File.SetCellValue(Sheet.Name, Cell.Name, Cell.Value)
		}
		Data.File.SetActiveSheet(index)
	}
	return Data.File.SaveAs(Data.Path)
}
