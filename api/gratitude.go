package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wty/bchain/connutils"
	"github.com/wty/models"
)

const (
	minerAddr = "localhost:3002"
)

// Quote -
func (api *API) Quote(w http.ResponseWriter, req *http.Request) {
	quote := api.quotes.RandomQuote()
	w.Write([]byte(quote.Text))
}

// SecretQuote -
func (api *API) SecretQuote(w http.ResponseWriter, req *http.Request) {
	quote := api.quotes.RandomQuote()
	w.Write([]byte(quote.Text))
}

// SendGratitude sends gratitude
func (api *API) SendGratitude(w http.ResponseWriter, r *http.Request) {
	//userName := mux.Vars(r)["username"]
	user := api.GetUserFromContext(r)

	decoder := json.NewDecoder(r.Body)
	gratitude := models.NewGratitude(user.Username)
	err := decoder.Decode(gratitude)

	if err != nil ||
		gratitude.Recipient == "" || gratitude.Content == "" {
		http.Error(w, "Missing recipient or message", http.StatusBadRequest)
		return
	}

	log.Println("gratitude received", gratitude)
	go sendGrat2Centre(gratitude)
	w.Write([]byte("sent gratitude successfully"))
}

func sendGrat2Centre(grat *models.Gratitude) {
	err := connutils.SendCmdData(minerAddr, "gratitude", grat)
	if err != nil {
		log.Println("gratitude miner send error ", err)
		return
	}
	log.Println("gratitude successfully sent for mining")
}
