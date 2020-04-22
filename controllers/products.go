package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web_store/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, req *http.Request) {
	allProducts := models.SelectAllProducts()
	temp.ExecuteTemplate(w, "index", allProducts)
}

func New(w http.ResponseWriter, req *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Edit(w http.ResponseWriter, req *http.Request) {
	productId := req.URL.Query().Get("id")
	product := models.EditProduct(productId)
	temp.ExecuteTemplate(w, "edit", product)
}

func Insert(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		name := req.FormValue("name")
		description := req.FormValue("description")
		price := req.FormValue("price")
		quantity := req.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error at price conversion")
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error at price quantity conversion")
		}
		models.CreateNewProduct(name, description, convertedPrice, convertedQuantity)

	}
	http.Redirect(w, req, "/", 301)
}

func Delete(w http.ResponseWriter, req *http.Request) {
	productId := req.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, req, "/", 301)
}

func Update(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		id := req.FormValue("id")
		name := req.FormValue("name")
		description := req.FormValue("description")
		price := req.FormValue("price")
		quantity := req.FormValue("quantity")

		convertedID, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error at id conversion", err)
		}

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error at price conversion", err)
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error at price quantity conversion", err)
		}

		models.UpdateProduct(convertedID, name, description, convertedPrice, convertedQuantity)

	}
	http.Redirect(w, req, "/", 301)
}
