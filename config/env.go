package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Env struct {
	Mode      string `yaml:"mode"`
	SecretKey string `yaml:"secret_key"`
	Server    Server `yaml:"server"`
	Mysql     Mysql  `yaml:"mysql"`
	Redis     Redis  `yaml:"redis"`
	Mongo     Mongo  `yaml:"mongo"`
}

type Server struct {
	Address        string `yaml:"address"`        // 服务地址
	Port           string `yaml:"port"`           // 端口
	ReadTimeout    int    `yaml:"readTimeout"`    // 读超时时间
	WriteTimeout   int    `yaml:"writeTimeout"`   // 写超时时间
	MaxHeaderBytes int    `yaml:"maxHeaderBytes"` // 最大头部字节数
}

type Mysql struct {
	Address  string `yaml:"address"`  // 服务地址
	Host     string `yaml:"host"`     // MySQL 地址
	Port     string `yaml:"port"`     // 端口
	User     string `yaml:"user"`     // MySQL 用户名
	Password string `yaml:"password"` // MySQL 密码
	Database string `yaml:"database"` // 数据库名
	Charset  string `yaml:"charset"`  // 字符集
	TimeOut  int    `yaml:"timeOut"`  // 超时时间
	Pool     int    `yaml:"maxConn"`  // 最大连接数
	MaxIdle  int    `yaml:"maxIdle"`  // 最大空闲连接数
}

type Mongo struct {
	Host     string `yaml:"host"`     // MongoDB 地址
	Port     string `yaml:"port"`     // 端口
	User     string `yaml:"user"`     // MongoDB 用户名
	Password string `yaml:"password"` // MongoDB 密码
	Database string `yaml:"database"` // 数据库名
}

type Redis struct {
	Host              string `yaml:"host"`              // redis 地址
	Port              int    `yaml:"port"`              // 端口
	Password          string `yaml:"password"`          // redis 密码
	Database          int    `yaml:"database"`          // 数据库
	Pool              int    `yaml:"pool"`              // 最大连接数
	Prefix            string `yaml:"prefix"`            // redis key 前缀
	Timeout           int    `yaml:"timeout"`           // 超时时间
	Retry             int    `yaml:"retry"`             // 重试次数
	RetryDelay        int    `yaml:"retryDelay"`        // 重试延迟
	RetryMultiplier   int    `yaml:"retryMultiplier"`   // 重试延迟倍数
	MaxIdle           int    `yaml:"maxIdle"`           // 最大空闲连接数
	IdleTimeoutMillis int    `yaml:"idleTimeoutMillis"` // 空闲超时时间
}

func NewEnv() *Env {
	viper.SetConfigFile("/app/config/dev.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	var env Env
	if err := viper.Unmarshal(&env); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(&env); err != nil {
			panic(err)
		}
		fmt.Println("配置文件修改了")

	})
	return &env
}

// 读取docker-compose.yaml生成的系统环境变量
//func _GetOsEnv() *Env {
//	//	viper获取系统环境变量
//	var env Env
//	env.Mode = os.Getenv("MODE")
//	//env.SecretKey = os.Getenv("SECRET_KEY")
//	env.SecretKey = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`
//	env.Server = newServer()
//	env.Mysql = newMysql()
//	env.Redis = newRedis()
//	env.Mongo = newMongo()
//
//	return &env
//}
//
//func newServer() Server {
//	return Server{
//		Address:        os.Getenv("SERVER_ADDRESS"),
//		Port:           os.Getenv("SERVER_PORT"),
//		ReadTimeout:    viper.GetInt("SERVER_READ_TIMEOUT"),
//		WriteTimeout:   viper.GetInt("SERVER_WRITE_TIMEOUT"),
//		MaxHeaderBytes: viper.GetInt("SERVER_MAX_HEADER_BYTES"),
//	}
//}
//
//func newMysql() Mysql {
//	timeOut, _ := strconv.Atoi(os.Getenv("MYSQL_TIMEOUT"))
//	pool, _ := strconv.Atoi(os.Getenv("MYSQL_POOL"))
//	maxIdle, _ := strconv.Atoi(os.Getenv("MYSQL_MAX_IDLE"))
//	return Mysql{
//		Host:     os.Getenv("MYSQL_HOST"),
//		Port:     os.Getenv("MYSQL_PORT"),
//		User:     os.Getenv("MYSQL_USER"),
//		Password: os.Getenv("MYSQL_PASSWORD"),
//		Database: os.Getenv("MYSQL_DATABASE"),
//		Charset:  os.Getenv("MYSQL_CHARSET"),
//		TimeOut:  timeOut,
//		Pool:     pool,
//		MaxIdle:  maxIdle,
//	}
//}
//
//func newRedis() Redis {
//	port, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
//	database, _ := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
//	pool, _ := strconv.Atoi(os.Getenv("REDIS_POOL"))
//	timeout, _ := strconv.Atoi(os.Getenv("REDIS_TIMEOUT"))
//	retry, _ := strconv.Atoi(os.Getenv("REDIS_RETRY"))
//	retryDelay, _ := strconv.Atoi(os.Getenv("REDIS_RETRY_DELAY"))
//	retryMultiplier, _ := strconv.Atoi(os.Getenv("REDIS_RETRY_MULTIPLIER"))
//	maxIdle, _ := strconv.Atoi(os.Getenv("REDIS_MAX_IDLE"))
//	idleTimeoutMillis, _ := strconv.Atoi(os.Getenv("REDIS_IDLE_TIMEOUT_MILLIS"))
//
//	return Redis{
//		Host:              os.Getenv("REDIS_HOST"),
//		Port:              port,
//		Password:          os.Getenv("REDIS_PASSWORD"),
//		Database:          database,
//		Pool:              pool,
//		Prefix:            os.Getenv("REDIS_PREFIX"),
//		Timeout:           timeout,
//		Retry:             retry,
//		RetryDelay:        retryDelay,
//		RetryMultiplier:   retryMultiplier,
//		MaxIdle:           maxIdle,
//		IdleTimeoutMillis: idleTimeoutMillis,
//	}
//}
//
//func newMongo() Mongo {
//	return Mongo{
//		Host:     os.Getenv("MONGO_HOST"),
//		Port:     os.Getenv("MONGO_PORT"),
//		User:     os.Getenv("MONGO_USER"),
//		Password: os.Getenv("MONGO_PASSWORD"),
//		Database: os.Getenv("MONGO_DATABASE"),
//	}
//}
