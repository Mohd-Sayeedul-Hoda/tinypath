package encoding

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/request"
)

func EncodeJson[T any](w http.ResponseWriter, r *http.Request, status int, data T) error {

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

func Validated[T request.Validator](r *http.Request) (T, map[string]string, error) {
	var data T

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return data, nil, err
	}

	if problems := data.Valid(r.Context()); len(problems) > 0 {
		return data, problems, fmt.Errorf("invalid %T: %d problems", data, len(problems))
	}

	return data, nil, nil

}
