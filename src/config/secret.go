package config

import (
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Secret struct {
	App struct {
		Name     string `yaml:"name"`
		Url      string `yaml:"url"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		TimeZone string `yaml:"timezone"`
	} `yaml:"app"`
	Db struct {
		Source string `yaml:"source"`
		ScGorm string `yaml:"scgorm"`
	} `yaml:"db"`
	Rmq struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	} `yaml:"rmq"`
	Telco struct {
		Url string `yaml:"url"`
	} `yaml:"telco"`
	Portal struct {
		Url string `yaml:"url"`
	} `yaml:"portal"`
	Report struct {
		Url string `yaml:"url"`
	} `yaml:"report"`
	Log struct {
		Path string `yaml:"path"`
	}
}

func LoadSecret(path string) (*Secret, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return LoadSecretFromBytes(data)
}

func LoadSecretFromBytes(data []byte) (*Secret, error) {
	fang := viper.New()
	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()
	fang.SetEnvPrefix("GO")
	fang.SetConfigType("yaml")

	if err := fang.ReadConfig(bytes.NewBuffer(data)); err != nil {
		return nil, err
	}
	var creds Secret
	err := fang.Unmarshal(&creds)
	if err != nil {
		log.Fatalf("Error loading creds: %v", err)
	}
	return &creds, nil
}
