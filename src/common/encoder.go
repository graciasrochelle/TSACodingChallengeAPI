package common

import (
	"encoding/json"
	"net/http"
)

func WriteWithPayload(w http.ResponseWriter, payload ResponseType) {
	if payload != nil {
		body := payload.([]byte)
		w.Write(body)
	}
}

func EncodeJsonResponse(w http.ResponseWriter, status int, response ResponseType) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	resp, e := json.Marshal(response)
	if e != nil {
		return e
	}
	WriteWithPayload(w, resp)
	return nil
}

func AddHeadersToResponse(w http.ResponseWriter, headers map[string]string) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
}
