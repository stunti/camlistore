package http_util

import (
	"fmt"
	"http"
	"json"
	"os"
)

func BadRequestError(conn http.ResponseWriter, errorMessage string) {
	conn.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(conn, "%s\n", errorMessage)
}

func ServerError(conn http.ResponseWriter, err os.Error) {
	conn.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(conn, "Server error: %s\n", err)
}

func ReturnJson(conn http.ResponseWriter, data interface{}) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		BadRequestError(conn, fmt.Sprintf(
			"JSON serialization error: %v", err))
		return
	}
	conn.Write(bytes)
	conn.Write([]byte("\n"))
}