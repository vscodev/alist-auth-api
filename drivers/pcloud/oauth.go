package pcloud

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/vscodev/alist-auth-api/conf"
)

type TokenReq struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`
	Code         string `json:"code" form:"code"`
	Hostname     string `json:"hostname" form:"hostname"`
}

func GetToken(c echo.Context) error {
	r := new(TokenReq)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if r.ClientID == "" || r.ClientSecret == "" {
		r.ClientID = conf.Conf.PCloud.ClientID
		r.ClientSecret = conf.Conf.PCloud.ClientSecret
	}

	params := url.Values{
		"client_id":     []string{r.ClientID},
		"client_secret": []string{r.ClientSecret},
		"code":          []string{r.Code},
	}
	req, err := http.NewRequestWithContext(c.Request().Context(), http.MethodPost, fmt.Sprintf("https://%s/oauth2_token", r.Hostname), strings.NewReader(params.Encode()))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
