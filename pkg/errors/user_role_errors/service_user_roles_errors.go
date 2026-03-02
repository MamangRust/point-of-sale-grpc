package userrole_errors

import (
	"net/http"
	"pointofsale/pkg/errors"
)

var (
	ErrFailedAssignRoleToUser = errors.NewErrorResponse("Failed to assign role to user", http.StatusInternalServerError)
	ErrFailedRemoveRole       = errors.NewErrorResponse("Failed to remove role from user", http.StatusInternalServerError)
)
