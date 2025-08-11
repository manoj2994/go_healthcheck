package handlers

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Status string `json:"status"`         // `json:"status"` is a struct tag to specify the JSON field name
	Code   int    `json:"code,omitempty"` // `omitempty` omits the field if its zero value
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {

	var s Status = Status{Status: "ok"}
	res, err := json.Marshal(s)
	if err != nil {
		res, _ = json.Marshal(map[string]string{"status": "Not Ok"})
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))

}
