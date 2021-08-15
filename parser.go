package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/tmccombs/hcl2json/convert"
)

func main() {
	js.Module.Get("exports").Set("parseToString", jsParseToString)
	js.Module.Get("exports").Set("parseToObject", jsParseToObject)
}

// Parse a HCL string into a JSON string
func parseToString(input string) (output string, err error) {
	convertedBytes, err := convert.Bytes([]byte(input), "", convert.Options{})
	if err != nil {
		return "", err
	}
	output = string(convertedBytes)
	return output, nil
}

// Parse a HCL string into a JSON object
func parseToObject(input string) (output *js.Object, err error) {
	jsonString, err := parseToString(input)
	if err != nil {
		empty := js.Object{}
		return &empty, err
	}

	obj := js.Global.Get("JSON").Call("parse", string(jsonString))
	return obj, nil
}

func jsParseToObject(input string) (output *js.Object) {
	obj, _ := parseToObject(input)

	return obj
}

func jsParseToString(input string) (output string) {
	jsonString, _ := parseToString(input)

	return jsonString
}
