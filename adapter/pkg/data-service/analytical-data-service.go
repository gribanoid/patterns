package data_service

type AnalyticalDataService interface {
	SendXMLData()
}

type XMLDocument struct {
}

func (doc XMLDocument) SendXMLData() {
	println("Отправка xml документа")
}
