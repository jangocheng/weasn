package service

import (
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/prodbox/weasn/kernel/context"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 服务端接口
type IGuard interface {
	ShouldReturnRawResponse(*http.Request) []byte
	Dispatch([]byte) interface{}
}

// 服务端处理基类
type ServerGuard struct {
	IGuard
	context.Context
	alwaysValidate bool
}

func NewServerGuard(guard IGuard, c context.Context) *ServerGuard {
	return &ServerGuard{IGuard: guard, Context: c}
}

func (this *ServerGuard) Serve() error {
	var payload []byte
	var err error

	// 请求校验
	if this.validate() == false {
		return errors.New("请求校验失败")
	}

	if rawbytes := this.ShouldReturnRawResponse(this.Request()); rawbytes != nil {
		this.String(http.StatusOK, string(rawbytes))
		return nil
	}

	// 获取消息
	if payload, err = this.GetMessage(); err != nil {
		return err
	}

	if response := this.Dispatch(payload); response != nil {
		this.XML(http.StatusOK, this.buildResponse(response))
	}
	return nil
}

func (this *ServerGuard) buildResponse(response interface{}) interface{} {
	var encrypted []byte
	var err error

	if encrypted, err = xml.Marshal(response); err == nil || this.IsSafeMode() == false {
		return response
	}

	// 安全模式加密消息
	if encrypted, err = this.Options().Encrypter.Encrypt(encrypted); err != nil {
		return nil
	}

	timeStamp := time.Now().Unix()
	nonce := this.Query("nonce")

	signature := this.Signature(this.Options().Token, strconv.FormatInt(timeStamp, 10), this.Query("nonce"), string(encrypted))

	return &EncryptResponse{
		Encrypt:      CDATA(string(encrypted)),
		MsgSignature: CDATA(signature),
		TimeStamp:    timeStamp,
		Nonce:        CDATA(nonce),
	}
}

// Validate 请求合法校验
func (this *ServerGuard) validate() bool {
	// 非安全模式
	if !this.alwaysValidate && this.IsSafeMode() == false {
		return true
	}

	// 验证签名
	return this.Query("signature") == this.Signature(
		this.Options().Token,
		this.Query("timestamp"),
		this.Query("nonce"),
	)
}

func (this *ServerGuard) GetMessage() ([]byte, error) {
	var (
		rawBytes []byte
		err      error
	)

	// 明文模式
	if this.IsSafeMode() == false {
		if rawBytes, err = ioutil.ReadAll(this.Request().Body); err == nil {
			return rawBytes, nil
		}
		return nil, err
	}

	// 安全模式
	var content EncryptRquest
	if err = xml.NewDecoder(this.Request().Body).Decode(&content); err != nil {
		return nil, err
	}

	// 解密数据
	if rawBytes, err = this.Options().Encrypter.Decrypt(content.Encrypt, this.Query("msg_signature"), this.Query("timestamp"), this.Query("nonce")); err != nil {
		return nil, fmt.Errorf("消息解密失败, err=%v", err)
	}

	return rawBytes, nil
}

// 是否为安全模式
func (this *ServerGuard) IsSafeMode() bool {
	return len(this.Query("signature")) != 0 && this.Query("encrypt_type") == "aes"
}

// 开启强制验证
func (this *ServerGuard) ForceValidate() *ServerGuard {
	this.alwaysValidate = true
	return this
}

// 签名
func (s *ServerGuard) Signature(params ...string) string {
	sort.Strings(params)
	tmpStr := strings.Join(params, "")
	return fmt.Sprintf("%x", sha1.Sum([]byte(tmpStr)))
}
