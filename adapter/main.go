package main

import "patterns/adapter/pkg/data"

func main() {
	json := new(data.JSONDocumentAdapter)
	json.SendXMLData()
}
