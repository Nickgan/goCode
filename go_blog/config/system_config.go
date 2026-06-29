package config

import (
	"fmt"
	"goCode/go_blog/models"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SysConfig SystemConfig

// DB 数据库链接
var DB *gorm.DB

type ServerConfig struct {
	Port string `yaml:"port"` //服务启动端口
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type JwtConfig struct {
	Secret string `yaml:"secret"`
	Expire string `yaml:"expire"`
}

type SystemConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Jwt      JwtConfig      `yaml:"jwt"`
}

func init() {
	// 1. 初始化 Viper
	viper.SetConfigName("config")           // 配置文件名（不含扩展名）
	viper.SetConfigType("yaml")             // 配置文件类型
	viper.AddConfigPath("./go_blog/config") // 优先查找 config 目录
	viper.AddConfigPath(".")                // 备选：当前目录

	// 2. 设置默认值（可选）
	viper.SetDefault("server.port", 8080)

	// 3. 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 绑定到结构体
	if err := viper.Unmarshal(&SysConfig); err != nil {
		log.Fatalf("绑定配置到结构体失败: %v", err)
	}

	fmt.Println("读取系统配置文件成功: ", SysConfig)
}

// InitDatabase 初始化数据库连接
func InitDatabase() {

	var err error
	dbConfig := SysConfig.Database
	fmt.Println("初始化数据库连接：", dbConfig)

	// 构建MySQL连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Dbname)

	// 连接MySQL数据库
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to MySQL database:", err)
	}

	// 自动迁移数据库表结构
	err = DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("数据库链接并且自动迁移表结构成功")
}

// GetDB 获取数据库连接实例
func GetDB() *gorm.DB {
	return DB
}
