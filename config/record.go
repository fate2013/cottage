package config

import (
	conf "github.com/nicholaskh/jsconf"
)

type RecordConfig struct {
	Type string

	Host     string
	Port     int
	Username string
	Password string
	Db       string
}

func (this *RecordConfig) LoadConfig(cf *conf.Conf) {
	this.Type = cf.String("type", "mysql")

	this.Host = cf.String("host", "127.0.0.1")
	this.Port = cf.Int("port", 3306)
	this.Username = cf.String("username", "root")
	this.Password = cf.String("password", "")
	this.Db = cf.String("db", "cottage")
}
