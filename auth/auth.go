package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"net/url"
)

type Config struct {
	Key    string
	Secret string
}

func New(key, secret string) *Config {
	config := &Config{
		Key:    key,
		Secret: secret,
	}

	return config
}

// getSha256 creates a sha256 hash for given []byte
func getSha256(input []byte) []byte {
	sha := sha256.New()
	sha.Write(input)
	return sha.Sum(nil)
}

// getHMacSha512 creates a hmac hash with sha512
func getHMacSha512(message, secret []byte) []byte {
	mac := hmac.New(sha512.New, secret)
	mac.Write(message)
	return mac.Sum(nil)
}

func encryptSignature(urlPath string, values url.Values, secret []byte) string {
	// See https://www.kraken.com/help/api#general-usage for more information
	shaSum := getSha256([]byte(values.Get("nonce") + values.Encode()))
	macSum := getHMacSha512(append([]byte(urlPath), shaSum...), secret)
	return base64.StdEncoding.EncodeToString(macSum)
}

func (p *Config) SetAuthValueAndCreateSignature(urlPath string, values url.Values) (string, url.Values) {
	secret, _ := base64.StdEncoding.DecodeString(p.Secret)

	signature := encryptSignature(urlPath, values, secret)
	return signature, values
}
