package utils

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Code    int
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
	//Data    map[string]string `json:"data"`
}

func (r *Result) Response(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	response, _ := json.Marshal(r)
	w.Write(response)
}
