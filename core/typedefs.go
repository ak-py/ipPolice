package core

const (
	StatusSuccess = "SUCCESS"
	StatusFailed  = "FAILED"
)

type IpPoliceConfig struct {
	Version          string `json:"version"`
	Locale           string `json:"locale"`
	RestApiPort      string `json:"rest_api_port"`
	DatabaseFilePath string `json:"database_file_path"`
	WorkingDir       string `json:"working_dir"`
}

type ApiStatus struct {
	Status      string `json:"status"`
	MessageCode string `json:"message_code"`
	Message     string `json:"message"`
}

type HttpReqValidateIp struct {
	IpAddress         string   `json:"ip_address"`
	WhiteListIsoCodes []string `json:"white_list_iso_codes"`
}
type HttpResValidateIp struct {
	Status        string `json:"status"`
	IsWhiteListed bool   `json:"is_white_listed"`
}

type HttpReqIpToCountry struct{}

type HttpResIpToCountry struct {
	Status         string `json:"status"`
	CountryIsoCode string `json:"country_iso_code"`
}
