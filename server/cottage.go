package server

import (
	"bytes"
	"net/http"

	"github.com/nicholaskh/cottage/config"
	"github.com/nicholaskh/cottage/record"
	"github.com/nicholaskh/cottage/storage"
	serverlib "github.com/nicholaskh/golib/server"
	log "github.com/nicholaskh/log4go"
)

type Cottage struct {
	*serverlib.HttpServer
	record  record.Record
	storage storage.Storage
	config  *config.CottageConfig
}

func NewCottageServer(conf *config.CottageConfig) *Cottage {
	this := new(Cottage)
	this.HttpServer = serverlib.NewHttpServer("cottage")
	this.initRouter()

	this.record = record.Factory(conf.Record)
	this.storage = storage.Factory(conf.Storage, this.record)
	this.config = conf
	return this
}

func (this *Cottage) initRouter() {
	this.HttpServer.RegisterHandler("/", this.test)
	this.HttpServer.RegisterHandler("/upload", this.upload)
	this.HttpServer.RegisterHandler("/search", this.search)
	this.HttpServer.RegisterHandler("/download", this.download)
	this.HttpServer.RegisterHandler("/max-version", this.maxVersion)
}

func (this *Cottage) test(w http.ResponseWriter, r *http.Request) {
	newResponse(w).setStatus(http.StatusOK).setData("Hello World").send()
}

func (this *Cottage) upload(w http.ResponseWriter, r *http.Request) {
	version := r.PostFormValue("version")
	name := r.PostFormValue("name")
	file, _, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		newResponse(w).error(RET_ERROR, err.Error()).send()
		log.Error("Receive file uploaded error: %s", err.Error())
		return
	}

	buf := make([]byte, 1024)
	_, err = file.Read(buf)
	if err != nil {
		newResponse(w).error(RET_ERROR, err.Error()).send()
		log.Error("Upload file to %s error: %s", this.config.Storage.Type, err.Error())
		return
	}
	buf = bytes.TrimRight(buf, "\x00")

	url, err := this.storage.Store(version, name, string(buf))
	if err != nil {
		newResponse(w).error(RET_ERROR, err.Error()).send()
		log.Error("Receive from storage %s error: %s", this.config.Storage.Type, err.Error())
		return
	}
	err = this.record.Record(version, name, url)
	if err != nil {
		newResponse(w).error(RET_ERROR, err.Error()).send()
		log.Error("Record package error: %s", err.Error())
		return
	}
	newResponse(w).setData("success").send()
}

func (this *Cottage) search(w http.ResponseWriter, r *http.Request) {
	word := r.PostFormValue("word")
	names, err := this.record.Search(word)
	if err != nil {
		newResponse(w).error(RET_ERROR, err.Error()).send()
		log.Error("Search from record error: %s", err.Error())
		return
	}
	newResponse(w).setData(names).send()
}

func (this *Cottage) download(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	version := r.PostFormValue("version")
	url, err := this.record.GetUrl(name, version)
	if err != nil {
		newResponse(w).error(RET_ERROR, err.Error()).send()
		log.Error("Get url from record error: %s", err.Error())
		return
	}
	content, err := this.storage.Fetch(url)
	if err != nil {
		newResponse(w).error(RET_ERROR, err.Error()).send()
		log.Error("Receive from storage %s error: %s", this.config.Storage.Type, err.Error())
		return
	}
	newResponse(w).sendRaw(content)
}

func (this *Cottage) maxVersion(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	maxVersion, err := this.record.MaxVersion(name)
	if err != nil {
		newResponse(w).error(RET_ERROR, err.Error()).send()
		log.Error("Receive from storage %s error: %s", this.config.Storage.Type, err.Error())
		return
	}
	newResponse(w).setData(maxVersion).send()
}

func (this *Cottage) Launch(listenAddr string) {
	this.HttpServer.Launch(listenAddr)
}
