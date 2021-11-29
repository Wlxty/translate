package server

import(
	"encoding/json"
	"net/http"
	"fmt"
)
func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	fmt.Fprintf(w, string(js))
}
