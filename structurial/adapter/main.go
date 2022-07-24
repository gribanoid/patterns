package main

import "patterns/structurial/adapter/pkg/data"

func main() {
	json := new(data.JSONDocumentAdapter)
	json.SendXMLData()
}
