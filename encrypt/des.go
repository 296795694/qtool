package encrypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

// EntryptDesCBC
// @Description: DES CBC模式加密
// @Auth syq
// @Date 2021-12-14 09:04:13
// @param origData
// @param key
// @param iv
// @return []byte
// @return error
func EntryptDesCBC(origData, key, iv []byte) ([]byte, error) {
	//iv即是向量，长度为8
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// DecryptDESCBC
// @Description: DES CBC模式解密
// @Auth syq
// @Date 2021-12-14 09:03:34
// @param crypted
// @param key
// @param iv
// @return []byte
// @return error
func DecryptDESCBC(crypted, key, iv []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// EntryptDesECB
// @Description: DES ECB模式加密
// @Auth syq
// @Date 2021-12-14 09:02:36
// @param data
// @param key
// @return string
func EntryptDesECB(data, key []byte) string {
	if len(key) > 8 {
		key = key[:8]
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return ""
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return ""
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return base64.StdEncoding.EncodeToString(out)
}

// DecryptDESECB
// @Description:  DES ECB模式解密
// @Auth syq
// @Date 2021-12-14 09:03:05
// @param d
// @param key
// @return string
func DecryptDESECB(d string, key []byte) string {
	data, err := base64.StdEncoding.DecodeString(d)
	if err != nil {
		return ""
	}
	if len(key) > 8 {
		key = key[:8]
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return ""
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return ""
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
