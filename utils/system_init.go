package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

//使用viper读取设置文件
func InitConfig(){
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err:= viper.ReadInConfig()
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("config app:",viper.Get("app"))
	fmt.Println("config mysql:",viper.Get("mysql"))
}

func InitMySQL(){
	//自定义日志打印sql语句
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),logger.Config{SlowThreshold: time.Second,//慢SQL阈值
		LogLevel: logger.Info,//级别
		Colorful: true,//彩色
	},)

	//从config yam中读取mysql信息
	DB, _= gorm.Open(mysql.Open(viper.GetString("mysql.dns")),&gorm.Config{Logger: newLogger})
	// user := models.UserBasic{}
	// DB.Find(&user)
	// fmt.Printf("user: %v\n", user)

}