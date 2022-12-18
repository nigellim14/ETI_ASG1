package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type AllDrivers struct {
	Drivers map[string]Driver `json:"Drivers"`
}

type AllPassengers struct {
	Passengers map[string]Passenger `json:"Passengers"`
}

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

// Starting main console
func main() {
outer:
	for {
		fmt.Println("===========================================")
		fmt.Println("Welcome to the Ride-Sharing Platform\n",
			"1. Proceed to create a Passenger Account\n",
			"2. Proceed to create a Driver Account\n",
			"3. Quit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1: //Passenger console
			passengermain()
		case 2: //Driver console
			drivermain()
		case 3: //quit
			break outer
		}
	}
}

// Driver main page
func drivermain() {
outer:
	for {
		fmt.Println("===========================================")
		fmt.Println("You are in the Driver Management Console\n",
			"1. List all drivers\n",
			"2. Create new drivers\n",
			"3. Update driver\n",
			"4. Back to main page\n",
			"5. Quit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1: //List all drivers
			driverlist()
		case 2: //Creating account
			drivercreate()
		case 3: //Updating account
			driverupdate()
		case 4: //Return to main page
			main()
		case 5: //Quit
			break outer
		}
	}
}

// Driver list display
func driverlist() {

	//Conneting to MYSQL Database 'my_db'
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")

	//Handle error
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Retrieving driver's information from database
	results, err := db.Query("SELECT FirstName, LastName, MobileNum, EmailAdd, IdNum, CarLicen FROM Driver")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var driver Driver

		err = results.Scan(&driver.FirstName, &driver.LastName, &driver.MobileNum, &driver.EmailAdd, &driver.IdNum, &driver.CarLicen)
		if err != nil {
			panic(err.Error())
		}

		//Display database retrieved
		fmt.Println("====================================")
		fmt.Println(" First Name:"+driver.FirstName+"\n",
			"Last Name:"+driver.LastName+"\n",
			"Mobile Number:", driver.MobileNum, "\n",
			"Email Address:"+driver.EmailAdd+"\n",
			"ID Number:", driver.IdNum, "\n",
			"Car License Number:", driver.CarLicen)
	}
}

// Driver creating an account
func drivercreate() {
	var newDriver Driver
	var FirstName string

	fmt.Print("\nPlease enter your first name: ")
	fmt.Scanf("%v\n", &FirstName)

	fmt.Print("Please enter your last name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newDriver.LastName = strings.TrimSpace(input)

	fmt.Print("Please enter your Mobile num: ")
	fmt.Scanf("%d\n", &(newDriver.MobileNum))

	fmt.Print("Please enter your email address: ")
	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	newDriver.EmailAdd = strings.TrimSpace(input1)

	fmt.Print("Please enter your identification number(E.g 1234567): ")
	fmt.Scanf("%d\n", &(newDriver.IdNum))

	fmt.Print("Please enter your car license number(E.g 1234): ")
	fmt.Scanf("%d\n", &(newDriver.CarLicen))

	//Adding user create account into MYSQL Database
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")

	//Handle error
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	}

	//Using SQL func'INSERT' to add account into database
	insert, err := db.Query("INSERT INTO `my_db`.`Driver` (`FirstName`, `LastName`, `MobileNum`, `EmailAdd`,  `IdNum`, `CarLicen`) VALUES (?, ?, ?, ?, ?, ?)", newDriver.FirstName, newDriver.LastName, newDriver.MobileNum, newDriver.EmailAdd, newDriver.IdNum, newDriver.CarLicen)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("You have successfully added into the driver database")
}

