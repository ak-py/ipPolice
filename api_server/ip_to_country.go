package api_server

import (
	"ipPolice/core"
	"net/http"
	"strings"
)

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
