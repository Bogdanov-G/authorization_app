package app

import (
	"encoding/json"
	"github.com/Bogdanov-G/authorization_app/dto"
	"github.com/Bogdanov-G/authorization_app/errs"
	"github.com/Bogdanov-G/authorization_app/service"
	"net/http"
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
	urlParams := map[string]string{}

	for k := range r.URL.Query() {
		urlParams[k] = r.URL.Query().Get(k)
	}
	if urlParams["token"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		writeResponse(rw, &Response{false}, r)
		return
	}

	authorized, appErr := th.service.Verify(urlParams)
	if appErr != nil {
		rw.WriteHeader(appErr.Code)
		writeResponse(rw, &Response{false}, r)
		return
	}
	if authorized {
		rw.WriteHeader(http.StatusOK)
	} else {
		rw.WriteHeader(http.StatusForbidden)
	}
	writeResponse(rw, &Response{authorized}, r)
}

type Response struct {
	IsAuthorized bool `xml:"isAuthorized" json:"isAuthorized,omitempty"`
}
