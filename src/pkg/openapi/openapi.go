package openapi

import (
	"encoding/json"
	"net/http"
)

func Resp(w http.ResponseWriter, status int, response interface{}) {
	bytes, _ := json.Marshal(response)
	w.WriteHeader(status)
	_, _ = w.Write(bytes)
}
