package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"valdson/store/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	allProducts := models.SearchAllProducts()
	templates.ExecuteTemplate(w, "index", allProducts)

}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "new", nil)

}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Price conversion error")
		}

		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Quantity conversion error")
		}

		models.CreateNewProduct(name, description, priceConverted, quantityConverted)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")

	models.DeleteProduct(productID)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {

	productID := r.URL.Query().Get("id")

	product := models.EditProduct(productID)
	templates.ExecuteTemplate(w, "edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		idToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Id conversion error")
		}

		priceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Price conversion error")
		}

		quantityToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Quantity conversion error")
		}

		models.UpdateProduct(idToInt, name, description, priceToFloat, quantityToInt)
	}

	http.Redirect(w, r, "/", 301)
}
