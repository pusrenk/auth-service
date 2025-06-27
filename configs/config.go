package configs

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var (
	ConfigInstance *Config
	once           sync.Once
)

type Config struct {
	App      AppConfig      `json:"app" mapstructure:"app"`
	Database DatabaseConfig `json:"database" mapstructure:"database"`
	Redis    RedisConfig    `json:"redis" mapstructure:"redis"`
}

type AppConfig struct {
	Name    string `json:"name" mapstructure:"name"`
	Version string `json:"version" mapstructure:"version"`
	Port    int    `json:"port" mapstructure:"port"`
	Env     string `json:"env" mapstructure:"env"`
}

type DBConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	User     string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
}

type DatabaseConfig struct {
	DBConfig
	Database        string        `json:"database" mapstructure:"database"`
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time" mapstructure:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime" mapstructure:"conn_max_lifetime"`
	MaxIdleConns    int           `json:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `json:"max_open_conns" mapstructure:"max_open_conns"`
}

type RedisConfig struct {
	DBConfig
	Database int `json:"database" mapstructure:"database"`
}

// get config instance
func GetConfig() *Config {
	return ConfigInstance
}

// load config from env.json file
func loadConfig(config *Config) {
	if ConfigInstance == nil {
		panic(fmt.Errorf("config not initialized"))
	}

	viper.SetConfigName("env")
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error reading config file: %v", err))
	}

	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("Error unmarshalling config: %v", err))
	}
}

// init config
func init() {
	once.Do(func() {
		ConfigInstance = &Config{}
		loadConfig(ConfigInstance)
	})
}
