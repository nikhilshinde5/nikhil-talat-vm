package main

import (
	"encoding/json"
	"io/ioutil"
)

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func main() {
	data := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)

	dataMap := map[int]string{

		1: "nikhil",
		2: "vaishnavi",
		3: "nilesh",
		4: "ramkrushnahari",
		5: "shruti",
	}
	file_2, _ := json.MarshalIndent(dataMap, "", " ")
	_ = ioutil.WriteFile("testMap.json", file_2, 0644)
}
