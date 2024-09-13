package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NSurf adds CRF protection to all POST request
func NoSurf(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)
	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return crsfHandler
}

// SessionLoad loads and seaves the sesion on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
