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
		writeResponse(rw, appErr.AsMessage(), r)
		return
	}
	token, appErr := th.service.Login(logReq)
	if appErr != nil {
		rw.WriteHeader(appErr.Code)
		writeResponse(rw, appErr.AsMessage(), r)
	} else {
		writeResponse(rw, token, r)
	}
}

func (th AuthHandler) Register(rw http.ResponseWriter, r *http.Request) {
	// TBD
}

func (th AuthHandler) Verify(rw http.ResponseWriter, r *http.Request) {
	// TBD
}
