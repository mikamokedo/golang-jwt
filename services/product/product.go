package productServices

import (
	productModels "jwt-api/models/product"
	"jwt-api/utils"
	"net/http"

	"github.com/gorilla/mux"
)


type Handler struct {
	store productModels.ProductStore
}

func NewHandler(s productModels.ProductStore) *Handler {
	return &Handler{store: s}
}


func (h *Handler) RegisRoutes(router *mux.Router) {
	router.HandleFunc("/product", h.getAllProducts).Methods("GET")
	router.HandleFunc("/product", h.createProduct).Methods("POST")
}



func (h *Handler) createProduct(w http.ResponseWriter, r *http.Request) {
	var payload productModels.CreateProductPayload
	if err := utils.ParseJsonBody(r, &payload); err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.CreateProduct(&payload); err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJsonResponse(w, http.StatusCreated, nil)
}



func (h *Handler) getAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJsonResponse(w, http.StatusOK, products)

} 