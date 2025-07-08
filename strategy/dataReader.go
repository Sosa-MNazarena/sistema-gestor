package strategy


type DataReader interface {
	ReadData() ([]map[string]interface{}, error)
}
