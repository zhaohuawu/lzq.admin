package config

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */
import (
	"fmt"
	"gopkg.in/ini.v1"
	"lzq-admin/pkg/utility"
	"time"
)

type AppConfig struct {
	ServerConfig `ini:"server"`
	Database     `ini:"database"`
	RedisConfig  `ini:"redis"`
	JwtConfig    `ini:"jwt"`
	LogConfig    `ini:"log"`
}

type Database struct {
	Type        string `ini:"type"`
	Host        string `ini:"host"`
	UserName    string `ini:"username"`
	Password    string `ini:"password"`
	Database    string `ini:"database"`
	MaxOpenConn int    `ini:"max_open_conn"`
	MaxIdleConn int    `ini:"max_idle_conn"`
	IsMigration bool   `ini:"is_migration"`
}

type ServerConfig struct {
	RunMode           string
	HttpPort          int
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	UseMultiTenancy   bool
	DefaultAvatar     string
	ServiceModuleCode string
}

type RedisConfig struct {
	RedisHost     string `ini:"RedisHost"`
	RedisPwd      string `ini:"RedisPwd"`
	RedisDB       int    `ini:"RedisDB"`
	RedisPoolName string `ini:"RedisPoolName"`
}

type JwtConfig struct {
	JwtIssuer     string `ini:"JwtIssuer"`
	JwtSecret     string `ini:"JwtSecret"`
	JwtExpireDate int    `ini:"JwtExpireDate"`
}

type LogConfig struct {
	ApplicationName   string `ini:"ApplicationName"`
	UseElasticSearch  bool   `ini:"UseElasticSearch"`
	ElasticSearchUrl  string `ini:"ElasticSearchUrl"`
	ElasticSearchUser string `ini:"ElasticSearchUser"`
	ElasticSearchPwd  string `ini:"ElasticSearchPwd"`
}

var Config = &AppConfig{}

// Init 初始化配置文件
func init() {
	path, err := utility.GetRootPath()
	if err != nil {
		fmt.Println("get root path err:", err)
	}
	fmt.Printf(path)
	err = ini.MapTo(Config, path+"/config/config.ini")
	if err != nil {
		fmt.Sprintf("load ini err:%v", err)
	}
}
