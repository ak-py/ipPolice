package api_server

import (
	"encoding/json"
	"ipPolice/core"
	"net/http"
)

const JsonIndentation = "    "

func customResponseCompose(rw http.ResponseWriter, data interface{}) {
	core.Logger().Println(core.LogInfoVariable("Composing custom response", data))
	js, _ := json.MarshalIndent(data, "", JsonIndentation)
	composeJsonResponse(rw, js, 0)
}

func failureResponseCompose(rw http.ResponseWriter, msgCode string, httpStatusCode int) {
	apiStatus := core.ApiStatus{
		Status:      core.StatusFailed,
		MessageCode: msgCode,
		Message:     core.GetMessage(msgCode),
	}
	core.Logger().Println(core.LogInfoVariable("Composing failure response. Api Status", apiStatus))
	js, _ := json.MarshalIndent(apiStatus, "", JsonIndentation)
	composeJsonResponse(rw, js, httpStatusCode)
}

func composeJsonResponse(rw http.ResponseWriter, js []byte, httpStatusCode int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("X-Content-Type-Options", "nosniff")
	if httpStatusCode != 0 {
		rw.WriteHeader(httpStatusCode)
	}
	_, _ = rw.Write(js)
}
