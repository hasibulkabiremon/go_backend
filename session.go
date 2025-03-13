package main

import (
	"errors"
	"net/http"
)

var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request)  error {
	username := r.FormValue("username")
	user, ok := users[username]
	if !ok {
		return AuthError
	}

	//Get session token from cookie
	sessionToken, err := r.Cookie("session_token")
	if err != nil || sessionToken.Value != user.SessionToken {
		return AuthError
	}

	//Get csrf token from cookie
	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRFToken {
		return AuthError
	}

	return nil
}


