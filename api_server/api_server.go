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

// ListenAndServe listens on the TCP network and serves REST API requests.
//
// @title IpPolice API
// @version v1
// @description This is the REST API provided by ipPolice.
// @contact.name akg009
// @contact.email contact@akgupta.tech
// @license.name akg009
// @in header
// @host localhost
// @BasePath /ipPolice/v1
func ListenAndServe() {
	http.HandleFunc(RestApiValidateIpPath, HandleValidateIP)
	http.HandleFunc(RestApiIpToCountryPath, HandleIpToCountry)

	port := core.GetConfiguration().RestApiPort
	address := ":" + port
	core.Logger().Println(core.LogInfoVariable("Starting api server on port", port))
	core.Logger().Fatal(http.ListenAndServe(address, nil))
}
