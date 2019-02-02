package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type GlobalInfo struct {
	Version string `yaml:"Version"`
	Host    string `yaml:"Host"`
	Port    int    `yaml:"Port"`
}

type LogInfo struct {
	Path  string `yaml:"Path"`
	Level string `yaml:"Level"`
}
type DbInfo struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
}
type Config struct {
	GlobalInfo `yaml:"Global"`
	LogInfo `yaml:"Log"`
	DbInfo `yaml:"Db"`
}

var Global GlobalInfo
var Log LogInfo
var Db DbInfo

func init() {
	conf := Config{}
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("yamlFile.Get err: ", err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Println("Unmarshal: ", err)
	}
	fmt.Println("conf ", conf)
	Global = conf.GlobalInfo
	Log = conf.LogInfo
	Db = conf.DbInfo
}
