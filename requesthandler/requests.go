package requesthanlder

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

// Request:
type Request struct {
	URL     string
	Headers map[string][]string
	Method  string
	Payload string
}

// Execute:
func (Request r) Execute() (int, string, error) {
	client := http.Client{
		Timeout: time.Second * 60,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	request := &http.Request{
		Host:   r.URL,
		Body:   r.Paload,
		Header: r.Headers,
	}
	resp, err := client.Do(request)
	if err != nil {
		return "", 500, err.Error()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 500, err.Error()
	}
	return resp.StatusCode, body, nil
}
