package conf

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

type AppConfig struct {
	Port  int
	Debug bool

	// 数据库配置
	Db struct {
		Path   string
		Log    bool
		Prefix string
	}
}

var (
	App AppConfig
	DB  *gorm.DB
)

func init() {
	readConfig()
	initDb()
}

func readConfig() {
	bytes, err := os.ReadFile("app.yml")
	if err != nil {
		panic(fmt.Sprintf("读取配置文件错误:%s", err))
	}
	err = yaml.Unmarshal(bytes, &App)
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败:%s", err))
	}
}

func initDb() {
	var logger logger2.Interface
	var nameStrategy schema.NamingStrategy
	if App.Db.Log {
		logger = logger2.Default.LogMode(logger2.Info)
	} else {
		logger = logger2.Default.LogMode(logger2.Silent)
	}
	if App.Db.Prefix == "" {
		nameStrategy = schema.NamingStrategy{
			SingularTable: true,
		}
	} else {
		nameStrategy = schema.NamingStrategy{
			TablePrefix:   App.Db.Prefix,
			SingularTable: true,
		}
	}
	db, err := gorm.Open(sqlite.Open(App.Db.Path), &gorm.Config{
		Logger:         logger,
		NamingStrategy: nameStrategy,
	})
	if err != nil {
		panic(fmt.Sprintf("连接数据库失败:%s", err))
	}
	//设置连接池
	s, _ := db.DB()
	s.SetMaxIdleConns(10)
	s.SetMaxOpenConns(100)
	s.SetConnMaxLifetime(time.Hour)
	DB = db
}
