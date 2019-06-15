package rest

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Send(resp http.ResponseWriter, data map[string]interface{}) {
	resp.Header().Add("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(data)
}
