package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Cliente struct {
	ID       string    `json: "id"`
	Nome     string    `json: "nome"`
	Telefone string    `json: "telefone"`
	Endereco *Endereco `json: "endereço"`
}

type Endereco struct {
	Rua    string `json: "rua"`
	Cidade string `json: "cidade"`
	Estado string `json: "estado"`
}

var clientes []Cliente

//Aplicação para utilizar com POSTMAN para fazer as requisições

func main() {
	r := mux.NewRouter()

	clientes = append(clientes, Cliente{ID: "1", Nome: "Paulo de Tarso", Telefone: "9.1011-1213", Endereco: &Endereco{Rua: "Aos Romanos", Cidade: "Capitulo Nove", Estado: "Predestinacao"}})
	clientes = append(clientes, Cliente{ID: "2", Nome: "Davi", Telefone: "1.123-456", Endereco: &Endereco{Rua: "Salmos", Cidade: "Capitulo Um", Estado: "Direcao"}})
	r.HandleFunc("/clientes", getClientes).Methods("GET")
	r.HandleFunc("/clientes/{id}", getCliente).Methods("GET")
	r.HandleFunc("/clientes", createCliente).Methods("POST")
	r.HandleFunc("/clientes/{id}", updateCliente).Methods("PUT")
	r.HandleFunc("/clientes/{id}", deleteCliente).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

func createCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cliente Cliente
	_ = json.NewDecoder(r.Body).Decode(&cliente)
	cliente.ID = strconv.Itoa(rand.Intn(10000000000))
	clientes = append(clientes, cliente)
	json.NewEncoder(w).Encode(cliente)
}

func deleteCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range clientes {

		if item.ID == params["id"] {
			clientes = append(clientes[:index], clientes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(clientes)

}

func getCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range clientes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func updateCliente(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop over the movies, range
	for index, item := range clientes {
		// delete de client with id the you i've sent
		//add a new client - the client that we send in the body of postman
		if item.ID == params["id"] {
			clientes = append(clientes[:index], clientes[index+1:]...)
			var cliente Cliente
			_ = json.NewDecoder(r.Body).Decode(&cliente)
			cliente.ID = params["id"]
			clientes = append(clientes, cliente)
			json.NewEncoder(w).Encode(cliente)
			return
		}
	}
}
