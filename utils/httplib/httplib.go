package httplib

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var (
	client = http.Client{Timeout: 10 * time.Second}
)

const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
)

func Request(method, url string, data interface{}, headers map[string]string) (*http.Response, error) {
	var dataParams []byte
	var err error
	if data != nil {
		dataParams, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(dataParams))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return client.Do(req)
}

func RequestBind(method, url string, data interface{}, headers map[string]string, obj interface{}) (*http.Response, error) {
	resp, err := Request(method, url, data, headers)
	if err != nil {
		return nil, err
	}
	if obj == nil {
		return resp, nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(body, obj)
	return resp, err
}

func Get(url string, data interface{}, headers map[string]string) (*http.Response, error) {
	return RequestBind(MethodGet, url, data, headers, nil)
}

func GetInto(url string, data interface{}, headers map[string]string, obj interface{}) (*http.Response, error) {
	return RequestBind(MethodGet, url, data, headers, obj)
}

func Post(url string, data interface{}, headers map[string]string) (*http.Response, error) {
	return RequestBind(MethodPost, url, data, headers, nil)
}

func PostInto(url string, data interface{}, headers map[string]string, obj interface{}) (*http.Response, error) {
	return RequestBind(MethodPost, url, data, headers, obj)
}

func Put(url string, data interface{}, headers map[string]string) (*http.Response, error) {
	return RequestBind(MethodPut, url, data, headers, nil)
}

func PutInto(url string, data interface{}, headers map[string]string, obj interface{}) (*http.Response, error) {
	return RequestBind(MethodPut, url, data, headers, obj)
}

func Delete(url string, data interface{}, headers map[string]string) (*http.Response, error) {
	return RequestBind(MethodDelete, url, data, headers, nil)
}

func DeleteInto(url string, data interface{}, headers map[string]string, obj interface{}) (*http.Response, error) {
	return RequestBind(MethodDelete, url, data, headers, obj)
}
