package adb_test

import (
	"adb-shell/adb"
	"fmt"
	"testing"
	"time"
)

func Test_Exec(t *testing.T) {
	r, err := adb.Exec("cmd", "/c", "chcp 65001 && dir")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func Test_Exec2(t *testing.T) {
	r, err := adb.Exec("adb", "-s", "ABGBB23130204411", "shell", "dpm", "set-device-owner", "com.ninelock.mobile/com.ninelock.mobile.core.manage.DeviceReceiver")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func Test_ExecAsync(t *testing.T) {
	ms, errs := adb.ExecAsync("cmd", "/c", "dir")
	go func() {
		for {
			select {
			case m := <-ms:
				fmt.Println(m)
			}
		}
	}()

	for {
		select {
		case err := <-errs:
			t.Fatalf("%+v", err)
		case <-time.After(5 * time.Second):
			return
		}
	}
}
