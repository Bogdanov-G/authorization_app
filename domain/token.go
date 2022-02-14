package domain

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	CustomerId  string   `xml:"customer-id" json:"customer_id,omitempty"`
	AccountsIds []string `xml:"accounts-ids" json:"accounts_ids,omitempty"`
	Username    string   `xml:"username" json:"username,omitempty"`
	Expiry      int64    `xml:"expiry" json:"expiry,omitempty"`
	Role        string   `xml:"role" json:"role,omitempty"`
}

func (c Claims) IsRequestVerified(urlParams map[string]string) bool {
	if c.Role == "employee" {
		return true
	}
	if c.Username != urlParams["customer_id"] {
		return false
	}
	if urlParams["account_id"] == "" {
		return true
	}
	for _, account := range c.AccountsIds {
		if account == urlParams["account_id"] {
			return true
		}
	}
	return false
}

func BuildClaims(mapClaims jwt.MapClaims) (*Claims, error) {
	bytes, err := json.Marshal(mapClaims)
	if err != nil {
		return nil, err
	}
	claims := Claims{}
	err = json.Unmarshal(bytes, &claims)
	if err != nil {
		return nil, err
	}
	return &claims, nil
}
