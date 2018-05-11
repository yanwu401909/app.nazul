package config

import (
	"time"
)

type Config struct {
	AppName       string        `yaml:"appName"`
	AppVersion    string        `yaml:"appVersion"`
	AppInstanceId uint32        `yaml:"appInstanceId"`
	Host          string        `yaml:"host"`
	Port          uint32        `yaml:"port"`
	Timeout       time.Duration `yaml:"timeout"`
	Start         time.Time
	Database      struct {
		Host    string `yaml:"host"`
		Port    uint32 `yaml:"port"`
		Name    string `yaml:"name"`
		Charset string `yaml:"charset"`
		User    string `yaml:"username"`
		Pass    string `yaml:"password"`
	}
	Elasticsearch struct {
		Host    string `yaml:"host"`
		Port    uint32 `yaml:"port"`
		EsIndex string `yaml:"esIndex"`
		EsType  string `yaml:"esType"`
	}
}
