package record

import (
	"github.com/nicholaskh/cottage/config"
)

type Record interface {
	Record(ver, name, url string) (err error)
	Search(word string) (names []string, err error)
	MaxVersion(name string) (maxVersion string, err error)
	GetUrl(name, version string) (url string, err error)
}

func Factory(config *config.RecordConfig) Record {
	switch config.Type {
	case "mysql":
		return newMysql(config)
	default:
		return newMysql(config)
	}
}
