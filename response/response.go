package response

import (
	"encoding/json"
	"net/http"
	"time"
)

type setResponse struct {
	Status     string      `json:"status"`
	AccessTime string      `json:"access_time"`
	Data       interface{} `json:"data"`
}

func HttpResponseSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	setResponse := setResponse{
		Status:     http.StatusText(http.StatusOK),
		AccessTime: time.Now().Format("02-01-2006 15:04:05"),
		Data:       data,
	}
	response, _ := json.Marshal(setResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func HttpResponseError(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	setResponse := setResponse{
		Status:     http.StatusText(code),
		AccessTime: time.Now().Format("02-01-2006 15:04:05"),
		Data:       data,
	}

	response, _ := json.Marshal(setResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(response)
}
