package config

import (
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Debug            bool
	Timeout          int
	Port             int
	Mysql            Mysql
	DefaultExpire    time.Duration
	SessionRedis     RedisInfo
	CacheRedis       RedisInfo
	Sources          map[string]string
	SignFields       map[string][]string
	IgnoreSignConfig map[string]bool
	TldCacheFile     string
	VersionTimestamp string
	DNSServer        string
	HedwiServer      string
	APIBaseUrl       string
	APIKey           string
	FromUser         string
	ToUser           string
}

type Mysql struct {
	Host string
	PORT int
	DB   string
}

type RedisInfo struct {
	Address  string
	Password string
	DB       int
	PoolSize int
}

type HTTPSever struct {
	Addr string
}

type RPCInfo struct {
	Addr string
}

var DefaultConfig Config

func init() {
	InitConfigInfo()
}

func InitConfigInfo() error {
	var confFileName = "./config.toml"
	_, err := toml.DecodeFile(confFileName, &DefaultConfig)
	return err
}
