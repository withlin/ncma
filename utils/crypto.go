package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
)

const (
	modulus = "00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
	nonce   = "0CoJUm6Qyw8W8jud"
	pubKey  = "010001"
	iv      = "0102030405060708"
	keys string = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`
)


// 创建指定长度的key
func createSecretKey(size int) string {
	// 也就是从 a~9 以及 +/ 中随机拿出指定数量的字符拼成一个 key
	rs := ""
	for i := 0; i < size; i++ {
		pos := rand.Intn(len(keys))
		rs += keys[pos : pos+1]
	}
	return rs
}



// 加密
func Encrypt(param string) (string, string, error) {

	if encTextStr, err := aesEncrypt(param, nonce); err == nil {
		secretKey :=createSecretKey(16)
		if encText, err := aesEncrypt(encTextStr,secretKey); err == nil {
			encSecKey := RSAEncrypt(secretKey, pubKey, modulus)
			return encText, encSecKey, nil
		}
		return  "","",err
	}else {
		return "","",err
	}
}

func Decrypt(decodeStr string) (string, error) {

	decTextStr, _ := aesDecrypt(decodeStr,createSecretKey(16))
	decText, _ := aesDecrypt(decTextStr, nonce)
	return decText, nil
}

func  RSAEncrypt(secKey string, pubKey string, modulus string) string {
	encSecKey := rsaEncrypt(secKey, pubKey, modulus)
	return encSecKey
}

// AES加密的具体算法为: AES-128-CBC，输出格式为 base64
// AES加密时需要指定 iv：0102030405060708
// AES加密时需要 padding
// https://github.com/darknessomi/musicbox/wiki/%E7%BD%91%E6%98%93%E4%BA%91%E9%9F%B3%E4%B9%90%E6%96%B0%E7%99%BB%E5%BD%95API%E5%88%86%E6%9E%90
func aesEncrypt(encodeStr string, secretKeyStr string) (string, error) {
	secretKey := []byte(secretKeyStr)
	encodeBytes := []byte(encodeStr)

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = pKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	cipherText := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(cipherText, encodeBytes)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func aesDecrypt(decodeStr string, secretKeyStr string) (string, error) {
	// decode base64
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		 return "", err
	}

	secretKey := []byte(secretKeyStr)
	block, _ := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	originData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(originData, decodeBytes)
	originData = pKCS5UnPadding(originData)

	var params interface{}
	json.Unmarshal(originData, &params)

	if params != nil {
		return params.(string), nil
	}
	return string(originData[:]), nil
}

func rsaEncrypt(secKey string, pubKey string, modulus string) string {
	// 倒序 key
	rKey := ""
	for i := len(secKey) - 1; i >= 0; i-- {
		rKey += secKey[i : i+1]
	}
	// 将 key 转 ascii 编码 然后转成 16 进制字符串
	hexRKey := ""
	for _, char := range []rune(rKey) {
		hexRKey += fmt.Sprintf("%x", int(char))
	}
	// 将 16进制 的 三个参数 转为10进制的 bigint
	bigRKey, _ := big.NewInt(0).SetString(hexRKey, 16)
	bigPubKey, _ := big.NewInt(0).SetString(pubKey, 16)
	bigModulus, _ := big.NewInt(0).SetString(modulus, 16)
	// 执行幂乘取模运算得到最终的bigint结果
	bigRs := bigRKey.Exp(bigRKey, bigPubKey, bigModulus)
	// 将结果转为 16进制字符串
	hexRs := fmt.Sprintf("%x", bigRs)
	// 可能存在不满256位的情况，要在前面补0补满256位
	return addRSAPadding(hexRs, modulus)
}

// 补0步骤
func addRSAPadding(encText string, modulus string) string {
	ml := len(modulus)
	for i := 0; ml > 0 && modulus[i:i+1] == "0"; i++ {
		ml--
	}
	num := ml - len(encText)
	prefix := ""
	for i := 0; i < num; i++ {
		prefix += "0"
	}
	return prefix + encText
}

func pKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize // 16, 32, 48 etc..
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddingText...)
}

func pKCS5UnPadding(originData []byte) []byte {
	length := len(originData)
	unPadding := int(originData[length-1])
	return originData[:(length - unPadding)]
}