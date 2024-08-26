package app

import (
	"net/http"
)

type application struct {
	server *http.Server
}
