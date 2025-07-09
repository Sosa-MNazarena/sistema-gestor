package strategy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiReader struct {
	BaseURL string
}

func (a *ApiReader) ReadData() ([]map[string]interface{}, error) {
	fmt.Println("Leyendo datos de la API:", a.BaseURL)

	resp, err := http.Get(a.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("error al hacer la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("la API devolvió estado %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
	}

	var data []map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("error al parsear JSON: %v", err)
	}

	return data, nil
}