package run

import (
	"net/http"
	"sync"
)

type Run struct {
	_lock    *sync.Mutex
	services map[string]string
	server   *http.Server
}
