package settings

// 使用 viper 管理配置

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")  // 指定配置文件的名称 不需要带文件后缀
	viper.SetConfigType("yaml")    // 指定配置文件的类型
	viper.AddConfigPath("./conf/") // 指定查找配置文件的路径
	err = viper.ReadInConfig()     // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed %v\n", err)
		panic(fmt.Errorf("Fater error config file: %s\n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config have changed")
	})

	r := gin.Default()
	if err := r.Run(fmt.Sprintf("%d", viper.Get("port"))); err != nil {
		panic(err)
	}
	return err
}
