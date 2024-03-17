package http

import (
	"net/http"

	"github.com/openidea-marketplace/domain"
	"github.com/sirupsen/logrus"
)

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrUserNotFound:
		return http.StatusNotFound
	case domain.ErrUserConflict:
		return http.StatusConflict
	case domain.ErrUserWrongPassword:
		return http.StatusBadRequest
	case domain.ErrUserDuplicateUsername:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
