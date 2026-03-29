package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/demo", demoHandler)

	log.Println("Server is starting....")
	err := http.ListenAndServe(":8080", nil) //localhost:8080

	if err != nil {
		log.Fatal("Server started fail, with error: ", err)
	}
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong http method, please change", http.StatusMethodNotAllowed)
		return
	}

	respone := map[string]string{
		"message": "Welcome bitch",
	}

	w.Header().Set("Content-Type", "application/json")
	// data, err := json.Marshal(respone)

	// if err != nil {
	// 	http.Error(w, "Encoding error", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(data)

	json.NewEncoder(w).Encode(respone)
}
