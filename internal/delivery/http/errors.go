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
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrWrongPassword:
		return http.StatusBadRequest
	case domain.ErrDuplicateUsername:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
