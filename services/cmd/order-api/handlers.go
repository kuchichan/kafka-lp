package main

import (
	"encoding/json"
	"kuchichan/kafka-lp/internal/common"
	"kuchichan/kafka-lp/internal/models"
	"net/http"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "OK",
		"environment": app.env,
		"version":     app.version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Json could not be processed", http.StatusInternalServerError)
	}

	js = append(js, '\n')
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func (app *application) createOrder(w http.ResponseWriter, r *http.Request) {
	input := models.Order{}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "bad json")
		return 
	}
	orderEvent := common.OrderToReceivedEvent(input)

	app.kafkaPublisher.PublishMessage("order-received", orderEvent)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
