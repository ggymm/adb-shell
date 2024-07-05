package adb

import (
	"testing"
)

func Test_Start(t *testing.T) {
	Start()
}

func Test_Connect(t *testing.T) {
	Connect()
}

func Test_Devices(t *testing.T) {
	devices, err := Devices()
	if err != nil {
		t.Fatal(err)
	}
	for _, device := range devices {
		t.Logf("%+v", device)
	}
}
