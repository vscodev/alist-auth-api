package alipan_tv

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"github.com/vscodev/alist-auth-api/pkg/crypto"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type QrCodeReq struct {
	Scopes string `json:"scopes" form:"scopes"`
	Width  int    `json:"width,omitempty" form:"width"`
	Height int    `json:"height,omitempty" form:"height"`
}

type QrCodeData struct {
	QrCodeUrl string `json:"qrCodeUrl"`
	Sid       string `json:"sid"`
}

type TokenReq struct {
	Code         string `json:"code,omitempty" form:"code"`
	RefreshToken string `json:"refresh_token,omitempty" form:"refresh_token"`
}

type EncryptedTokenData struct {
	Ciphertext string `json:"ciphertext"`
	Iv         string `json:"iv"`
}

type TokenData struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func (d EncryptedTokenData) Decrypt(key []byte) (*TokenData, error) {
	cipherText, err := base64.StdEncoding.DecodeString(d.Ciphertext)
	if err != nil {
		return nil, err
	}

	iv, err := hex.DecodeString(d.Iv)
	if err != nil {
		return nil, err
	}

	plainText, err := crypto.DecryptAESCBC(cipherText, key, iv)
	if err != nil {
		return nil, err
	}

	v := new(TokenData)
	if err = json.Unmarshal(plainText, v); err != nil {
		return nil, err
	}

	return v, nil
}
