package userServices

import (
	"fmt"
	userModels "jwt-api/models/user"
	authService "jwt-api/services/auth"
	"jwt-api/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)




type Handler struct {
	store userModels.UserStore
}

func NewHandler(store userModels.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.login).Methods("POST")
	router.HandleFunc("/register", h.register).Methods("POST")
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	
	var payload userModels.UserLoginPayload
	if error := utils.ParseJsonBody(r, &payload); error != nil{
		utils.WriteJsonError(w, http.StatusBadRequest, error)
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteJsonError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}
	if err := authService.ComparePassword(user.Password, payload.Password); err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	token, err := authService.CreateJwtToken(user)
	if err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJsonResponse(w, http.StatusOK, map[string]string{"token":token})
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	var payload userModels.UserRegisterPayload
	if error := utils.ParseJsonBody(r, &payload); error != nil{
		utils.WriteJsonError(w, http.StatusBadRequest, error)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteJsonError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	if _, err := h.store.GetUserByEmail(payload.Email); err == nil {
		utils.WriteJsonError(w, http.StatusBadRequest, fmt.Errorf("email %s already exists", payload.Email))
		return
	}

	password, err := authService.HasPassword(payload.Password)
	if err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, err)
		return
	}


	if err := h.store.CreateUser(userModels.UserRegisterPayload{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: password,
	}); err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, err)
		return
	}	

		utils.WriteJsonResponse(w, http.StatusOK, nil)
}