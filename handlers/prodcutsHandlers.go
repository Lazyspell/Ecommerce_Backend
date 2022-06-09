package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lazyspell/Ecommerce_Backend/helpers"
	"github.com/lazyspell/Ecommerce_Backend/models"
)

func (m *Repository) NewProducts(w http.ResponseWriter, r *http.Request) {
	var payload models.Products

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helpers.BadRequest400(w, "invalid type please check request body")
		return
	}

	var product models.Products

	product.ProductName = payload.ProductName
	product.ImageUrl = payload.ImageUrl
	product.Price = payload.Price
	product.Product = payload.Product
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	_, err = m.DB.NewProductDB(product)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	helpers.Create201(w)

}

func (m *Repository) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := m.DB.AllProducts()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (m *Repository) GetAllProductsAsObjects(w http.ResponseWriter, r *http.Request) {
	var objectList []models.Items
	var ProductObject models.Items
	hats, err := m.DB.GetProductsByCategoryDB("Hats")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	ProductObject.Hats = hats

	jackets, err := m.DB.GetProductsByCategoryDB("Jackets")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	ProductObject.Jackets = jackets

	sneakers, err := m.DB.GetProductsByCategoryDB("Sneakers")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	ProductObject.Sneakers = sneakers

	womens, err := m.DB.GetProductsByCategoryDB("Womens")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	ProductObject.Womens = womens

	mens, err := m.DB.GetProductsByCategoryDB("Mens")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	ProductObject.Mens = mens

	objectList = append(objectList, ProductObject)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objectList)

}

func (m *Repository) GetAllProductsByCategory(w http.ResponseWriter, r *http.Request) {

	payload := r.URL.Query().Get("product")
	products, err := m.DB.GetProductsByCategoryDB(payload)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
