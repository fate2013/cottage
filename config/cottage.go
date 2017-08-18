package config

import (
	conf "github.com/nicholaskh/jsconf"
)

var (
	Cottage *CottageConfig
)

type CottageConfig struct {
	ListenAddr string
}

func (this *CottageConfig) LoadConfig(cf *conf.Conf) {
	this.ListenAddr = cf.String("listen_addr", ":8844")
}
