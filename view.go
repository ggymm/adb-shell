package main

import (
	"strings"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"

	"adb-shell/adb"
)

type MainForm struct {
	*vcl.TForm

	devices *vcl.TComboBox
	connect *vcl.TButton
	refresh *vcl.TButton

	init   *vcl.TButton
	remove *vcl.TButton

	active   *vcl.TButton
	inactive *vcl.TButton

	clear   *vcl.TButton
	content *vcl.TMemo
}

func (f *MainForm) OnFormCreate(_ vcl.IObject) {
	f.SetCaption("授权工具")
	f.SetWidth(720)
	f.SetHeight(480)
	f.SetPosition(types.PoScreenCenter)
	f.SetBorderStyle(types.BsSingle)

	f.setupView()
	f.setupEvent()

	// 模拟点击刷新设备
	f.refresh.Click()
}

func (f *MainForm) print(s string) {
	vcl.ThreadSync(func() {
		s = strings.Replace(s, "\r\n", "\n", -1)
		lines := strings.Split(s, "\n")
		for _, line := range lines {
			if len(line) == 0 {
				continue
			}
			f.content.Lines().Add(line)
		}
	})
}

func (f *MainForm) enable() {
	f.devices.SetEnabled(true)
	f.connect.SetEnabled(true)
	f.refresh.SetEnabled(true)

	f.init.SetEnabled(true)
	f.remove.SetEnabled(true)
	f.active.SetEnabled(true)
	f.inactive.SetEnabled(true)

	f.clear.SetEnabled(true)
	f.content.SetEnabled(true)
}

func (f *MainForm) disable() {
	f.devices.SetEnabled(false)
	f.connect.SetEnabled(false)
	f.refresh.SetEnabled(false)

	f.init.SetEnabled(false)
	f.remove.SetEnabled(false)
	f.active.SetEnabled(false)
	f.inactive.SetEnabled(false)

	f.clear.SetEnabled(false)
	f.content.SetEnabled(false)
}

func (f *MainForm) setupView() {
	enFont := vcl.NewFont()
	//enFont.SetName("JetBrains Mono Medium")
	enFont.SetName("Microsoft YaHei UI")

	zhFont := vcl.NewFont()
	zhFont.SetName("Microsoft YaHei UI")

	f.devices = vcl.NewComboBox(f)
	f.devices.SetParent(f)
	f.devices.SetFont(zhFont)
	f.devices.SetStyle(types.CsDropDownList)
	f.devices.SetOnCloseUp(func(sender vcl.IObject) {
	})

	f.connect = vcl.NewButton(f)
	f.connect.SetParent(f)
	f.connect.SetFont(zhFont)
	f.connect.SetCaption("连接设备")

	f.refresh = vcl.NewButton(f)
	f.refresh.SetParent(f)
	f.refresh.SetFont(zhFont)
	f.refresh.SetCaption("刷新设备")

	f.devices.SetBounds(20, 25, 480, 25)
	f.connect.SetBounds(520, 25, 80, 25)
	f.refresh.SetBounds(620, 25, 80, 25)

	f.init = vcl.NewButton(f)
	f.init.SetParent(f)
	f.init.SetFont(zhFont)
	f.init.SetCaption("安装应用")

	f.remove = vcl.NewButton(f)
	f.remove.SetParent(f)
	f.remove.SetFont(zhFont)
	f.remove.SetCaption("卸载应用")

	f.active = vcl.NewButton(f)
	f.active.SetParent(f)
	f.active.SetFont(zhFont)
	f.active.SetCaption("设备授权")

	f.inactive = vcl.NewButton(f)
	f.inactive.SetParent(f)
	f.inactive.SetFont(zhFont)
	f.inactive.SetCaption("解除授权")

	f.clear = vcl.NewButton(f)
	f.clear.SetParent(f)
	f.clear.SetFont(zhFont)
	f.clear.SetCaption("清空日志")

	f.init.SetBounds(20, 70, 80, 25)
	f.remove.SetBounds(120, 70, 80, 25)
	f.active.SetBounds(220, 70, 80, 25)
	f.inactive.SetBounds(320, 70, 80, 25)
	f.clear.SetBounds(620, 70, 80, 25)

	f.content = vcl.NewMemo(f)
	f.content.SetParent(f)
	f.content.SetFont(enFont)
	f.content.SetScrollBars(types.SsVertical)
	f.content.SetReadOnly(true)
	f.content.SetWordWrap(false)

	f.content.SetBounds(20, 115, 680, 340)
}

func (f *MainForm) setupEvent() {
	device := func() string {
		// 获取设备
		if f.devices.ItemIndex() < 0 {
			vcl.ShowMessage("请选择设备")
			return ""
		}
		device := f.devices.Items().S(f.devices.ItemIndex())
		fields := strings.Fields(device)
		if len(fields) < 1 {
			vcl.ShowMessage("设备信息错误")
			return ""
		}
		return fields[0]
	}

	f.connect.SetOnClick(func(sender vcl.IObject) {
		defer f.enable()
		f.disable()

		out, err := adb.Connect()
		if err != nil {
			f.print(err.Error())
		}
		f.print(out)
	})
	f.refresh.SetOnClick(func(sender vcl.IObject) {
		defer f.enable()
		f.disable()

		devices, err := adb.Devices()
		if err != nil {
			f.print(err.Error())
			return
		}
		vcl.ThreadSync(func() {
			f.devices.Items().Clear()
			for _, dev := range devices {
				if dev.Status != "device" {
					f.print(dev.String())
					continue
				}
				f.devices.Items().Add(dev.String())
			}

			// 默认选中第一个
			if len(devices) > 0 {
				f.devices.SetItemIndex(0)
			}
		})
	})

	f.init.SetOnClick(func(sender vcl.IObject) {
		defer f.enable()
		f.disable()

		// 获取设备
		dev := device()
		if dev == "" {
			return
		}

		// 解除授权
		out, err := adb.InstallApk(dev, InstallApk)
		if err != nil {
			f.print(err.Error())
		}
		f.print(out)
	})
	f.remove.SetOnClick(func(sender vcl.IObject) {
		defer f.enable()
		f.disable()

		// 获取设备
		dev := device()
		if dev == "" {
			return
		}

		// 解除授权
		out, err := adb.Uninstall(dev, PackageName)
		if err != nil {
			f.print(err.Error())
		}
		f.print(out)
	})

	f.active.SetOnClick(func(sender vcl.IObject) {
		defer f.enable()
		f.disable()

		// 获取设备
		dev := device()
		if dev == "" {
			return
		}

		// 激活设备
		out, err := adb.EnableOwner(dev, ReceiverClasspath)
		if err != nil {
			f.print(err.Error())
		}
		f.print(out)
	})
	f.inactive.SetOnClick(func(sender vcl.IObject) {
		defer f.enable()
		f.disable()

		// 获取设备
		dev := device()
		if dev == "" {
			return
		}

		// 解除授权
		out, err := adb.InstallApk(dev, UninstallApk)
		if err != nil {
			f.print(err.Error())
		}
		f.print(out)
	})

	f.clear.SetOnClick(func(sender vcl.IObject) {
		defer f.enable()
		f.disable()

		f.content.Lines().Clear()
	})
}
