package main

import (
	"go-web-native/config"
	categorycontroller "go-web-native/controllers/categoryController"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// memanggil homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//Memanggil category
	http.HandleFunc("/category", categorycontroller.Index)
	http.HandleFunc("/category/add", categorycontroller.Add)
	http.HandleFunc("/category/edit", categorycontroller.Edit)
	http.HandleFunc("/category/delete", categorycontroller.Delete)

	//Memanggil Product
	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/add", productcontroller.Add)
	http.HandleFunc("/product/detail", productcontroller.Detail)
	http.HandleFunc("/product/edit", productcontroller.Edit)
	http.HandleFunc("/product/delete", productcontroller.Delete)

	//log serve
	log.Println("Server run in port 8080")
	http.ListenAndServe(":8080",nil)
}