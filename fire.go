package fire

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var defaultUserAgent = "fire"

var methods = map[string]bool{
	"GET":    true,
	"POST":   true,
	"PUT":    true,
	"DELETE": true,
	"PATCH":  true,
}

type Request struct {
	Method  string            `json:"METHOD"`
	URL     string            `json:"URL"`
	Headers map[string]string `json:"HEADERS"`
	Payload map[string]string `json:"PAYLOAD"`
}

func IsSupportedMethod(check string) bool {
	if _, ok := methods[check]; !ok {
		return false
	}
	return true
}

func IsValidURL(check string) bool {
	_, err := url.ParseRequestURI(check)
	if err != nil {
		return false
	}
	return true
}

func (r *Request) hasUserAgent() bool {
	if _, ok := r.Headers["User-Agent"]; !ok {
		return false
	}
	return true
}

func (r *Request) Fire() error {

	if !IsSupportedMethod(r.Method) {
		return errors.New("Unsupported http-method passed.")
	}

	if !IsValidURL(r.URL) {
		return errors.New("Invalid URL passed.")
	}

	req, err := http.NewRequest(r.Method, r.URL, nil)
	if err != nil {
		return err
	}

	if !r.hasUserAgent() {
		r.Headers["User-Agent"] = defaultUserAgent
	}

	for header, value := range r.Headers {
		req.Header.Set(header, value)
	}

	if r.Method == "GET" {
		payload := req.URL.Query()
		for key, value := range r.Payload {
			payload.Add(key, value)
		}
		req.URL.RawQuery = payload.Encode()
	} else {
		payload := url.Values{}
		for key, value := range r.Payload {
			payload.Add(key, value)
		}
		req.Form = payload
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	outHeader, _ := json.Marshal(req.Header)
	outForm, _ := json.Marshal(req.Form)

	fmt.Printf("REQUEST:\t[%s]\t%s\nHEADERS:\t%s\nPAYLOAD:\t%s\nRESPONSE:\t%d\n", r.Method, req.URL, string(outHeader), string(outForm), resp.StatusCode)

	return nil
}