// Driver updating information
func driverupdate() {
	var newDriver Driver
	var FirstName string

	fmt.Print("Please enter your original first name: ")
	fmt.Scanf("%v\n", &FirstName)

	fmt.Print("Please enter your updated last name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newDriver.LastName = strings.TrimSpace(input)

	fmt.Print("Please enter your updated Mobile num: ")
	fmt.Scanf("%d\n", &(newDriver.MobileNum))

	fmt.Print("Please enter your updated email address: ")
	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	newDriver.EmailAdd = strings.TrimSpace(input1)

	fmt.Print("Please enter your car license number: ")
	fmt.Scanf("%d\n", &(newDriver.CarLicen))

	jsonString, _ := json.Marshal(newDriver)
	resbody := bytes.NewBuffer(jsonString)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/drivers/"+FirstName, resbody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Driver", FirstName, "is being updated")
			} else if res.StatusCode == 409 {
				fmt.Println("Error - driver", FirstName, "is not being updated")
			}
		}
	}
}

// Passenger main page
func passengermain() {
outer:
	for {
		fmt.Println("===========================================")
		fmt.Println("Passenger Management Console\n",
			"1. List all passengers\n",
			"2. Create new passengers\n",
			"3. Update passenger\n",
			"4. Back to main page\n",
			"5. Quit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1: //List all passenger
			passengerlist()
		case 2: //Creating account
			passengercreate()
		case 3: //updating account
			passengerupdate()
		case 4: //Return to main page
			main()
		case 5: //quit
			break outer
		}
	}
}

// Passenger list display
func passengerlist() {

	//Connecting to localhost to retrieved records
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/passengers", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				var results AllPassengers
				json.Unmarshal(body, &results)
				fmt.Println("===========================================")
				fmt.Println(" ")
				fmt.Println("List of passengers")
				fmt.Println("-------------------------------------------")

				for k, v := range results.Passengers {
					fmt.Println("\nFirst Name:", v.PasFirstName, "", k, "")
					fmt.Println("Last Name:", v.PasLastName)
					fmt.Println("Mobile Num:", v.PasMobileNum)
					fmt.Println("Email Address:", v.PasEmailAdd)
					fmt.Println()
				}
			}
		}

	}
}

// Passenger creating an account
func passengercreate() {
	var newPassenger Passenger
	var PasFirstName string

	fmt.Print("\nPlease enter your first name: ")
	fmt.Scanf("%v\n", &PasFirstName)

	fmt.Print("Please enter your last name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newPassenger.PasLastName = strings.TrimSpace(input)

	fmt.Print("Please enter your Mobile num: ")
	fmt.Scanf("%d\n", &(newPassenger.PasMobileNum))

	fmt.Print("Please enter your email address: ")
	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	newPassenger.PasEmailAdd = strings.TrimSpace(input1)

	jsonString, _ := json.Marshal(newPassenger)
	resbody := bytes.NewBuffer(jsonString)

	//Adding into the localhost of created account
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/passengers/"+PasFirstName, resbody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Passenger", PasFirstName, "is being created into the passenger list")
			} else if res.StatusCode == 409 {
				fmt.Println("Error - passenger", PasFirstName, "exists in the system, kindly please try with smallercase or uppercase letters")
			}

		}
	}
}

// Passenger updating account
func passengerupdate() {
	var newPassenger Passenger
	var PasFirstName string

	fmt.Print("Please enter your original first name: ")
	fmt.Scanf("%v\n", &PasFirstName)

	fmt.Print("Please enter your updated last name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newPassenger.PasLastName = strings.TrimSpace(input)

	fmt.Print("Please enter your updated Mobile num: ")
	fmt.Scanf("%d\n", &(newPassenger.PasMobileNum))

	fmt.Print("Please enter your updated email address: ")
	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	newPassenger.PasEmailAdd = strings.TrimSpace(input1)

	jsonString, _ := json.Marshal(newPassenger)
	resbody := bytes.NewBuffer(jsonString)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/passengers/"+PasFirstName, resbody); err == nil {
		if res, err := client.Do(req); err == nil {

			if res.StatusCode == 202 {
				fmt.Println("Passenger", PasFirstName, "is being updated")
			} else if res.StatusCode == 409 {
				fmt.Println("Error - passenger", PasFirstName, "is not being updated")
			}
		}
	}
}
