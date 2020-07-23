package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HttpClient struct {
	uri     string
	method  string
	timeout time.Duration
	client  *http.Client
}

func NewHttpClient(uri string) *HttpClient {
	c := &http.Client{
		Timeout: 3 * time.Second,
	}
	return &HttpClient{
		uri:     uri,
		method:  "POST",
		timeout: 3 * time.Second,
		client:  c,
	}
}

func (h *HttpClient) SetUri(uri string) {
	h.uri = uri
}

func (h *HttpClient) SetMethod(method string) {
	h.method = method
}

func (h *HttpClient) SetTimeout(timeout time.Duration) {
	h.timeout = timeout
}

func (h *HttpClient) Push(payload *Event) error {
	payload.Validate()
	b, err := json.Marshal(payload)

	req, err := http.NewRequest(h.method, h.uri, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	resp, err := h.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := make(map[string][]string)
	json.Unmarshal(body, res)
	if failIds, ok := res["fail"]; ok {
		if len(failIds) > 0 {
			return fmt.Errorf("following ids failed, %s", strings.Join(failIds, ","))
		}
	}

	return nil
}
