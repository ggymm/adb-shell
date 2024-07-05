package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/ying32/govcl/vcl"

	"adb-shell/adb"
)

const (
	InstallApk   = "install.apk"
	UninstallApk = "uninstall.apk"

	PackageName       = "com.ninelock.mobile"
	ReceiverClasspath = "com.ninelock.mobile/com.ninelock.mobile.core.manage.DeviceReceiver"
)

type App struct {
	wd      string
	view    *MainForm
	devices []*adb.Device
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	a.init()
	a.runTask()
	a.showView()
}

func (a *App) init() {
	dir := ""
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path := filepath.Base(exe)
	if !strings.HasPrefix(exe, os.TempDir()) && !strings.HasPrefix(path, "___") {
		dir = filepath.Dir(exe)
	} else {
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			// 需要根据当前文件所处目录，修改相对位置
			dir = filepath.Dir(filename)
		}
	}
	a.wd = filepath.Join(dir, "data")
}

func (a *App) runTask() {
	t := time.NewTicker(3 * time.Second)
	go func() {
		for {
			select {
			case <-t.C:
				func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Printf("recover:%v\n", err)
						}
					}()

					_ = adb.Start()
				}()
			}
		}
	}()
}

func (a *App) showView() {
	vcl.RunApp(&a.view)
}
