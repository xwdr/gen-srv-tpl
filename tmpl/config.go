package tmpl

var Config = `
# 日志配置
[app]
PageSize = 10
JwtSecret = 233
PrefixUrl = http://127.0.0.1:8000

# MB
ImageMaxSize = 5
ImageAllowExts = .jpg,.jpeg,.png

ExportSavePath = export/
QrCodeSavePath = qrcode/
FontSavePath = fonts/

LogSavePath = logs/
LogSaveName = log
LogFileExt = log
TimeFormat = 20060102

[server]
#debug or release
RunMode = debug
HttpPort = 8000
ReadTimeout = 60
WriteTimeout = 60

[database]
Type = mysql
User = root
Password = x2014.mc
Host = 129.204.12.208:3306
Name = sdata
TablePrefix = blog_

[redis]
Host = 192.168.19.131:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200

`

var Module = `module {{.Dir}}

go 1.14

require (
	github.com/gin-gonic/gin v1.7.1
	github.com/jinzhu/gorm v1.9.16
	github.com/sirupsen/logrus v1.8.1
)

`