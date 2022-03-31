package output

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
  Name string       "json:\"name\""
  Versions []string "json:\"versions\""
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

func (c *QiitaClient) UpdateItem(id, title, body string, tags []QiitaTag) error {
	
	url := fmt.Sprintf("https://qiita.com/api/v2/items/%s", id)
	p, err := json.Marshal(map[string]interface{}{
		"title": title,
		"body":  body,
		"tags":  tags,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(p))
	if err != nil {
		return errors.WithStack(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	resp, err := c.httpAPI.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.WithStack(err)
		}
		return errors.New(string(b))
	}

	return nil
}