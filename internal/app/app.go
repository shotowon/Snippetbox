package app

import (
	"net/http"
	"snippetbox/internal/config"
)

type application struct {
	cfg    *config.Config
	server *http.Server
}
