package handlers

import (
	"action-logger/pkg/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shopify/sarama"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var submitLogger models.SubmitLogger
	err := json.NewDecoder(r.Body).Decode(&submitLogger)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Create a new Kafka producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Printf("Failed to create Kafka producer: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer producer.Close()

	// Convert submitLogger to a string representation
	submitLoggerBytes, err := json.Marshal(submitLogger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	submitLoggerStr := string(submitLoggerBytes)

	// Create a Kafka message
	message := &sarama.ProducerMessage{
		Topic: "action-logger",
		Value: sarama.StringEncoder(submitLoggerStr),
	}

	// Publish the message to Kafka
	_, _, err = producer.SendMessage(message)
	if err != nil {
		log.Printf("Failed to publish message to Kafka: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Log the received action logger
	log.Printf("Received action logger: %+v", submitLogger)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Action logger submitted successfully"))
}
