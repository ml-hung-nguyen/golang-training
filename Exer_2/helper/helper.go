package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// respondwithJSON write json response format
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func TranDataJson(origin interface{}, response interface{}) error {
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
