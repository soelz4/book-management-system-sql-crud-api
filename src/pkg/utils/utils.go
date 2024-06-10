package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err == nil {
		// Unmarshal Parses the JSON-Encoded Data and Stores the Result
		err = json.Unmarshal([]byte(body), x)
	}
	if err != nil {
		return
	}
}
