// author: wsfuyibing <websearch@163.com>
// date: 2021-08-16

package es

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var Config = &Configuration{}

const (
	DefaultAddress              = "http://127.0.0.1:9200"
	DefaultMaxConcurrency int64 = 128
)

// ES配置.
type Configuration struct {
	Address        []string `yaml:"address"`         // ES API地址列表.
	MaxConcurrency int64    `yaml:"max-concurrency"` // 队列调用ES最大并发数.
}

// 从文件中加载.
func (o *Configuration) init() {
	// 1. 解析YAML文件.
	for _, file := range []string{"tmp/es.yaml", "config/es.yaml", "../tmp/es.yaml", "../config/es.yaml"} {
		body, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}
		if yaml.Unmarshal(body, o) == nil {
			break
		}
	}

	// 2. 导入默认值.
	if Config.Address == nil || len(Config.Address) < 1 {
		Config.Address = []string{DefaultAddress}
	}
	if o.MaxConcurrency == 0 {
		o.MaxConcurrency = DefaultMaxConcurrency
	}
}
