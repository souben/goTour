package main

import (
	"github.com/souben/json/marshaling"
	"github.com/souben/json/unmarshaling"
)

func main() {
	unmarshaling.UnmarshalWithUnStructuredData()
	unmarshaling.UnmarshalWithStructuredData()
	marshaling.MarshalingWithStructuredData()
	marshaling.MarshalingWithUnStructuredData()
}
