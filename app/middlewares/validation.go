package middlewares

import (
	"net/http"

	"github.com/golang/gddo/httputil/header"
)

//ValContentType : controlla che la richiesta sia di tipo application/json
func ValContentType(r *http.Request) error {
	var err error = nil
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			err = http.ErrNotSupported
		}
	}
	return err
}
