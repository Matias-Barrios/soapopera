package requesthanlder

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
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
func (r Request) Execute() (int, string, error) {
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
		Body:   ioutil.NopCloser(strings.NewReader(r.Payload)),
		Header: r.Headers,
	}
	resp, err := client.Do(request)
	if err != nil {
		return 500, "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 500, "", err
	}
	return resp.StatusCode, string(body), nil
}
