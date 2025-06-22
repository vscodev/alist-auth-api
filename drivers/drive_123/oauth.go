package drive_123

import (
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
	GrantType    string `json:"grant_type" form:"grant_type"`
	Code         string `json:"code,omitempty" form:"code"`
	RefreshToken string `json:"refresh_token,omitempty" form:"refresh_token"`
	RedirectURI  string `json:"redirect_uri" form:"redirect_uri"`
}

func GetToken(c echo.Context) error {
	r := new(TokenReq)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if r.ClientID == "" || r.ClientSecret == "" {
		r.ClientID = conf.Conf.Drive123.ClientID
		r.ClientSecret = conf.Conf.Drive123.ClientSecret
	}

	params := url.Values{
		"grant_type":    []string{r.GrantType},
		"client_id":     []string{r.ClientID},
		"client_secret": []string{r.ClientSecret},
	}
	if r.GrantType == "authorization_code" {
		params.Set("code", r.Code)
		params.Set("redirect_uri", r.RedirectURI)
	} else if r.GrantType == "refresh_token" {
		params.Set("refresh_token", r.RefreshToken)
	}

	req, _ := http.NewRequestWithContext(c.Request().Context(), http.MethodPost, "https://open-api.123pan.com/api/v1/oauth2/access_token", strings.NewReader(params.Encode()))
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
