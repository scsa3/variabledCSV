package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	vPath, tPath, rPath := getFlags()
	t := getTemplate(tPath)
	m := getVariables(vPath)

	writeRendered(rPath, t, m)
}

func getFlags() (variablesPath, templatePath, resultPath string) {
	flag.StringVar(&variablesPath, "v", "./variables.json", "default is './variables.json'")
	flag.StringVar(&templatePath, "t", "./cases.csv.gohtml", "default is './cases.csv.gohtml'")
	flag.StringVar(&resultPath, "r", "./cases.csv", "default is './cases.csv'")
	flag.Parse()
	return
}

func getTemplate(path string) *template.Template {
	t, err := template.ParseFiles(path)
	check(err)
	return t
}

func getVariables(path string) (m map[string]interface{}) {
	j, err := ioutil.ReadFile(path)
	check(err)
	err = json.Unmarshal(j, &m)
	check(err)
	return
}

func writeRendered(path string, t *template.Template, m map[string]interface{}) {
	f, err := os.Create(path)
	check(err)
	err = t.Execute(f, m)
	check(err)
	err = f.Close()
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
