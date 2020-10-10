package main

import (
	"fmt"
	"log"
	"net/http"

	Handlers "github.com/LuisAcerv/goeth-api/handler"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func main() {
	client, err := ethclient.Dial("http://localhost:7545")

	if err != nil {
		fmt.Println(err)
	}

	_ = client
	fmt.Println("we have a connection")

	r := mux.NewRouter()

	r.Handle("/api/v1/eth/{module}", Handlers.ClientHandler{client})

	log.Fatal(http.ListenAndServe(":8080", r))

}
