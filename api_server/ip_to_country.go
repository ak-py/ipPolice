package api_server

import (
	"ipPolice/core"
	"net/http"
	"strings"
)

// Handler for /ipToCountry GET
//
// @Summary ipToCountry - get country from ip address
// @Description This API is for getting the country iso code for given ip address by using geo db.
// @ID ipPolice.ipçToCountry.get
// @Tags ipPoliceService
// @Accept json
// @Produce  json
// @Param ip path string true "Ip Address"
// @Param payload body core.HttpReqIpToCountry true "JSON payload of the request"
// @Success 200 {object} core.HttpResIpToCountry "OK"
// @Failure 400 {object} core.ApiStatus "Bad Request"
// @Failure 405 {object} core.ApiStatus "Method Not Allowed"
// @Failure 500 {object} core.ApiStatus "Internal Server Error"
// @Router /ipToCountry/{ip} [get]
func HandleIpToCountry(rw http.ResponseWriter, req *http.Request) {
	core.Logger().Println(LogHttpReqInfo(req))

	if !strings.EqualFold(req.Method, http.MethodGet) {
		failureResponseCompose(rw, core.McMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	ip := strings.TrimPrefix(req.URL.Path, RestApiIpToCountryPath)
	core.Logger().Println(core.LogInfoVariable("ip", ip))

	countryIsoCode, err := core.IpToCountryCode(ip)
	if err != nil {
		core.Logger().Println(core.LogErrorMessage(err))
		failureResponseCompose(rw, core.McErrorCountryApi, http.StatusInternalServerError)
		return
	}

	response := core.HttpResIpToCountry{
		Status:         core.StatusSuccess,
		CountryIsoCode: countryIsoCode,
	}
	customResponseCompose(rw, response)
	return
}
