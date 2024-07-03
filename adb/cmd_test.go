package adb

import (
	"fmt"
	"testing"
	"time"
)

func Test_Exec(t *testing.T) {
	r, err := Exec("cmd", "/c", "chcp 65001 && dir")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func Test_ExecAsync(t *testing.T) {
	ms, errs := ExecAsync("cmd", "/c", "dir")
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
