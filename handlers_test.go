package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestHandlerYou(t *testing.T) {
	url := "http://localhost:5001/you"
	fakeName := gofakeit.BeerName()

	buf, err := json.Marshal(username{fakeName})
	if err != nil {
		t.Errorf("Bad json %v", err)
	}

	responsebody := bytes.NewBuffer(buf)
	res, err := http.Post(url, "application/json", responsebody)
	if err != nil {
		t.Errorf("Server not find %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Body is empty %v", err)
	}
	sb := string(body)

	gr, err := json.Marshal(greeting{fakeName})
	if err != nil {
		t.Errorf("Bad json %v", err)
	}

	if sb != string(gr) {
		t.Errorf("response %v not equal greeting = %s", sb, gr)
	}
}
