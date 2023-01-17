package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/palomabarroso/go-web-application/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.ListAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvert, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço: ", price)
		}

		qtdConvert, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão da quantity: ", quantity)
		}
		models.NewProduct(name, description, priceConvert, qtdConvert)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeleteProduct(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	productEdit := models.EditProduct(idProduto)
	temp.ExecuteTemplate(w, "EditProduct", productEdit)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		println("ID do produto", id)
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		priceConvertidoParaFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na convesão do preço para float64:", err)
		}

		quantityConvertidaParaInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na convesão da quantity para int:", err)
		}

		models.UpdateProduct(idConvertidaParaInt, name, description, priceConvertidoParaFloat, quantityConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
