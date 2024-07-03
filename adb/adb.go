package adb

import (
	"fmt"
	"strings"
)

type Device struct {
	Name    string
	Model   string
	Brand   string
	Version string
}

func (d *Device) String() string {
	return fmt.Sprintf("%s %s %s Android%s", d.Name, d.Model, d.Brand, d.Version)
}

func Start() error {
	_, err := Exec("adb", "start-server")
	return err
}

func Connect() error {
	_, err := Exec("adb", "reconnect")
	return err
}

func Devices() ([]*Device, error) {
	out, err := Exec("adb", "devices")
	if err != nil {
		return nil, err
	}

	ds := make([]*Device, 0)
	lines := strings.Split(out, "\n")
	for i, line := range lines {
		if i == 0 {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		name := fields[0]

		// get device model
		out, err = Exec("adb", "-s", name, "shell", "getprop", "ro.product.model")
		if err != nil {
			continue
		}
		model := strings.TrimSpace(out)

		// get device brand
		out, err = Exec("adb", "-s", name, "shell", "getprop", "ro.product.brand")
		if err != nil {
			continue
		}
		brand := strings.TrimSpace(out)

		// get device version
		out, err = Exec("adb", "-s", name, "shell", "getprop", "ro.build.version.release")
		if err != nil {
			continue
		}
		version := strings.TrimSpace(out)

		ds = append(ds, &Device{
			Name:    name,
			Model:   model,
			Brand:   brand,
			Version: version,
		})
	}
	return ds, nil
}

func InstallApk(device, apk string) (string, error) {
	return Exec("adb", "-s", device, "install", "-r", apk)
}

func UninstallApk(device, pkg string) (string, error) {
	return Exec("adb", "-s", device, "uninstall", pkg)
}

func EnableDeviceOwner(device, classpath string) (string, error) {
	return Exec("adb", "-s", device, "shell", "dpm", "set-device-owner", classpath)
}
