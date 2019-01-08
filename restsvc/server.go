package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//NOT THREAD SAFE
var employees = *new([]Employee)

func init() {
	employees = append(employees, Employee{"Scott", "Tiger", "Scott.Tiger@oracle.com"})
	employees = append(employees, Employee{"John", "Doe", "John.Doe@oracle.com"})
}

func main() {
	http.HandleFunc("/employees", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(res).Encode(employees); err != nil {
				http.Error(res, "Internal Error", http.StatusInternalServerError)
			}
		} else if req.Method == "POST" {
			if req.Header.Get("Content-Type") != "application/json" {
				http.Error(res, "Bad Request", http.StatusBadRequest)
			}
			var employee Employee
			err := json.NewDecoder(req.Body).Decode(&employee)
			if err != nil {
				http.Error(res, err.Error(), 400)
				return
			}
			employees = append(employees, employee)
			res.WriteHeader(http.StatusCreated)
		} else {
			res.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Employee struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
