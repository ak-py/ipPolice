package core

func GetMessage(msgCode string) string {
	locale := GetConfiguration().Locale
	switch locale {
	case "EN":
		return messagesEN[msgCode]
	// TODO - add more languages later for internationalizing
	default:
		return messagesEN[msgCode]
	}
}

const (
	McIpInvalid        = "IP_INVALID"
	McIpNil            = "IP_NIL"
	McErrorGeoDb       = "ERROR_GEO_DB"
	McErrorCountryApi  = "ERROR_COUNTRY_API"
	McCountryNotFound  = "COUNTRY_NOT_FOUND"
	McMethodNotAllowed = "METHOD_NOT_ALLOWED"
	McBadRequest       = "BAD_REQUEST"
)

var messagesEN = map[string]string{
	McIpInvalid:        "Ip Address is invalid",
	McIpNil:            "IP passed to Lookup cannot be nil",
	McErrorGeoDb:       "Error encountered using geo database",
	McErrorCountryApi:  "Error encountered using country api",
	McCountryNotFound:  "Country not found",
	McMethodNotAllowed: "Method is not allowed. Please check the api documentation.",
	McBadRequest:       "Bad Request. Please check the api documentation.",
}
