package internal

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Item struct {
	Val     interface{}
	Created int64
	Expire  int64
}

func (item *Item) hasExpired() bool {
	return item.Expire > 0 && (time.Now().Unix()-item.Created) >= item.Expire
}

// 文件缓存
type FileCacher struct {
	rootPath string
}

func NewFileCacher(path string) *FileCacher {
	return &FileCacher{rootPath: path}
}

func (c *FileCacher) filepath(key string) string {
	m := md5.Sum([]byte(key))
	hash := hex.EncodeToString(m[:])
	return filepath.Join(c.rootPath, string(hash[0]), string(hash[1]), hash)
}

func (f *FileCacher) Set(key string, val interface{}, expire int64) error {
	filename := f.filepath(key)
	item := &Item{val, time.Now().Unix(), expire}
	data, err := EncodeGob(item)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, os.ModePerm)
}

func (f *FileCacher) read(key string) (*Item, error) {
	filename := f.filepath(key)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	item := new(Item)
	return item, DecodeGob(data, item)
}

// Get gets cached value by given key.
func (f *FileCacher) Get(key string) interface{} {
	item, err := f.read(key)
	if err != nil {
		return nil
	}

	if item.hasExpired() {
		os.Remove(f.filepath(key))
		return nil
	}
	return item.Val
}

func (f *FileCacher) Delete(key string) error {
	return os.Remove(f.filepath(key))
}

func (f *FileCacher) IsExist(key string) bool {
	_, err := os.Stat(f.filepath(key))
	return err == nil || os.IsExist(err)
}

func (f *FileCacher) Flush() error {
	return os.RemoveAll(f.rootPath)
}

func (c *FileCacher) Start() error {
	root, _ := os.Getwd()

	if !filepath.IsAbs(c.rootPath) {
		c.rootPath = filepath.Join(root, c.rootPath)
	}

	if err := os.MkdirAll(c.rootPath, os.ModePerm); err != nil {
		return err
	}
	return nil
}
