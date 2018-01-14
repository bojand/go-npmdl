package npmdl

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseURL string = "https://api.npmjs.org/downloads"

// GetPointCounts gets the total downloads for a given period, for all packages or a specific package.
// Period is the desired period to be queried. If empty defaults to "last-day".
// Optionally pass in the package name in the pkg param.
func GetPointCounts(period string, pkg string) (out *PointCounts, err error) {
	if period == "" {
		period = "last-day"
	}

	body, err := call("/point/" + period + "/" + pkg)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetRangeCounts gets the downloads per day for a given period, for all packages or a specific package.
// Period is the desired period to be queried for the range. If empty defaults to "last-week".
// Optionally pass in the package name in the pkg param.
func GetRangeCounts(period string, pkg string) (out *RangeCounts, err error) {
	if period == "" {
		period = "last-week"
	}

	body, err := call("/range/" + period + "/" + pkg)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// call rpc style endpoint.
func call(path string) (io.ReadCloser, error) {
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
