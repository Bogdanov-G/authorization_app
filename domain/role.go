package domain

import "strings"

type RolePermissions struct {
	permissions map[string][]string
}

func (rp RolePermissions) IsAuthorized(role string, routeName string) bool {
	for _, r := range rp.permissions[role] {
		if r == strings.TrimSpace(routeName) {
			return true
		}
	}
	return false
}

func NewRolePermissions() RolePermissions {
	return RolePermissions{map[string][]string{
		"customer": []string{"GetCustomer", "CreateTransaction"},
		"employee": []string{"GetAllCustomers", "GetCustomer", "CreateAccount", "CreateTransaction"},
	}}
}
