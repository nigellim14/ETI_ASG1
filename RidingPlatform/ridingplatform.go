package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Driver struct {
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	MobileNum int    `json:"Mobile Num"`
	EmailAdd  string `json:"Email Address"`
	IdNum     int    `json:"Identification Number"`
	CarLicen  int    `json:"Car License Number"`
}

type Passenger struct {
	PasFirstName string `json:"Passenger First Name"`
	PasLastName  string `json:"Passenger Last Name"`
	PasMobileNum int    `json:"Passenger Mobile Num"`
	PasEmailAdd  string `json:"Passenger Email Address"`
}

type AllDrivers struct {
	Drivers map[string]Driver `json:"Drivers"`
}

type AllPassengers struct {
	Passengers map[string]Passenger `json:"Passengers"`
}

var drivers map[string]Driver = map[string]Driver{
	" ": Driver{"Peter", "Lim", 97865422, "peter@gmail.com", 1358965, 1523},
}

var passengers map[string]Passenger = map[string]Passenger{
	" ": Passenger{"Casey", "Tan", 97865422, "casey@gmail.com"},
}

type Test struct {
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	Value     int    `json:"Value"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/drivers", driversFilter)
	router.HandleFunc("/api/v1/drivers/{driver_id}", alldrivers).Methods("GET", "POST", "PUT")

	router.HandleFunc("/api/v1/passengers", passengersFilter)
	router.HandleFunc("/api/v1/passengers/{passenger_id}", allpassengers).Methods("GET", "POST", "PUT")

	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func driversFilter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	results := map[string]Driver{}
	if value := query.Get("q"); len(value) > 0 {
		for k, v := range drivers {
			if strings.Contains(strings.ToLower(v.FirstName), strings.ToLower(value)) {
				results[k] = v
			}
		}

		if len(results) == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No driver found")
		} else {
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Driver `json:"Search Results"`
			}{results})
		}
	} else if value := query.Get("value"); len(value) > 0 {
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body), "\n", err)

		for k, v := range drivers {
			value, _ := strconv.Atoi(value)
			if v.CarLicen >= value {
				results[k] = v
			}
		}

		if len(results) == 0 {
			fmt.Fprintf(w, "No driver eligible")
		} else {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Driver `json:"Eligible Driver(s)"`
			}{results})
		}
	} else {
		alldrivers := AllDrivers{drivers}

		json.NewEncoder(w).Encode(alldrivers)
	}
}

func passengersFilter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	results := map[string]Passenger{}
	if value := query.Get("q"); len(value) > 0 {
		for l, x := range passengers {
			if strings.Contains(strings.ToLower(x.PasFirstName), strings.ToLower(value)) {
				results[l] = x
			}
		}

		if len(results) == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No passenger found")
		} else {
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Passenger `json:"Search Results"`
			}{results})
		}
	} else if value := query.Get("value"); len(value) > 0 {
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body), "\n", err)

		for l, x := range passengers {
			value, _ := strconv.Atoi(value)
			if x.PasMobileNum >= value {
				results[l] = x
			}
		}

		if len(results) == 0 {
			fmt.Fprintf(w, "No passenger eligible")
		} else {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Passenger `json:"Eligible Passenger(s)"`
			}{results})
		}
	} else {
		allpassengers := AllPassengers{passengers}

		json.NewEncoder(w).Encode(allpassengers)
	}
}

func allpassengers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Println(params["passenger_id"], r.Method)

	if v, ok := passengers[params["passenger_id"]]; ok {
		if r.Method == "GET" {
			json.NewEncoder(w).Encode(v)
		} else if r.Method == "POST" {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "Passenger ID exists")
		} else if r.Method == "PUT" {
			if body, err := ioutil.ReadAll(r.Body); err == nil {
				var data Passenger

				if err := json.Unmarshal(body, &data); err == nil {
					fmt.Printf("PUT ### %v", data)
					w.WriteHeader(http.StatusAccepted)
					passengers[params["passenger_id"]] = data
				}
			}
		} else {
			delete(drivers, params["passenger_id"])
			fmt.Fprintf(w, params["passenger_id"]+" Deleted")
		}
	} else if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Passenger

			if err := json.Unmarshal(body, &data); err == nil {
				w.WriteHeader(http.StatusAccepted)
				passengers[params["passenger_id"]] = data
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		if r.Method == "PUT" {
			fmt.Fprintf(w, "Passenger ID does not exist")
		} else {
			fmt.Fprintf(w, "Invalid Passenger ID")
		}
	}
}

func alldrivers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Println(params["driver_id"], r.Method)

	if v, ok := drivers[params["driver_id"]]; ok {
		if r.Method == "GET" {
			json.NewEncoder(w).Encode(v)
		} else if r.Method == "POST" {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "Driver ID exists")
		} else if r.Method == "PUT" {
			if body, err := ioutil.ReadAll(r.Body); err == nil {
				var data Driver

				if err := json.Unmarshal(body, &data); err == nil {
					fmt.Printf("PUT ### %v", data)
					w.WriteHeader(http.StatusAccepted)
					drivers[params["driver_id"]] = data
				}
			}
		} else {
			delete(drivers, params["driver_id"])
			fmt.Fprintf(w, params["driver_id"]+" Deleted")
		}
	} else if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Driver

			if err := json.Unmarshal(body, &data); err == nil {
				w.WriteHeader(http.StatusAccepted)
				drivers[params["driver_id"]] = data
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		if r.Method == "PUT" {
			fmt.Fprintf(w, "Driver ID does not exist")
		} else {
			fmt.Fprintf(w, "Invalid Driver ID")
		}
	}
}
