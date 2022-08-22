package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("Customer_out.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Arfan",
		MiddleName: "Ganteng",
		LastName: "Baik",
	}

	_ = encoder.Encode(customer)

	fmt.Println(customer)
}