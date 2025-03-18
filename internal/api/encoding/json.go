package encoding

import (
	"encoding/json"
	"net/http"
)

func WriteJson[T any](w http.ResponseWriter, status int, headers http.Header, data T) error {

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}

	return nil
}

func DecodeJson[T any](r *http.Request) (T, error) {

	var data T
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return data, err
	}

	return data, nil

}
