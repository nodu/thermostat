package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/exec"
	"strconv"
)

type Temperature struct {
	Value float64 `json:"value"`
}

var temperature Temperature

func main() {
	//Setup Port and initalize temperature
	port := ":8080"
	temperature.Value = readDatabase()

	fmt.Println("Server Running on", port)
	fmt.Println("Inital Temperature is:", temperature.Value)

	// Define the API endpoint
	http.HandleFunc("/temperature", temperatureHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(port, nil))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func readDatabase() float64 {
	dat, err := os.ReadFile("/home/pi/thermostat/api/database")
	check(err)
	num, error := strconv.ParseFloat(string(dat), 32)
	if error != nil {
		num = 0
	}
	return num
}

func writeDatabase(temp float64) {
	err := os.WriteFile("/home/pi/thermostat/api/database", []byte(strconv.FormatFloat(temp, 'f', -1, 32)), 0644)
	check(err)
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
	var position float32 // TODO= read file or init from memory

	fmt.Printf("r.Body: %v\n", r.Body)

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	// todo Move set temperature at hw level out of the setTemperature http handler
	// Unmarshal the JSON request body into the temperature struct
	decodeErrQuestion := json.NewDecoder(r.Body).Decode(&temperature)
	fmt.Printf("temperature.Value: %v\n", temperature.Value)
	switch temperature.Value {
	case 0:
		position = -.8
	case 70:
		position = .3
	case 72:
		position = .45
	case 75:
		position = .6
	}

	cmd := exec.Command("/usr/bin/python", "/home/pi/thermostat/hw/on.py", fmt.Sprintf("%f", position))
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
	// Write temperature from UI to databsase
	writeDatabase(temperature.Value)

	// Set the content type header and write the success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Temperature updated"}`))
}
