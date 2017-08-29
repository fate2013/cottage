package storage

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/nicholaskh/cottage/config"
)

type tfs struct {
	baseUrl string
	config  *config.StorageConfig
}

func NewTfs(baseUrl string, config *config.StorageConfig) *tfs {
	this := new(tfs)
	this.baseUrl = baseUrl
	this.config = config
	return this
}

func (this *tfs) Store(ver, name, content string) (url string, err error) {
	resp, err := http.Post(this.baseUrl, "application/x-www-form-urlencoded", strings.NewReader(content))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	url = fmt.Sprintf("%s/%s", this.baseUrl, body)
	return
}

func (this *tfs) Fetch(url string) (content []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
