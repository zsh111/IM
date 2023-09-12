package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

// 使用viper读取设置文件
func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config mysql:", viper.Get("mysql"))
}

func InitMySQL() {
	//自定义日志打印sql语句
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{SlowThreshold: time.Second, //慢SQL阈值
		LogLevel: logger.Info, //级别
		Colorful: true,        //彩色
	})

	//从config yam中读取mysql信息
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	// user := models.UserBasic{}
	// DB.Find(&user)
	// fmt.Printf("user: %v\n", user)
}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	// ping redis测试
	// ctx := context.Background()
	// pong, err := Red.Ping(ctx).Result()
	// if err != nil {
	// 	fmt.Println("init redis ...", err)
	// } else {
	// 	fmt.Println("init success...", pong)
	// }
}

const (
	PublishKey = "websocket"
)

// Publish发送消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("Publish success...", msg)
	err = Red.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅redis的消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	fmt.Println("Subscribe success...", ctx) //测试是否连接
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Subscribe success...", msg.Payload)
	return msg.Payload, err
}
