package alipan_tv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetQrCode(c echo.Context) error {
	r := new(QrCodeReq)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, _ := json.Marshal(r)
	req, _ := http.NewRequestWithContext(c.Request().Context(), http.MethodPost, "https://api.extscreen.com/aliyundrive/qrcode", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.Request().UserAgent())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	defer resp.Body.Close()
	var v Response[QrCodeData]
	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if v.Code != http.StatusOK {
		return echo.NewHTTPError(http.StatusInternalServerError, v.Message)
	}

	return c.JSON(http.StatusOK, v.Data)
}

func CheckQrCodeStatus(c echo.Context) error {
	sid := c.Param("sid")
	req, _ := http.NewRequestWithContext(c.Request().Context(), http.MethodGet, fmt.Sprintf("https://openapi.alipan.com/oauth/qrcode/%s/status", sid), nil)
	req.Header.Set("User-Agent", c.Request().UserAgent())
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

func GetToken(c echo.Context) error {
	r := new(TokenReq)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, _ := json.Marshal(r)
	req, _ := http.NewRequestWithContext(c.Request().Context(), http.MethodPost, "https://api.extscreen.com/aliyundrive/v3/token", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.Request().UserAgent())

	deviceParams := generateDeviceParams()
	for k, v := range deviceParams {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	defer resp.Body.Close()
	var v Response[EncryptedTokenData]
	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if v.Code != http.StatusOK {
		return echo.NewHTTPError(http.StatusInternalServerError, v.Message)
	}

	tokenData, err := v.Data.Decrypt(calculateKey(deviceParams))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tokenData)
}
