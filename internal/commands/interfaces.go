package commands

import (
	"io"
	"net/http"
)

type fetchClient interface {
	Get(url string, headers map[string]string) (resp *http.Response, err error)
	Post(url string, body io.Reader, headers map[string]string) (resp *http.Response, err error)
	Put(url string, body io.Reader, headers map[string]string) (resp *http.Response, err error)
	Patch(url string, body io.Reader, headers map[string]string) (resp *http.Response, err error)
	Delete(url string, body io.Reader, headers map[string]string) (resp *http.Response, err error)
}

type storage interface {
	Create() error
	Save(data interface{}) error
	Load(dest interface{}) error
	Exists() bool
}
