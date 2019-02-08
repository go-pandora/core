// Package setting only accomplishes loading configuration.
package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Configuration struct {
	*Database
	*Server
	*Redis
	*JWT
}

type Database struct {
	//Type     string
	DBName     string `yaml:"name"`
	DBUser     string `yaml:"user"`
	DBPassword string `yaml:"password"`
	DBHost     string `yaml:"host"`
	DBPort     string `yaml:"port"`
}

type Server struct {
	RunMode      string
	Port         string
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type Redis struct {
	RedisHost     string `yaml:"host"`
	RedisPort     string `yaml:"port"`
	RedisPassword string `yaml:"password"`
}

type JWT struct {
	SigningAlgorithm string `yaml:"signing_algorithm"`
	Secret           string
	Timeout          time.Duration
	Issuer           string
	MaxRefresh       time.Duration `yaml:"max_refresh_time"`
}

var Config Configuration

func init() {
	data, err := ioutil.ReadFile("conf/config.yaml")
	if err != nil {
		log.Panicln("failed to load config file")
	}
	if err = yaml.Unmarshal(data, &Config); err != nil {
		log.Panicln("failed to load configuration")
	} else {
		if Config.Database == nil {
			log.Panicln("failed to init database configuration")
		}
		if Config.Server == nil {
			log.Panicln("failed to init server configuration")
		}
		if Config.Redis == nil {
			log.Panicln("failed to init redis configuration")
		}
		if Config.JWT == nil {
			log.Println("failed to init jwt configuration, use default jwt config...")
			Config.JWT = &JWT{}
			Config.Timeout = time.Hour
			Config.Secret = "Hatsune Miku"
			Config.Issuer = "Fallensouls"
		} else {
			Config.Timeout *= time.Minute
		}
	}
}
