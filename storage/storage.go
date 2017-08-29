package storage

import (
	"github.com/nicholaskh/cottage/config"
	"github.com/nicholaskh/cottage/record"
)

type Storage interface {
	Store(ver, name, content string) (url string, err error)
	Fetch(url string) (content []byte, err error)
}

func Factory(cf *config.StorageConfig, record record.Record) (storage Storage) {
	switch cf.Type {
	case "tfs":
		storage = NewTfs(cf.BaseUrl, cf)
	default:
		storage = NewTfs(cf.BaseUrl, cf)
	}

	return
}
