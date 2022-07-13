package data

type JSONDocument struct {
}

func (doc *JSONDocument) ConvertToXML() string {
	return "<xml></xml>"
}

type JSONDocumentAdapter struct {
	jsonDocument JSONDocument
}

func (adapter *JSONDocumentAdapter) SendXMLData() {
	adapter.jsonDocument.ConvertToXML()
	println("отправка XML данных")
}
