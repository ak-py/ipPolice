package api_server

import (
	"fmt"
	"ipPolice/core"
	"net/http"
	"time"
)

func LogHttpReqInfo(req *http.Request) string {
	return core.LogInfoMessage(getHttpReqInfo(req))
}

func getHttpReqInfo(req *http.Request) string {

	method := req.Method
	route := req.URL.Path
	currTime := time.Now()

	msg := fmt.Sprintf("\t %#v \t %#v \t %#v", method, route, currTime)
	return msg
}
