package aliyun_drive

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/vscodev/alist-auth-api/conf"
)

type TokenReq struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`
	GrantType    string `json:"grant_type" form:"grant_type"`
	Code         string `json:"code,omitempty" form:"code"`
	RefreshToken string `json:"refresh_token,omitempty" form:"refresh_token"`
}

func GetToken(c echo.Context) error {
	r := new(TokenReq)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if r.ClientID == "" || r.ClientSecret == "" {
		r.ClientID = conf.Conf.AliyunDrive.ClientID
		r.ClientSecret = conf.Conf.AliyunDrive.ClientSecret
	}

	data, _ := json.Marshal(r)
	req, _ := http.NewRequestWithContext(c.Request().Context(), http.MethodPost, "https://openapi.alipan.com/oauth/access_token", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	responseHeader := c.Response().Header()
	for k, v := range resp.Header {
		responseHeader[k] = v
	}
	c.Response().WriteHeader(resp.StatusCode)

	defer resp.Body.Close()
	_, err = io.Copy(c.Response().Writer, resp.Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
