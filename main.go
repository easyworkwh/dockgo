package main

/*
	The source code is from https://zhuanlan.zhihu.com/p/382175578
*/

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// AppConfig 应用配置结构体
type AppConfig struct {
	AppName     string
	Port        int
	Description string
}

// conf 全局配置信息实例
var conf AppConfig

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在路径
	viper.AddConfigPath("$HOME/.appname") // 多次调用AddConfigPath，可以添加多个搜索路径
	viper.AddConfigPath(".")              // 还可以在工作目录中搜索配置文件
	viper.AddConfigPath("./conf")         // 还可以在工作目录中搜索配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("read config failed: %v \r", err)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Println(err)
	}
}

// main 主函数入口，使用gin启动一个简单的http交互服务
func main() {
	r := gin.Default()
	r.Any("/", func(c *gin.Context) {
		message := strings.Join([]string{"Hello", conf.AppName, conf.Description}, " ")
		c.JSON(200, gin.H{
			"message": message,
		})
	})
	appPort := strings.Join([]string{":", strconv.Itoa(conf.Port)}, "")
	r.Run(appPort)
}
