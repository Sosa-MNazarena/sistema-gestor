package strategy

import "fmt"

type ApiReader struct {
	BaseURL string
}

func (a *ApiReader) ReadData() ([]map[string]interface{}, error) {
	fmt.Println("Leyendo datos de la API:", a.BaseURL)
	return []map[string]interface{}{
		{"nombre": "Producto API 1", "precio": 120.0},
	}, nil
}