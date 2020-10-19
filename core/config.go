package core

const DefaultDbFileName = "GeoIP2-Country-Test.mmdb"

var ipPoliceConfig = defaultConfig()

func defaultConfig() IpPoliceConfig {
	return IpPoliceConfig{
		Version:     "v1",
		Locale:      "EN",
		RestApiPort: "8080",
		// TODO - refactor file path later
		WorkingDir:       GetCurrentWorkingDirectory(),
		DatabaseFilePath: "",
	}
}

func GetConfiguration() IpPoliceConfig {
	//Logger().Println(LogInfoVariable("Ip Police Configuration", ipPoliceConfig))
	return ipPoliceConfig
}

func SetConfiguration(config IpPoliceConfig) IpPoliceConfig {

	// TODO - add validation
	ipPoliceConfig = config
	return ipPoliceConfig
}
