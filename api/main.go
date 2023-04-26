package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"math"
	"net/http"
	"net/http/httputil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Temperature struct {
	Set         float64 `json:"set"`
	Real        float64 `json:"real"`
	CronEnabled bool    `json:"cron"`
}

var temperature Temperature

func main() {
	//Setup Port and initalize temperature
	port := ":8080"
	temperature.Set = readDatabaseTemp()
	temperature.Real = getTemperatureHW()
	temperature.CronEnabled = readDatabaseCron()

	fmt.Println("Server Running on", port)
	fmt.Println("Inital Temperature is:", temperature.Set)
	fmt.Println("Is Cron enabled?", temperature.CronEnabled)
	setupCron()
	// Define the API endpoint
	http.HandleFunc("/api/realTemperature", realTemperatureHandler)
	http.HandleFunc("/api/temperature", temperatureHandler)
	http.HandleFunc("/api/cron", cronHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(port, nil))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func readDatabaseCron() bool {
	dat, err := os.ReadFile("/home/pi/thermostat/api/database-cron")
	check(err)
	sdat := strings.ReplaceAll(string(dat), "\n", "")
	isCronEnabled, error := strconv.ParseBool(sdat)
	if error != nil {
		fmt.Println(error)
		isCronEnabled = false
	}
	return isCronEnabled
}
func readDatabaseTemp() float64 {
	dat, err := os.ReadFile("/home/pi/thermostat/api/database")
	check(err)
	num, error := strconv.ParseFloat(string(dat), 32)
	if error != nil {
		num = 0
	}
	return num
}

func writeDatabaseTemp(temp float64) {
	err := os.WriteFile("/home/pi/thermostat/api/database", []byte(strconv.FormatFloat(temp, 'f', -1, 32)), 0644)
	check(err)
}

func writeDatabaseCron(isEnabled bool) {
	err := os.WriteFile("/home/pi/thermostat/api/database-cron", []byte(strconv.FormatBool(isEnabled)), 0644)
	check(err)
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTemperature(w, r)
	case "POST":
		setTemperaturePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func realTemperatureHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getRealTemperature(w, r)
	case "POST":
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func cronHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getCron(w, r)
	case "POST":
		setCronEnabled(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func setupCron() {
	fmt.Println("Current Time: ", time.Now().String())
	s := gocron.NewScheduler(time.Local)

	var turnOn = func() {
		if temperature.CronEnabled == true {
			fmt.Println("Turning Heater on to 70")
			temperature.Set = 70
			setTemperatureHW()
		}
	}
	var turnOff = func() {
		if temperature.CronEnabled == true {
			fmt.Println("Turning Heater off")
			temperature.Set = 0
			setTemperatureHW()
		}
	}
	_, _ = s.Every(1).Day().At("06:50").Do(turnOn)
	_, _ = s.Every(1).Day().At("22:00").Do(turnOff)
	s.StartAsync()
}

// func imCold() {
// triggered from post to /cold
// temperature.Set = 75
// setTemperatureHW()
// in 5 minutes from now, set temperature back to previous value
// }

func getTemperatureHW() float64 {
	cmd := exec.Command("/usr/bin/python", "/home/pi/thermostat/hw/checkTemp.py", strconv.Itoa(4))

	stdouterr, exiterr := cmd.CombinedOutput()
	if exiterr != nil {
		log.Fatal(exiterr)
	}

	cel, err := strconv.ParseFloat(strings.TrimSuffix(string(stdouterr), "\n"), 32)
	if err != nil {
		panic(err)
	}
	far := (cel * 1.8) + 32
	farRound := math.Round(far*100) / 100

	temperature.Real = farRound
	return farRound
}

func getRealTemperature(w http.ResponseWriter, r *http.Request) {
	// Marshal the temperature struct into JSON
	getTemperatureHW()
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

func getCron(w http.ResponseWriter, r *http.Request) {
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

func setCronEnabled(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Request Dump: ", string(requestDump))

	// Unmarshal the JSON request body into the temperature struct
	decodeErrQuestion := json.NewDecoder(r.Body).Decode(&temperature)

	if decodeErrQuestion != nil {
		http.Error(w, decodeErrQuestion.Error(), http.StatusBadRequest)
		return
	}

	// Write cron from UI to databsase
	writeDatabaseCron(temperature.CronEnabled)

	// Set the content type header and write the success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Cron updated"}`))
}

func setTemperatureHW() {
	var position float32
	switch temperature.Set {
	case 0:
		position = -.8
	case 70:
		position = .3
	case 72:
		position = .45
	case 75:
		position = .6
	case 80:
		position = .75
	}

	cmd := exec.Command("/usr/bin/python", "/home/pi/thermostat/hw/on.py", fmt.Sprintf("%f", position))
	cmd.Env = append(cmd.Environ(), "GPIOZERO_PIN_FACTORY=pigpio")

	stdouterr, exiterr := cmd.CombinedOutput()
	if exiterr != nil {
		log.Fatal(exiterr)
	}
	fmt.Printf("%s\n", stdouterr)

	// Write temperature from UI to databsase
	writeDatabaseTemp(temperature.Set)
}

func setTemperaturePost(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("r.Body: %v\n", r.Body)

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	// Unmarshal the JSON request body into the temperature struct
	decodeErrQuestion := json.NewDecoder(r.Body).Decode(&temperature)
	fmt.Printf("temperature.Set: %v\n", temperature.Set)

	if decodeErrQuestion != nil {
		http.Error(w, decodeErrQuestion.Error(), http.StatusBadRequest)
		return
	}
	setTemperatureHW()
	// Set the content type header and write the success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Temperature updated"}`))
}
