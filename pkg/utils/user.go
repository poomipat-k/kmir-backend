package utils

import (
	"errors"
	"net/http"
	"strconv"
)

func GetUsernameFromRequestHeader(r *http.Request) (string, error) {
	username := r.Header.Get("username")
	if username == "" {
		return "", errors.New("empty username")
	}
	return username, nil
}

func GetUserRoleFromRequestHeader(r *http.Request) (string, error) {
	userRole := r.Header.Get("userRole")
	if userRole == "" {
		return "", errors.New("empty userRole")
	}
	return r.Header.Get("userRole"), nil
}

func GetUserIdFromRequestHeader(r *http.Request) (int, error) {
	strUserId := r.Header.Get("userId")
	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
