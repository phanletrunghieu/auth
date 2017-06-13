package auth

import (
	"fmt"
	"net/http"
)

// DefaultLoginHandler default login behaviour
var DefaultLoginHandler = func(req *http.Request, w http.ResponseWriter, session *Session, authorize func(*http.Request, http.ResponseWriter, *Session) (interface{}, error)) {
	currentUser, err := authorize(req, w, session)
	if err == nil {
		if currentUser != nil {
			// write cookie, json
		} else {
		}
	}

	session.Auth.Config.Render.Execute("auth/login", session, req, w)
}

// DefaultRegisterHandler default register behaviour
var DefaultRegisterHandler = func(req *http.Request, w http.ResponseWriter, session *Session, register func(*http.Request, http.ResponseWriter, *Session) (interface{}, error)) {
	user, err := register(req, w, session)
	fmt.Println(user)
	fmt.Println(err)
	if err == nil {
		if user != nil {
			// registered
		}
	}

	session.Auth.Config.Render.Execute("auth/register", session, req, w)
}
