package api_server

import (
	"encoding/json"
	"io/ioutil"
	"ipPolice/core"
	"net/http"
	"strings"
)

func HandleValidateIP(rw http.ResponseWriter, req *http.Request) {
	core.Logger().Println(LogHttpReqInfo(req))

	if !strings.EqualFold(req.Method, http.MethodPost) {
		failureResponseCompose(rw, core.McMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	var reqBody core.HttpReqValidateIp
	reqBodyBytes, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		core.Logger().Println(core.LogErrorMessage(readErr))
		failureResponseCompose(rw, core.McBadRequest, http.StatusBadRequest)
		return
	}

	if unMarshalErr := json.Unmarshal(reqBodyBytes, &reqBody); unMarshalErr != nil {
		core.Logger().Println(core.LogErrorMessage(unMarshalErr))
		failureResponseCompose(rw, core.McBadRequest, http.StatusBadRequest)
		return
	}

	// todo - validate request json

	ip := reqBody.IpAddress
	whiteList := reqBody.WhiteListIsoCodes

	isWhiteListed, err := core.ValidateIP(ip, whiteList)
	if err != nil {
		core.Logger().Println(core.LogErrorMessage(err))
		failureResponseCompose(rw, core.McCountryNotFound, http.StatusInternalServerError)
		return
	}

	response := core.HttpResValidateIp{
		Status:        core.StatusSuccess,
		IsWhiteListed: isWhiteListed,
	}
	customResponseCompose(rw, response)
	return
}
