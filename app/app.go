package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mskydream/banking/domain"
	"github.com/mskydream/banking/service"
)

func Start() {
	// mux := http.NewServeMux()

	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
