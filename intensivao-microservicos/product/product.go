package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)
// Estrutura de produtos
type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Products struct {
	Products []Product
}
// Ler o conteudo da pasta PRODUCTS.JSON, e trata o erro qdo for diferente NUll.
func loadData() []byte {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println("erro: ",err.Error())
	}
	defer jsonFile.Close() // Fechar o arquivo
	
// Guarda o valor na variavel Data:
	data, err := ioutil.ReadAll(jsonFile)
	return data
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadData()
	w.Write([]byte(products))
}
//para buscar uma unico produto.
func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := loadData()

	var products Products
	json.Unmarshal(data, &products)

	for _, v := range products.Products {
		if v.Uuid == vars["id"] {
			product, _ := json.Marshal(v)
			w.Write([]byte(product))
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", ListProducts)
	r.HandleFunc("/product/{id}", GetProductById)
	http.ListenAndServe(":8081", r)
}
