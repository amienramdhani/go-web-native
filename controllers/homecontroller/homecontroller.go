package homecontroller

import (
	"html/template"
	"net/http"
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	temp := template.Must(template.ParseFiles("views/home/index.html"))

	temp.Execute(writer, nil)
}