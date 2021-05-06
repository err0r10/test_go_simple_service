package main

import (
	"encoding/json"
	"net/http"
)

type username struct {
	Username string
}

type greeting struct {
	Greeting string
}

func you(w http.ResponseWriter, r *http.Request) {

	var u username
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method == "POST" {
		g := greeting{u.Username}
		js, err := json.Marshal(g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func main() {
	http.HandleFunc("/you", you)
	http.ListenAndServe(":5001", nil)
}
