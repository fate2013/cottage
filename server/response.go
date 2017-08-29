package server

import (
	"encoding/json"
	"net/http"

	log "github.com/nicholaskh/log4go"
)

const (
	RET_SUCCESS = 0
	RET_ERROR   = 1
)

type response struct {
	w      http.ResponseWriter
	status int
	ret    int
	data   interface{}
	err    string
}

func newResponse(w http.ResponseWriter) *response {
	this := new(response)
	this.w = w
	this.status = http.StatusOK
	this.ret = RET_SUCCESS
	return this
}

func (this *response) setStatus(status int) *response {
	this.status = status
	return this
}

func (this *response) setData(data interface{}) *response {
	this.data = data
	return this
}

func (this *response) success() *response {
	this.ret = RET_SUCCESS
	return this
}

func (this *response) error(errcode int, errmsg string) *response {
	this.ret = errcode
	this.err = errmsg
	return this
}

func (this *response) send() {
	var payload map[string]interface{}
	if this.ret == RET_SUCCESS {
		payload = map[string]interface{}{"ret": this.ret, "data": this.data}
	} else {
		payload = map[string]interface{}{"ret": this.ret, "error": this.err}
	}
	body, err := json.Marshal(payload)
	if err != nil {
		log.Warn("Json marshal error: %#v", payload)
	}
	this.w.WriteHeader(this.status)
	this.w.Write(body)
}

func (this *response) sendRaw(raw []byte) {
	this.w.Write(raw)
}
