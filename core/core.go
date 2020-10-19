package core

import (
	"errors"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)

func ValidateIP(ipAddress string, whiteListISOCodes []string) (bool, error) {
	Logger().Println(LogInfoVariable("Validating for Ip Address", ipAddress))

	countryISO, err := IpToCountryCode(ipAddress)

	if err != nil {
		Logger().Println(LogErrorMessage(err))
		return false, err
	}

	for _, isoCode := range whiteListISOCodes {
		if countryISO == isoCode {
			return true, nil
		}
	}

	return false, nil
}

func IpToCountryCode(ipAddress string) (string, error) {
	Logger().Println(LogInfoVariable("Finding country for Ip Address", ipAddress))

	ip := net.ParseIP(ipAddress)
	if ip == nil {
		Logger().Println(LogWarningMessage(GetMessage(McIpInvalid)))
	}

	dbPath := GetConfiguration().DatabaseFilePath
	if dbPath == "" {
		dbPath = dbFilePath()
	}

	geoReader, err := geoip2.Open(dbPath)
	if err != nil {
		Logger().Println(LogErrorMessage(err))
		Logger().Println(LogErrorMessageWithCode(McErrorGeoDb))
		err = errors.New(GetMessage(McErrorGeoDb))
		return "", err
	}
	defer geoReader.Close()

	country, err := geoReader.Country(ip)

	if err != nil {
		Logger().Println(LogErrorMessage(err))
		Logger().Println(LogErrorMessageWithCode(McErrorCountryApi))
		err = errors.New(GetMessage(McErrorCountryApi))
		return "", err
	}

	if country == nil || country.Country.IsoCode == "" {
		Logger().Println(LogErrorMessageWithCode(McCountryNotFound))
		err = errors.New(GetMessage(McCountryNotFound))
		return "", err
	}

	countryISO := country.Country.IsoCode

	infoMsg := fmt.Sprintf("Ip Address %#v maps to County ISO Code = %#v", ipAddress, countryISO)
	Logger().Println(LogInfoMessage(infoMsg))

	return countryISO, nil
}
