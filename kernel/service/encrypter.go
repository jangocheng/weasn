package service

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// 加密/解密实例
type encrypter struct {
	appId  string
	token  string
	aesKey []byte
}

func NewEncryptor(appId, token, aesKey string) *encrypter {
	en := new(encrypter)
	en.appId = appId
	en.token = token
	en.aesKey, _ = base64.StdEncoding.DecodeString(aesKey + "=")
	return en
}

func (this *encrypter) Encrypt(content []byte) ([]byte, error) {

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, int32(len(content))); err != nil {
		return nil, fmt.Errorf("Binary write err:", err)
	}

	bodyLength := buf.Bytes()
	randomBytes := []byte("abcdefghijklmnop")
	plainData := bytes.Join([][]byte{randomBytes, bodyLength, content, []byte(this.appId)}, nil)

	cipherData, err := this.aesEncrypt(plainData, this.aesKey)
	if err != nil {
		return nil, fmt.Errorf("AesEncrypt error,err=%v", err)
	}

	return []byte(base64.StdEncoding.EncodeToString(cipherData)), nil
}

func (this *encrypter) Decrypt(content, msgSignature, timestamp, nonce string) ([]byte, error) {

	if this.signature(this.token, timestamp, nonce, content) != msgSignature {
		return nil, errors.New("invalid signature")
	}

	var (
		contentBytes, decrypted []byte
		err                     error
	)

	if contentBytes, err = base64.StdEncoding.DecodeString(content); err != nil {
		return nil, err
	}

	if decrypted, err = this.aesDecrypt(contentBytes, this.aesKey, this.aesKey[:aes.BlockSize]); err != nil {
		return nil, err
	}

	if _, _, msg, appId2 := ParseFullMsg(decrypted); this.appId == appId2 {
		return msg, nil
	}

	return nil, errors.New("invalid appId")
}

func (this *encrypter) aesDecrypt(cipherData, aesKey, iv []byte) ([]byte, error) {
	if len(cipherData) < len(aesKey) {
		return nil, fmt.Errorf("the length of encrypted message too short: %d", len(cipherData))
	}
	if len(cipherData)&(len(aesKey)-1) != 0 { // or len(enc)%len(key) != 0
		return nil, fmt.Errorf("encrypted message is not a multiple of the key size(%d), the length is %d", len(aesKey), len(cipherData))
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(cipherData, cipherData)
	return PKCS7UnPadding(cipherData), nil
}

func (this *encrypter) aesEncrypt(plainData []byte, aesKey []byte) ([]byte, error) {
	k := len(aesKey)
	if len(plainData)%k != 0 {
		plainData = PKCS7Pad(plainData, k)
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cipherData := make([]byte, len(plainData))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cipherData, plainData)

	return cipherData, nil
}

func (this *encrypter) signature(token, timestamp, nonce, encrypt string) string {
	params := []string{token, timestamp, nonce, encrypt}
	sort.Strings(params)
	tmpStr := strings.Join(params, "")
	return fmt.Sprintf("%x", sha1.Sum([]byte(tmpStr)))
}

func PadLength(slice_length, blocksize int) (padlen int) {
	padlen = blocksize - slice_length%blocksize
	if padlen == 0 {
		padlen = blocksize
	}
	return padlen
}

func PKCS7Pad(message []byte, blocksize int) (padded []byte) {
	if blocksize < 1<<1 {
		panic("block size is too small (minimum is 2 bytes)")
	}
	if blocksize < 1<<8 {
		padlen := PadLength(len(message), blocksize)

		padding := bytes.Repeat([]byte{byte(padlen)}, padlen)

		padded = append(message, padding...)
		return padded
	}
	panic("unsupported block size")
}

func ParseFullMsg(fullMsg []byte) ([]byte, int, []byte, string) {
	randBytes := fullMsg[:16]

	msgLen := (int(fullMsg[16]) << 24) |
		(int(fullMsg[17]) << 16) |
		(int(fullMsg[18]) << 8) |
		int(fullMsg[19])
	return randBytes, msgLen, fullMsg[20 : 20+msgLen], string(fullMsg[20+msgLen:])
}

func PKCS7UnPadding(src []byte) (padded []byte) {
	padLen := int(src[len(src)-1])
	return src[:len(src)-padLen]
}
