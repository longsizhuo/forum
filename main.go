// main.go
package main

import (
	"fmt"
	"github.com/longsizhuo/forum/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
)

func main() {
	InitConfig()
	// Init SQL
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 0,            // Slow SQL threshold
			LogLevel:      logger.Error, // Log level
			Colorful:      true,
		},
	)
	dsn := viper.GetString("mysql.dns")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(
		&models.AuthRole{},
		&models.User{},
		&models.Admin{},
		&models.Section{},
		&models.Topic{},
		&models.Reply{},
	)
	if err != nil {
		panic(err)
	}

}

var BasePath = "./"

func InitConfig() {
	if basePath := os.Getenv("BASE_PATH"); basePath != "" {
		BasePath = basePath
	}
	configPath := filepath.Join(BasePath, "Config", "app.yml")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app", viper.Get("app"))
	fmt.Println("config mysql", viper.Get("mysql"))
}
