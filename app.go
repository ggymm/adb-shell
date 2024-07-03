package main

import (
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ying32/govcl/vcl"

	"adb-shell/adb"
)

type App struct {
	wd   string
	view *MainForm
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	a.init()
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

func (a *App) showView() {
	vcl.DEBUG = false
	vcl.Application.SetScaled(true)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	// 主窗口
	vcl.Application.CreateForm(&a.view)

	// 显示事件
	a.view.SetOnShow(func(sender vcl.IObject) {
		err := adb.Start()
		if err != nil {
			slog.Error("adb start error", err)
		}
	})

	// 启动应用
	vcl.Application.Run()
}
