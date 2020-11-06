package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Status string
}

func CheckDiscount(code string) string {

	if code == "abc" {
		return "-10% de desconto"
	}

	return "-5% de desconto"
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	coupon := r.PostFormValue("coupon")

	description := CheckDiscount(coupon)

	result := Result{Status: description}
	jsonResult, err := json.Marshal(result)

	log.Println(r.FormValue(coupon))
	log.Println(r.FormValue(description))

	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprint(w, string(jsonResult))
}
