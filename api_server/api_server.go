package api_server

import (
	"ipPolice/core"
	"net/http"
)

var (
	RestApiBasePath        = "/" + "ipPolice" + "/" + core.GetConfiguration().Version + "/"
	RestApiValidateIpPath  = RestApiBasePath + "validate/"
	RestApiIpToCountryPath = RestApiBasePath + "ipToCountry/"
)

func ListenAndServe() {
	http.HandleFunc(RestApiValidateIpPath, HandleValidateIP)
	http.HandleFunc(RestApiIpToCountryPath, HandleIpToCountry)

	port := core.GetConfiguration().RestApiPort
	address := ":" + port
	core.Logger().Println(core.LogInfoVariable("Starting api server on port", port))
	core.Logger().Fatal(http.ListenAndServe(address, nil))
}
