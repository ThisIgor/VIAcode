package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type elasticSearch struct {
	Address string
}

type mediaStorage struct {
	RootPath string
}

type routingConf struct {
	Type     string
	Http     string
	Https    string
	CertFile string
	KeyFile  string
}

type logging struct {
	LogPath         string
	CommonLogName   string
	DatabaseLogName string
}

type dbConfiguration struct {
	DatabaseType string
	Host         string
	Port         uint
	DatabaseName string
	Login        string
	Password     string
}

type configuration struct {
	ElasticSearch elasticSearch
	MediaStorage  mediaStorage
	Routing       routingConf
	Logging       logging
	DBConf        dbConfiguration
}

var Config configuration

func configFronCommandLine() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	pos := strings.LastIndex(os.Args[0], "/")
	return fmt.Sprintf("/usr/local/etc%v.conf", os.Args[0][pos:])
}

func Load() error {
	data, err := ioutil.ReadFile(configFronCommandLine())
	if err == nil {
		err = json.Unmarshal(data, &Config)
	}
	return err
}
