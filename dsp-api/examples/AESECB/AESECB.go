package main

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"errors"
	"fmt"
)

// AES-256-ECB 加密后返回 HEX 字符串
func AesECBEncryptHexJUJIA(plaintext string, keyHex string) (string, error) {

	// 1. 解析 HEX 密钥
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return "", err
	}
	if len(key) != 32 {
		return "", errors.New("key must be 32 bytes (AES-256)")
	}

	// 2. 创建 AES block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()

	// 3. PKCS7 填充
	data := pkcs7PaddingJUJIA([]byte(plaintext), blockSize)

	// 4. ECB 加密
	encrypted := make([]byte, len(data))
	for i := 0; i < len(data); i += blockSize {
		block.Encrypt(encrypted[i:i+blockSize], data[i:i+blockSize])
	}

	// 5. HEX 编码
	return hex.EncodeToString(encrypted), nil
}

// PKCS7 Padding
func pkcs7PaddingJUJIA(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	pad := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, pad...)
}

func AesECBDecryptHex(cipherHex string, keyHex string) (string, error) {

	// 1. 解析 key
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return "", err
	}

	if len(key) != 32 {
		return "", errors.New("key must be 32 bytes (AES-256)")
	}

	// 2. 解析密文
	cipherText, err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()

	if len(cipherText)%blockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of block size")
	}

	// 3. ECB 解密
	decrypted := make([]byte, len(cipherText))

	for i := 0; i < len(cipherText); i += blockSize {
		block.Decrypt(decrypted[i:i+blockSize], cipherText[i:i+blockSize])
	}

	// 4. 去 PKCS7 padding
	decrypted, err = pkcs7UnPadding(decrypted)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

// PKCS7 去填充
func pkcs7UnPadding(data []byte) ([]byte, error) {

	length := len(data)
	if length == 0 {
		return nil, errors.New("invalid padding size")
	}

	padding := int(data[length-1])

	if padding > length {
		return nil, errors.New("invalid padding")
	}

	return data[:length-padding], nil
}

//func main() {
//	encryptHex, err := Encrypt.AesECBEncryptHex("100", "b0b133a49843e5946b1979e685370d7a8bafbd0d1a102dcaf9ed1352dd72d280")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("加密结果:", encryptHex)
//
//}

func main() {
	//encryptHex, err := AesECBEncryptHex("10000", "c65b05b1ba9ae8570d999540e71b55bf9b242b761d6fa953ccbff299da267600")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("加密结果:", encryptHex)

	decryptHex, err := AesECBDecryptHex("80687d85dd47fc5928a7545a6596c24a", "b0b133a49843e5946b1979e685370d7a8bafbd0d1a102dcaf9ed1352dd72d280")
	if err != nil {
		panic(err)
	}

	fmt.Println("解密结果：", decryptHex)
}
