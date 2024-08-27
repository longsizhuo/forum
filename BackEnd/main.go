// main.go
package main

import (
	"fmt"
	"github.com/longsizhuo/forum/models"
	pb "github.com/longsizhuo/forum/proto"
	"github.com/longsizhuo/forum/services"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net"
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

	// 启动 gRPC 服务
	lis, err := net.Listen("tcp", viper.GetString("port"))
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	//注册服务
	pb.RegisterForumServiceServer(s, &services.Server{Db: db})
	log.Printf("server listening at %v", lis.Addr())

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

var BasePath = "./BackEnd"

func InitConfig() {
	if basePath := os.Getenv("BASE_PATH"); basePath != "" {
		BasePath = basePath
	}
	configPath := filepath.Join(BasePath, "config", "app.yml")
	fmt.Println("configPath", configPath)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic("Failed to read config file")
	}
	fmt.Println("config app", viper.Get("app"))
	fmt.Println("config mysql", viper.Get("mysql"))
}
