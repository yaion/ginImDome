package conf

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

var (
	// Conf 配置
	Conf *Config
)

// AppConfig 代表应用级别的配置
type AppConfig struct {
	Name  string `yaml:"name"`
	Port  int    `yaml:"port"`
	Host  string `yaml:"host"`
	Debug string `yaml:"debug"`
}

// LogConfig 代表日志相关的配置
type LogConfig struct {
	LogLevel      string `yaml:"log_level"`
	LogFile       string `yaml:"log_file"`
	LogFormat     string `yaml:"log_format"`
	LogDateFormat string `yaml:"log_date_format"`
	LogMaxSize    int    `yaml:"log_max_size"`
	LogMaxBackups int    `yaml:"log_max_backups"`
	LogMaxAge     int    `yaml:"log_max_age"`
	LogCompress   bool   `yaml:"log_compress"`
	LogLocaltime  bool   `yaml:"log_localtime"`
}

// MySQLConfig 代表MySQL数据库的配置
type MySQLConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	DB              string `yaml:"db"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int64  `yaml:"conn_max_lifetime"`  // 转换为时间Duration类型
	ConnMaxIdleTime int64  `yaml:"conn_max_idle_time"` // 同上
}

// RedisConfig 代表Redis数据库的配置
type RedisConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Password        string `yaml:"password"`
	DB              int    `yaml:"db"`
	MaxIdle         int    `yaml:"max_idle"`
	MaxActive       int    `yaml:"max_active"`
	IdleTimeout     int64  `yaml:"idle_timeout"`      // 转换为时间Duration类型
	MaxConnLifetime int64  `yaml:"max_conn_lifetime"` // 注意这里的重命名以避免冲突且转换为时间Duration类型
}

// Config 整合所有配置
type Config struct {
	App   AppConfig   `yaml:"app"`
	Log   LogConfig   `yaml:"log"`
	MySQL MySQLConfig `yaml:"mysql"`
	Redis RedisConfig `yaml:"redis"`
}

func Init() {
	Conf = new(Config)
	file, err := os.Open("conf.yaml")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	confData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(confData, &Conf)
	if err != nil {
		panic(err)
	}
}
