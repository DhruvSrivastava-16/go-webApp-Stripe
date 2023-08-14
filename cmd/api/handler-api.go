package main

import (
	"encoding/json"
	"net/http"
)

type stripePayload struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	ID      int    `json:"id"`
}

func (app *application) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	test_response := jsonResponse{
		OK:      true,
		Message: "Successully Tested!",
	}

	out, err := json.MarshalIndent(test_response, "", " 	")
	if err != nil {
		app.errorLog.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
