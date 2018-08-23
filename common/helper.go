package common

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/form"
)

func ParseForm(r *http.Request, i interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	decoder := form.NewDecoder()
	err = decoder.Decode(&i, r.Form)
	return err
}

func JsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ParseJson(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if len(body) < 1 {
		return nil, errors.New("No body")
	}
	return body, nil
}
