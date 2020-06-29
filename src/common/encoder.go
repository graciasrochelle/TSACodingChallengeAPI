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
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Header().Set("Access-Control-Allow-Origin", "http://graciasrochelle.github.io")
	w.WriteHeader(status)
	resp, e := json.Marshal(response)
	if e != nil {
		return e
	}
	WriteWithPayload(w, resp)
	return nil
}
