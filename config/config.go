package config

import (
	"io/ioutil"
	"log"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	AppName       string        `yaml:"appName"`
	AppVersion    string        `yaml:"appVersion"`
	AppInstanceId uint32        `yaml:"appInstanceId"`
	Host          string        `yaml:"host"`
	Port          uint32        `yaml:"port"`
	Timeout       time.Duration `yaml:"timeout"`
	AccessLog     bool          `yaml:"accessLog"`
	Start         time.Time
	Database      struct {
		Host         string `yaml:"host"`
		Port         uint32 `yaml:"port"`
		Name         string `yaml:"name"`
		Charset      string `yaml:"charset"`
		Autocommit   string `yaml:"autocommit"`
		User         string `yaml:"username"`
		Pass         string `yaml:"password"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
	}
	Elasticsearch struct {
		Host    string `yaml:"host"`
		Port    uint32 `yaml:"port"`
		EsIndex string `yaml:"esIndex"`
		EsType  string `yaml:"esType"`
	}
}

var CONFIG Config = Config{}

func init() {
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &CONFIG)
	if err != nil {
		log.Fatal(err)
	}
	CONFIG.Start = time.Now()
}
