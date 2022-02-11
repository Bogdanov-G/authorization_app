package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/Bogdanov-G/authorization_app/dto"
	"github.com/Bogdanov-G/authorization_app/errs"
	"github.com/Bogdanov-G/authorization_app/service"
)

type AuthHandler struct {
	service service.AuthService
}

func (th AuthHandler) Login(rw http.ResponseWriter, r *http.Request) {
	var logReq dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		appErr := errs.NewBadRequestError("bad request error")
		rw.WriteHeader(appErr.Code)
		writeResponse(rw, appErr.AsMessage(), r.Header.Get("Content-Type"))
		return
	}
	token, appErr := th.service.Login(logReq)
	if appErr != nil {
		rw.WriteHeader(appErr.Code)
		writeResponse(rw, appErr.AsMessage(), r.Header.Get("Content-Type"))
	} else {
		writeResponse(rw, token, r.Header.Get("Content-Type"))
	}
}

func (th AuthHandler) Register(rw http.ResponseWriter, r *http.Request) {
	// TBD
}

func (th AuthHandler) Verify(rw http.ResponseWriter, r *http.Request) {
	// TBD
}

// writeResponse() writes HTTP responses in JSON or XML formats depending on
// give contentType value.
func writeResponse(rw http.ResponseWriter, rwInterface interface{},
	contentType string) {
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
