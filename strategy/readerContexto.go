package strategy

type ReaderContext struct {
	Reader DataReader
}

func (rc* ReaderContext) SetReader(reader DataReader) {
	rc.Reader = reader
}

func (rc* ReaderContext) ProcessData() ([]map[string]interface{}, error) {
	if rc.Reader == nil {
		return nil, nil
	}
	return rc.Reader.ReadData()
}