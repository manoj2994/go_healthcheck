package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/manoj2994/go_healthcheck/db"
)

type Status struct {
	Status    string `json:"status"`         // `json:"status"` is a struct tag to specify the JSON field name
	Code      int    `json:"code,omitempty"` // `omitempty` omits the field if its zero value
	ID        int64  `json:"id,omitempty"`
	CheckedAt string `json:"checked_at,omitempty"`
}
type Handler struct {
	DB *sql.DB
}

var res []byte

func (h *Handler) Healthcheckwrite(w http.ResponseWriter, r *http.Request) {

	var err error
	_, re := db.InsertData(h.DB)

	var s Status = Status{Status: "ok", ID: re.ID, CheckedAt: re.CheckedAt}
	res, err = json.Marshal(s)
	if err != nil {
		res, _ = json.Marshal(map[string]string{"status": "Not Ok"})
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))

}
func (h *Handler) Healthcheckping(w http.ResponseWriter, r *http.Request) {
	err := h.DB.Ping()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "Failed to ping DB", http.StatusInternalServerError)
		return
	}
	res, _ = json.Marshal(map[string]string{"status": "ok"})

	w.Write([]byte(res))
}
