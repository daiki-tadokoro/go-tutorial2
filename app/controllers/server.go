package controllers

import (
	"net/http"
	"todo/config"
)

func StartMainServer() error {
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
