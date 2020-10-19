package api_server

import (
	"encoding/json"
	"io/ioutil"
	"ipPolice/core"
	"net/http"
	"strings"
)

// Handler for /validate POST
//
// @Summary validate - validate the ip address against given white list countries
// @Description This API is for validating the given ip address by using geo db.
// @ID ipPolice.validate.post
// @Tags ipPoliceService
// @Accept json
// @Produce  json
// @Param payload body core.HttpReqValidateIp true "JSON payload of the request"
// @Success 200 {object} core.HttpResValidateIp "OK"
// @Failure 400 {object} core.ApiStatus "Bad Request"
// @Failure 405 {object} core.ApiStatus "Method Not Allowed"
// @Failure 500 {object} core.ApiStatus "Internal Server Error"
// @Router /validate [post]
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
