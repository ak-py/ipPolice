package core

import "testing"

func TestGetMessage_Local_JK(t *testing.T) {
	SetConfiguration(IpPoliceConfig{
		DatabaseFilePath: "random.txt",
		Locale:           "JK",
	})
	msg := GetMessage(McCountryNotFound)
	expected := "Country not found"
	if msg != expected {
		t.Errorf("msg was incorrect, got: %#v, want: %#v.", msg, expected)
	}
}
