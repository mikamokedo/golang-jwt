package api

import (
	"database/sql"
	"fmt"
	productModels "jwt-api/models/product"
	userModels "jwt-api/models/user"
	productServices "jwt-api/services/product"
	userServices "jwt-api/services/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db *sql.DB
}

func NewAPIServer(addr string , db *sql.DB) *APIServer {
	return &APIServer{fmt.Sprintf(":%s",addr) ,db}
}

func RunServer(s *APIServer) error {
	r := mux.NewRouter()
    subrouter := r.PathPrefix("/api/v1").Subrouter()

	//user router
	userStore := userModels.NewStore(s.db)
	userHandler := userServices.NewHandler(userStore);
	userHandler.RegisRoutes(subrouter)

	//product router
	productStore := productModels.NewStore(s.db)
	productHandler := productServices.NewHandler(productStore);
	productHandler.RegisRoutes(subrouter)


    log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr,r)
}