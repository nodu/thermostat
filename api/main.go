package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type Temperature struct {
	Value float64 `json:"value"`
}

var temperature Temperature

func main() {

	// Define the API endpoint
	http.HandleFunc("/temperature", temperatureHandler)

	fmt.Println("Server Running on:")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTemperature(w, r)
	case "POST":
		setTemperature(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTemperature(w http.ResponseWriter, r *http.Request) {
	// Marshal the temperature struct into JSON
	tempJSON, err := json.Marshal(temperature)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type header and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tempJSON)
}

func setTemperature(w http.ResponseWriter, r *http.Request) {
	// Unmarshal the JSON request body into the temperature struct
	decodeErrQuestion := json.NewDecoder(r.Body).Decode(&temperature)
	fmt.Printf("%+v\n", temperature.Value)

	cmd := exec.Command("/usr/bin/python", "/home/pi/thermostat/hw/on.py", fmt.Sprintf("%f", temperature.Value))
	cmd.Env = append(cmd.Environ(), "GPIOZERO_PIN_FACTORY=pigpio")

	stdouterr, exiterr := cmd.CombinedOutput()
	if exiterr != nil {
		log.Fatal(exiterr)
	}
	fmt.Printf("%s\n", stdouterr)

	if decodeErrQuestion != nil {
		http.Error(w, decodeErrQuestion.Error(), http.StatusBadRequest)
		return
	}

	// Set the content type header and write the success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Temperature updated"}`))
}
