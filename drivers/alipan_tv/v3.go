package alipan_tv

import (
	"crypto/md5"
	"encoding/hex"
	"slices"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/vscodev/alist-auth-api/pkg/hashset"
	"github.com/vscodev/alist-auth-api/pkg/secrets"
)

func generateDeviceParams() map[string]string {
	deviceID, _ := secrets.TokenHex(16)
	return map[string]string{
		"akv": "2.8.1496",                               // 版本号
		"apv": "1.3.9",                                  // 内部版本号
		"b":   "XiaoMi",                                 // 手机品牌
		"d":   deviceID,                                 // 设备ID，可随机生成
		"m":   "23046RP50C",                             // 手机型号
		"n":   "23046RP50C",                             // 手机型号
		"t":   strconv.FormatInt(time.Now().Unix(), 10), // 时间戳
	}
}

func calculateKey(deviceParams map[string]string) []byte {
	keys := make([]string, 0, len(deviceParams))
	for k := range deviceParams {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	var uniqueChars []rune
	charSet := hashset.New[rune]()
	for _, k := range keys {
		if k != "t" {
			for _, r := range deviceParams[k] {
				if !charSet.Contains(r) {
					uniqueChars = append(uniqueChars, r)
					charSet.Add(r)
				}
			}
		}
	}

	var numericModifier int64
	if t := deviceParams["t"]; len(t) > 7 {
		numericModifier, _ = strconv.ParseInt(t[7:], 10, 32)
	}
	offset := rune(numericModifier%127) + 1

	b := make([]byte, 0, len(uniqueChars))
	for _, r := range uniqueChars {
		newCharCode := r - offset
		if newCharCode < 0 {
			newCharCode = -newCharCode
		}

		// prevent control characters
		if newCharCode < 33 {
			newCharCode += 33
		}

		buf := make([]byte, 4)
		n := utf8.EncodeRune(buf, newCharCode)
		b = append(b, buf[:n]...)
	}

	src := md5.Sum(b)
	dst := make([]byte, 32)
	hex.Encode(dst, src[:])
	return dst
}
