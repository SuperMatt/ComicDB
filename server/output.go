package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func jsonHeader(w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
}

//SendJSONData ...
func SendJSONData(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		SendInternalServerError(w, err)
	}

	jsonHeader(w)
	fmt.Fprintf(w, string(b))
}

//SendInternalServerError ...
func SendInternalServerError(w http.ResponseWriter, err error) {
	jsonHeader(w)
	fmt.Fprintf(w, err.Error())

}
