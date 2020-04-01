package internal

import (
	"bytes"
	"encoding/gob"
)

func init() {
	// 注册 map[string]string{} 类型
	gob.Register(map[string]string{})
}

func EncodeGob(item *Item) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := gob.NewEncoder(buf).Encode(item)
	return buf.Bytes(), err
}

func DecodeGob(data []byte, out *Item) error {
	buf := bytes.NewBuffer(data)
	return gob.NewDecoder(buf).Decode(&out)
}
