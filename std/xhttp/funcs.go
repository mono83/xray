package xhttp

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
)

// Do performs HTTP request
func Do(ray xray.Ray, client *http.Client, req *http.Request) (*http.Response, error) {
	if ray == nil {
		ray = xray.ROOT.Fork().WithLogger("http")
	}
	if req == nil {
		return nil, errors.New("empty request")
	}
	if client == nil {
		client = &http.Client{}
	}

	ray = ray.With(args.URL(req.URL.String()))

	ray.Debug("Sending request to :url")
	before := time.Now()

	resp, err := client.Do(req)
	if err != nil {
		ray.Error("Request to :url failed - :err", args.Error{Err: err})
		return nil, err
	}

	ray.Debug(
		"Request to :url finished with code :code in :delta",
		args.Int{N: "code", V: resp.StatusCode},
		args.Delta(time.Now().Sub(before)),
	)

	return resp, nil
}

// DoRead performs HTTP request and reads response body
func DoRead(ray xray.Ray, client *http.Client, req *http.Request) (code int, body []byte, headers http.Header, cookies []*http.Cookie, err error) {
	if ray == nil {
		ray = xray.ROOT.Fork().WithLogger("http")
	}

	resp, err := Do(ray, client, req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	code = resp.StatusCode
	headers = resp.Header
	cookies = resp.Cookies()
	body, err = ioutil.ReadAll(resp.Body)
	if err == nil {
		ray.InBytes(body)
	}

	return
}

// DoReadGet performs GET request and reads all response
func DoReadGet(ray xray.Ray, client *http.Client, url string) (code int, body []byte, headers http.Header, cookies []*http.Cookie, err error) {
	// Building request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	return DoRead(ray, client, req)
}
