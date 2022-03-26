package util

import (
	"encoding/json"
	"net/http"
)

func Series(f ...func() error) error {
	for _, v := range f {
		err := v()
		if err != nil {
			return err
		}
	}
	return nil
}

func Decode(i interface{}, r *http.Request) error {
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(i)
	return err
}

func Encode(i interface{}, w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(i)
}

func HttpErr(err error, w http.ResponseWriter) error {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return err
}
