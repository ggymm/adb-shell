# adb-shell


```shell

# 启动 adb 服务
adb start-server

# 重新连接设备
adb reconnect

# 查看设备列表
adb devices
# 如果需要指定设备
# 在 adb 命令后面加上 -s 设备号

# 获取设备信息
adb shell getprop ro.product.model
adb shell getprop ro.product.brand
adb shell getprop ro.product.manufacturer
adb shell getprop ro.build.version.release

# 启用设备管理器
adb shell dpm set-device-owner com.ninelock.mobile/com.ninelock.mobile.core.manage.DeviceReceiver

# 安装应用
adb install -r app.apk

# 卸载应用
adb uninstall com.ninelock.mobile

```