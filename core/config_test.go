package core

import "testing"

func TestSetConfiguration(t *testing.T) {
	defaultConfig := defaultConfig()

	newConfig := IpPoliceConfig{DatabaseFilePath: "fake_file.txt"}
	config := SetConfiguration(newConfig)

	if config == defaultConfig || config != newConfig {
		t.Errorf("config was incorrect, got: %#v, want: %#v.", config, newConfig)
	}
}
