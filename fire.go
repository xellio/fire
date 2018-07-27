package fire

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
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
	Method    string            `json:"METHOD"`
	URL       string            `json:"URL"`
	Headers   map[string]string `json:"HEADERS"`
	Payload   map[string]string `json:"PAYLOAD"`
	Timestamp int64
	Response  *Response
}

type Response struct {
	*http.Response
	Timestamp int64
	Duration  int64
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

func (r *Request) Fire() (*Response, error) {

	if !IsSupportedMethod(r.Method) {
		return r.Response, errors.New("Unsupported http-method passed.")
	}

	if !IsValidURL(r.URL) {
		return r.Response, errors.New("Invalid URL passed.")
	}

	req, err := http.NewRequest(r.Method, r.URL, nil)
	if err != nil {
		return r.Response, err
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
	r.Timestamp = time.Now().UnixNano()
	resp, err := client.Do(req)
	if err != nil {
		return r.Response, err
	}
	defer resp.Body.Close()
	respTs := time.Now().UnixNano()

	r.Response = &Response{
		resp,
		respTs,
		(respTs - r.Timestamp) / int64(time.Millisecond),
	}

	outHeader, _ := json.Marshal(req.Header)
	outForm, _ := json.Marshal(req.Form)

	fmt.Printf("REQUEST:\t[%s]\t%s\nHEADERS:\t%s\nPAYLOAD:\t%s\nRESPONSE:\t%d\nDURATION:\t%dms\n\n", r.Method, req.URL, string(outHeader), string(outForm), r.Response.StatusCode, r.Response.Duration)

	return r.Response, nil
}
