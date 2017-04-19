package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

func main() {
	normalCards, crisisCards, err := readFile("Existential crisis - cardz.csv")
	if err != nil {
		panic(err)
	}
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	normalData := make(map[string]interface{})
	crisisData := make(map[string]interface{})
	normalData["cards"] = normalCards
	crisisData["cards"] = crisisCards
	v := bytes.Buffer{}
	temp.Execute(&v, normalData)
	ioutil.WriteFile("normal_cards.html", v.Bytes(), 0644)
	v = bytes.Buffer{}
	temp.Execute(&v, crisisData)
	ioutil.WriteFile("crisis_cards.html", v.Bytes(), 0644)
}
