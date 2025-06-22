package disk_115

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/vscodev/alist-auth-api/conf"
)

type QrCodeReq struct {
	ClientID            string `json:"client_id" form:"client_id"`
	CodeChallenge       string `json:"code_challenge" form:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method" form:"code_challenge_method"`
}

type DeviceCodeTokenReq struct {
	Uid          string `json:"uid" form:"uid"`
	CodeVerifier string `json:"code_verifier" form:"code_verifier"`
}

type AuthCodeTokenReq struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`
	Code         string `json:"code" form:"code"`
	RedirectURI  string `json:"redirect_uri" form:"redirect_uri"`
}

func GetQrCode(c echo.Context) error {
	r := new(QrCodeReq)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	params := url.Values{
		"client_id":             []string{r.ClientID},
		"code_challenge":        []string{r.CodeChallenge},
		"code_challenge_method": []string{r.CodeChallengeMethod},
	}
	req, _ := http.NewRequestWithContext(c.Request().Context(), http.MethodPost, "https://passportapi.115.com/open/authDeviceCode", strings.NewReader(params.Encode()))
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

func GetTokenViaDeviceCode(c echo.Context) error {
	r := new(DeviceCodeTokenReq)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	params := url.Values{
		"uid":           []string{r.Uid},
		"code_verifier": []string{r.CodeVerifier},
	}
	req, _ := http.NewRequestWithContext(c.Request().Context(), http.MethodPost, "https://passportapi.115.com/open/deviceCodeToToken", strings.NewReader(params.Encode()))
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

func GetTokenViaAuthCode(c echo.Context) error {
	r := new(AuthCodeTokenReq)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if r.ClientID == "" || r.ClientSecret == "" {
		r.ClientID = conf.Conf.Disk115.ClientID
		r.ClientSecret = conf.Conf.Disk115.ClientSecret
	}

	params := url.Values{
		"grant_type":    []string{"authorization_code"},
		"code":          []string{r.Code},
		"client_id":     []string{r.ClientID},
		"client_secret": []string{r.ClientSecret},
	}
	req, _ := http.NewRequestWithContext(c.Request().Context(), http.MethodPost, "https://passportapi.115.com/open/authCodeToToken", strings.NewReader(params.Encode()))
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
