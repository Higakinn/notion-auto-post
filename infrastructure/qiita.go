package infrastructures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type HTTPAPI interface {
	Do(req *http.Request) (*http.Response, error)
}
type QiitaTag struct {
  Name string       `json:"name"`
  Versions []string `json:"versions"`
}
type QiitaClient struct {
	accessToken string
	httpAPI 		HTTPAPI
}

func NewQiitaClient(accessToken string) *QiitaClient {
	return &QiitaClient{
		accessToken: accessToken,
		httpAPI:     http.DefaultClient,
	}
}




