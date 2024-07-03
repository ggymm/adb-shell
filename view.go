package main

import (
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type MainForm struct {
	*vcl.TForm
	combo *vcl.TComboBox

	devices *vcl.TComboBox
	connect *vcl.TButton
	refresh *vcl.TButton

	init    *vcl.TButton
	enable  *vcl.TButton
	disable *vcl.TButton

	clear   *vcl.TButton
	content *vcl.TMemo
}

func (f *MainForm) OnFormCreate(_ vcl.IObject) {
	f.SetCaption("授权工具")
	f.SetWidth(720)
	f.SetHeight(480)
	f.SetPosition(types.PoScreenCenter)
	f.SetBorderStyle(types.BsSingle)

	f.combo = vcl.NewComboBox(f)
	f.combo.SetParent(f)
	f.combo.SetBounds(0, 0, 0, 0)
	f.combo.SetStyle(types.CsDropDownList)

	f.setupContent()
}

func (f *MainForm) fixFocus() {
	f.combo.SetFocus()
}

func (f *MainForm) setupContent() {
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
		f.fixFocus()
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

	f.enable = vcl.NewButton(f)
	f.enable.SetParent(f)
	f.enable.SetFont(zhFont)
	f.enable.SetCaption("设备授权")

	f.disable = vcl.NewButton(f)
	f.disable.SetParent(f)
	f.disable.SetFont(zhFont)
	f.disable.SetCaption("解除授权")

	f.clear = vcl.NewButton(f)
	f.clear.SetParent(f)
	f.clear.SetFont(zhFont)
	f.clear.SetCaption("清空日志")

	f.init.SetBounds(20, 70, 80, 25)
	f.enable.SetBounds(120, 70, 80, 25)
	f.disable.SetBounds(220, 70, 80, 25)
	f.clear.SetBounds(620, 70, 80, 25)

	f.content = vcl.NewMemo(f)
	f.content.SetParent(f)
	f.content.SetFont(enFont)
	f.content.SetScrollBars(types.SsVertical)
	f.content.SetWordWrap(true)

	f.content.SetBounds(20, 115, 680, 340)
}
