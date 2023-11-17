package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env   string
	Gin   *GinConfig
	DB    *DBConfig
	Jwt   *JwtConfig
	Redis *RedisConfig
}

type GinConfig struct {
	Host   string
	Domain string
	Port   string
	Mode   string
}

type DBConfig struct {
	Dialect  string
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Flag     string
}

type JwtConfig struct {
	Issuer               string
	Secret               string
	Refresh_token_exp    int
	Access_token_exp_sec int
}

type RedisConfig struct {
	Host           string
	Port           int
	Database       int
	Auth           string
	Max_idle       int
	Max_active     int
	Idle_timeout   int
	Notify_active  int
	Polling_active int
}

func NewConfig() *Config {
	configPath := "./"
	runPath, _ := os.Getwd()
	matchPathStatus := false
	pathArr := strings.Split(runPath, "/")
	for i := len(pathArr) - 1; i > 0; i-- {
		configPath += "../"
		if pathArr[i] == "cmd" || pathArr[i] == "test" || pathArr[i] == "migration" {
			matchPathStatus = true
			break
		}
	}
	if !matchPathStatus {
		configPath = "./"
	}
	configPath += "config"

	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		Env: viper.GetString("env"),
		Gin: &GinConfig{
			Host:   viper.GetString("server.host"),
			Domain: viper.GetString("server.domain"),
			Port:   viper.GetString("server.port"),
			Mode:   viper.GetString("server.mode"),
		},
		DB: &DBConfig{
			Dialect:  viper.GetString("db.dialect"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Database: viper.GetString("db.database"),
			Flag:     viper.GetString("db.flag"),
		},
		Jwt: &JwtConfig{
			Issuer:               viper.GetString("jwt.issuer"),
			Secret:               viper.GetString("jwt.secret"),
			Refresh_token_exp:    viper.GetInt("jwt.refresh_token_exp"),
			Access_token_exp_sec: viper.GetInt("jwt.access_token_exp_sec"),
		},
		Redis: &RedisConfig{
			Host:           viper.GetString("redis.host"),
			Port:           viper.GetInt("redis.port"),
			Database:       viper.GetInt("redis.database"),
			Auth:           viper.GetString("redis.auth"),
			Max_idle:       viper.GetInt("redis.max_idle"),
			Max_active:     viper.GetInt("redis.max_active"),
			Idle_timeout:   viper.GetInt("redis.idle_timeout"),
			Notify_active:  viper.GetInt("redis.notify_active"),
			Polling_active: viper.GetInt("redis.polling_active"),
		},
	}
}
