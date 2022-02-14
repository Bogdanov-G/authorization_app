package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// writeResponse writes HTTP responses in JSON or XML formats depending on
// give in the request.
func writeResponse(rw http.ResponseWriter, rwInterface interface{},
	r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		if err := xml.NewEncoder(rw).Encode(rwInterface); err != nil {
			panic(err)
		}
	} else {
		rw.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(rw).Encode(rwInterface); err != nil {
			panic(err)
		}
	}
}
