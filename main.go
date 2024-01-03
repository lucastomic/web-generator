package main

import (
	"html/template"
	"os"
)

type PageData struct {
	Title string
	Body  string
}

var data = PageData{
	Title: "Mi página de prueba",
	Body:  "Este es el contenido de la página.",
}

func main() {
	t, err := template.ParseFiles("./templates/template.html")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("./generated/pagina.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}
