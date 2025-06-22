package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func pkcs7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, paddingText...)
}

func pkcs7UnPadding(src []byte) []byte {
	n := len(src)
	return src[:n-int(src[n-1])]
}

func encrypt(plainText []byte, block cipher.Block, getMode func() cipher.BlockMode) (cipherText []byte, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
		}
	}()

	plainText = pkcs7Padding(plainText, block.BlockSize())
	cipherText = make([]byte, len(plainText))
	getMode().CryptBlocks(cipherText, plainText)
	return
}

func decrypt(cipherText []byte, getMode func() cipher.BlockMode) (plainText []byte, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
		}
	}()

	plainText = make([]byte, len(cipherText))
	getMode().CryptBlocks(plainText, cipherText)
	plainText = pkcs7UnPadding(plainText)
	return
}

func EncryptAESCBC(plainText, key, iv []byte) (cipherText []byte, err error) {
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}

	return encrypt(plainText, block, func() cipher.BlockMode {
		return cipher.NewCBCEncrypter(block, iv)
	})
}

func DecryptAESCBC(cipherText, key, iv []byte) (plainText []byte, err error) {
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}

	return decrypt(cipherText, func() cipher.BlockMode {
		return cipher.NewCBCDecrypter(block, iv)
	})
}

func EncryptAESECB(plainText, key []byte) (cipherText []byte, err error) {
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}

	return encrypt(plainText, block, func() cipher.BlockMode {
		return NewECBEncrypter(block)
	})
}

func DecryptAESECB(cipherText, key []byte) (plainText []byte, err error) {
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}

	return decrypt(cipherText, func() cipher.BlockMode {
		return NewECBDecrypter(block)
	})
}

func EncryptDESECB(plainText, key []byte) (cipherText []byte, err error) {
	var block cipher.Block
	block, err = des.NewCipher(key)
	if err != nil {
		return
	}

	return encrypt(plainText, block, func() cipher.BlockMode {
		return NewECBEncrypter(block)
	})
}

func DecryptDESECB(cipherText, key []byte) (plainText []byte, err error) {
	var block cipher.Block
	block, err = des.NewCipher(key)
	if err != nil {
		return
	}

	return decrypt(cipherText, func() cipher.BlockMode {
		return NewECBDecrypter(block)
	})
}
