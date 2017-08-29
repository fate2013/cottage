package config

import (
	conf "github.com/nicholaskh/jsconf"
)

type StorageConfig struct {
	Type    string
	BaseUrl string
}

func (this *StorageConfig) LoadConfig(cf *conf.Conf) {
	this.Type = cf.String("type", "tfs")
	this.BaseUrl = cf.String("base_url", "")
}
