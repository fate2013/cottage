package config

import (
	conf "github.com/nicholaskh/jsconf"
)

var (
	Cottage *CottageConfig
)

type CottageConfig struct {
	ListenAddr string

	Storage *StorageConfig
	Record  *RecordConfig
}

func (this *CottageConfig) LoadConfig(cf *conf.Conf) {
	this.ListenAddr = cf.String("listen_addr", ":8844")

	this.Storage = new(StorageConfig)
	section, err := cf.Section("storage")
	if err == nil {
		this.Storage.LoadConfig(section)
	}

	this.Record = new(RecordConfig)
	section, err = cf.Section("record")
	if err != nil {
		panic("Record Config not found")
	}
	this.Record.LoadConfig(section)
}
