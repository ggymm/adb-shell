package adb

import (
	"testing"
)

func Test_Start(t *testing.T) {
	err := Start()
	if err != nil {
		t.Error(err)
	}
}

func Test_Connect(t *testing.T) {
	err := Connect()
	if err != nil {
		t.Error(err)
	}
}

func Test_Devices(t *testing.T) {
	devices, err := Devices()
	if err != nil {
		t.Error(err)
	}
	for _, device := range devices {
		t.Logf("%+v", device)
	}
}
