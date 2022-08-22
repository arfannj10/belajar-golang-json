package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Arfan",
		MiddleName: "Ganteng",
		LastName:   "Baik",
		Hobbies:    []string{"Code", "Design", "Read"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Arfan","MiddleName":"Ganteng","LastName":"Baik","Age":0,"Married":false,"Hobbies":["Code","Design","Read"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Arfan",
		Addresses: []Address{
			{
				Street: "Jalan Kironggo",
				Country: "Indonesia",
				PostalCode: "2929",
			},
			{
				Street: "Jalan Raung",
				Country: "Indonesia",
				PostalCode: "1229",
			},
		},
	} 

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Arfan","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Kironggo","Country":"Indonesia","PostalCode":"2929"},{"Street":"Jalan Raung","Country":"Indonesia","PostalCode":"1229"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArray(t *testing.T) {
	jsonString := `
	[{"Street":"Jalan Kironggo","Country":"Indonesia","PostalCode":"2929"},{"Street":"Jalan Raung","Country":"Indonesia","PostalCode":"1229"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	Addresses:= []Address{
		{
			Street: "Jalan Kironggo",
			Country: "Indonesia",
			PostalCode: "2929",
		},
		{
			Street: "Jalan Raung",
			Country: "Indonesia",
			PostalCode: "1229",
		},
	}
	

	bytes, _  := json.Marshal(Addresses)
	fmt.Println(string(bytes))
}

