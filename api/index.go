package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/vscodev/alist-auth-api/conf"
	"github.com/vscodev/alist-auth-api/drivers/alipan_tv"
	"github.com/vscodev/alist-auth-api/drivers/aliyun_drive"
	"github.com/vscodev/alist-auth-api/drivers/baidu_netdisk"
	"github.com/vscodev/alist-auth-api/drivers/box"
	"github.com/vscodev/alist-auth-api/drivers/disk_115"
	"github.com/vscodev/alist-auth-api/drivers/drive_123"
	"github.com/vscodev/alist-auth-api/drivers/dropbox"
	"github.com/vscodev/alist-auth-api/drivers/google_drive"
	"github.com/vscodev/alist-auth-api/drivers/onedrive"
	"github.com/vscodev/alist-auth-api/drivers/pcloud"
	"github.com/vscodev/alist-auth-api/drivers/yandex_disk"
)

var e *echo.Echo

func init() {
	e = echo.New()
	e.Use(middleware.CORS())
	e.GET("/config", getConfig)
	e.POST("/aliyun_drive/token", aliyun_drive.GetToken)
	e.POST("/alipan_tv/qrcode", alipan_tv.GetQrCode)
	e.GET("/alipan_tv/qrcode/:sid/status", alipan_tv.CheckQrCodeStatus)
	e.POST("/alipan_tv/token", alipan_tv.GetToken)
	e.POST("/baidu_netdisk/token", baidu_netdisk.GetToken)
	e.POST("/115_disk/qrcode", disk_115.GetQrCode)
	e.POST("/115_disk/device_code_token", disk_115.GetTokenViaDeviceCode)
	e.POST("/115_disk/auth_code_token", disk_115.GetTokenViaAuthCode)
	e.POST("/123_drive/token", drive_123.GetToken)
	e.POST("/onedrive/token", onedrive.GetToken)
	e.POST("/google_drive/token", google_drive.GetToken)
	e.POST("/dropbox/token", dropbox.GetToken)
	e.POST("/box/token", box.GetToken)
	e.POST("/yandex_disk/token", yandex_disk.GetToken)
	e.POST("/pcloud/token", pcloud.GetToken)
	e.RouteNotFound("/*", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "https://alist-auth.pages.dev")
	})
}

func getConfig(c echo.Context) error {
	return c.JSON(http.StatusOK, conf.Conf.Public())
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e.ServeHTTP(w, r)
}

func Run(host string, port int) {
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", host, port)))
}
