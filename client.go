package npmdl

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseURL string = "https://api.npmjs.org/downloads"

// Client implements a NPM Downloads Counts API client.
// You may use the Point and Range clients directly if preferred,
// however Client exposes them both.
type Client struct {
	// Point *Point
	// Range *Range
}

// New client.
func New() *Client {
	c := &Client{}
	// c.Point = &Point{c}
	// c.Range = &Range{c}
	return c
}

// call rpc style endpoint.
func (c *Client) call(path string) (io.ReadCloser, error) {
	url := baseURL + path

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 400 {
		return res.Body, err
	}

	defer res.Body.Close()

	e := &Error{
		Status:     http.StatusText(res.StatusCode),
		StatusCode: res.StatusCode,
	}

	kind := res.Header.Get("Content-Type")
	if strings.Contains(kind, "text/plain") {
		if b, err := ioutil.ReadAll(res.Body); err == nil {
			e.Message = string(b)
			return nil, e
		}
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(e); err != nil {
		return nil, err
	}

	return nil, e
}
