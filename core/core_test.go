package core

import (
	"testing"
)

func TestValidateIP_GB(t *testing.T) {
	SetConfiguration(defaultConfig())

	ip := "81.2.69.142"
	countries := []string{"USA", "GB"}

	isValid, err := ValidateIP(ip, countries)

	if !isValid {
		t.Errorf("isValid was incorrect, got: %#v, want: %#v.", isValid, true)
	}

	if err != nil {
		t.Errorf("err was incorrect, got: %#v, want: %#v.", err, nil)
	}

	// verify if GB not in array
	isValid, err = ValidateIP(ip, []string{"USA"})

	if isValid {
		t.Errorf("isValid was incorrect, got: %#v, want: %#v.", isValid, false)
	}

	if err != nil {
		t.Errorf("err was incorrect, got: %#v, want: %#v.", err, nil)
	}

}

func TestValidateIP_LocalHost(t *testing.T) {
	SetConfiguration(defaultConfig())

	ip := "127.0.0.1"
	countries := []string{"USA", "GB"}

	isValid, err := ValidateIP(ip, countries)

	if isValid {
		t.Errorf("isValid was incorrect, got: %#v, want: %#v.", isValid, true)
	}

	if err == nil {
		t.Errorf("err was incorrect, got: %#v, want: %#v.", err, "Country not found")
	}
}

func TestValidateIP_BadInput(t *testing.T) {
	SetConfiguration(defaultConfig())

	ip := "fake-ip"
	countries := []string{"USA", "GB"}

	isValid, err := ValidateIP(ip, countries)

	if isValid {
		t.Errorf("isValid was incorrect, got: %#v, want: %#v.", isValid, true)
	}

	if err == nil {
		t.Errorf("err was incorrect, got: %#v, want: %#v.", err, "IP passed to Lookup cannot be nil")
	}
}

func TestIpToCountryCode_GB(t *testing.T) {
	SetConfiguration(defaultConfig())

	gbIp := "81.2.69.142"
	countryCode, err := IpToCountryCode(gbIp)

	if countryCode != "GB" {
		t.Errorf("countryCode was incorrect, got: %#v, want: %#v.", countryCode, "GB")
	}

	if err != nil {
		t.Errorf("err was incorrect, got: %#v, want: %#v.", err, nil)
	}
}

func TestIpToCountryCode_LocalHost(t *testing.T) {
	SetConfiguration(defaultConfig())

	localhost := "127.0.0.1"
	countryCode, err := IpToCountryCode(localhost)

	if countryCode != "" {
		t.Errorf("countryCode was incorrect, got: %#v, want: %#v.", countryCode, "")
	}

	if err == nil {
		t.Errorf("err was incorrect, got: %#v, want: %#v.", err, "Country not found")
	}
}

func TestIpToCountryCode_BadInput(t *testing.T) {
	SetConfiguration(defaultConfig())

	badIp := "fake-ip"
	countryCode, err := IpToCountryCode(badIp)

	if countryCode != "" {
		t.Errorf("countryCode was incorrect, got: %#v, want: %#v.", countryCode, "")
	}

	if err == nil {
		t.Errorf("err was incorrect, got: %#v, want: %#v.", err, "IP passed to Lookup cannot be nil")
	}

}

func TestIpToCountryCode_BadConfig(t *testing.T) {
	SetConfiguration(IpPoliceConfig{DatabaseFilePath: "invalid_file_name.txt"})

	badIp := "127.0.0.1"
	countryCode, err := IpToCountryCode(badIp)

	if countryCode != "" {
		t.Errorf("countryCode was incorrect, got: %#v, want: %#v.", countryCode, "")
	}

	if err == nil {
		t.Errorf("err was incorrect, got: %#v, want: %#v.", err, "IP passed to Lookup cannot be nil")
	}

	SetConfiguration(defaultConfig())

	countryCode, err = IpToCountryCode(badIp)

	if countryCode != "" {
		t.Errorf("countryCode was incorrect, got: %#v, want: %#v.", countryCode, "")
	}

	if err == nil {
		t.Errorf("err was incorrect, got: %#v, want: %#v.", err, "IP passed to Lookup cannot be nil")
	}
}
