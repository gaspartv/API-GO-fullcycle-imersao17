package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/gaspartv/API-GO-fullcycle-imersao17/internal/entity"
	"github.com/gaspartv/API-GO-fullcycle-imersao17/internal/service"
	"github.com/go-chi/chi/v5"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(productService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: productService}
}

func (wph *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wph.ProductService.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	error := json.NewEncoder(w).Encode(products)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
}

func (wph *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	product, err := wph.ProductService.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	error := json.NewEncoder(w).Encode(product)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
}

func (wph *WebProductHandler) GetProductsByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryID")
	if categoryID == "" {
		http.Error(w, "categoryID is required", http.StatusBadRequest)
		return
	}
	products, err := wph.ProductService.GetProductByCategoryID(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	error := json.NewEncoder(w).Encode(products)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
}

func (wph *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := wph.ProductService.CreateProduct(product.Name, product.Description, product.CategoryID, product.ImageURL, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	error := json.NewEncoder(w).Encode(result)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
}
