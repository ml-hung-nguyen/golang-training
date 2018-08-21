package helper

import (
	"encoding/json"
	"net/http"
)

func respondwithJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json, _ := json.Marshal(data)
		_, _ = w.Write(json)
	}
}

func tranDataJson(origin interface{}, response interface{}) error {
	data, err := json.Marshal(&origin)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}
	return nil
}
