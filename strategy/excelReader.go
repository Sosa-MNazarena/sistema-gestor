package strategy

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)


type ExcelReader struct {
	Path string
}


func (e *ExcelReader) ReadData() ([]map[string]interface{}, error) {
	f, err := excelize.OpenFile(e.Path)
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo Excel: %v", err)
	}

	rows, err := f.GetRows("Hoja1")
	if err != nil {
		return nil, fmt.Errorf("no se pueden leer las filas del archivo Excel: %v", err)
	}
	var data []map[string]interface{}
	for i, row := range rows {
		if i == 0 {
			continue 
		}
		if len(row) < 2 {
			continue
		}
		entry := map[string]interface{}{
			"nombre": row[0],
			"descripcion": row[1],
			"categoria": row[2],
			"proveedor": row[3],
			"precio": row[4],

		}
		data = append(data, entry)
	}
	
	return data, nil
}
