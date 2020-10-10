package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	Modules "github.com/LuisAcerv/goeth-api/modules"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

// ClientHandler ethereum client instance
type ClientHandler struct {
	*ethclient.Client
}

func (client ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	module := vars["module"]
	address := r.URL.Query().Get("address")
	hash := r.URL.Query().Get("hash")
	fmt.Println(hash, address)

	switch module {
	case "latest-block":
		_block := Modules.GetLatestBlock(*client.Client)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(_block)
	}

}
