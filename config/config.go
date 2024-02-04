package config

import "github.com/BurntSushi/toml"

var cfg Config

type Config struct {
	Server    ServerConfig    `toml:"server"`
	Database  DatabaseConfig  `toml:"database"`
	AliyunOSS AliyunOSSConfig `toml:"aliyun_oss"`
}

type ServerConfig struct {
	Mode string `toml:"mode"`
	Port string `toml:"port"`
}

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Database string `toml:"database"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type AliyunOSSConfig struct {
	AccessKey    string `toml:"access_key"`
	SecretKey    string `toml:"secret_key"`
	Bucket       string `toml:"bucket"`
	AliyunServer string `toml:"aliyun_server"`
}

func InitConfig() {
	_, err := toml.DecodeFile("config/config.toml", &cfg)
	if err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	return cfg
}

func GetServerConfig() ServerConfig {
	return cfg.Server
}

func GetDatabaseConfig() DatabaseConfig {
	return cfg.Database
}

func GetAliyunOSSConfig() AliyunOSSConfig {
	return cfg.AliyunOSS
}
