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
	w.Header().Set("Content-Type", "application/json;")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS,GET,PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,crossdomain, mode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	resp, e := json.Marshal(response)
	if e != nil {
		return e
	}
	WriteWithPayload(w, resp)
	return nil
}
