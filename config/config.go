package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ReadConfig 获取配置文件
func ReadConfig(configName, configType, configPath string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(configName) // 设置文件名称（无后缀）
	v.SetConfigType(configType) // 设置后缀名 {"1.6以后的版本可以不设置该后缀"}
	v.AddConfigPath(configPath) // 设置文件所在路径
	v.Set("verbose", true)      // 设置默认参数

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf(" Config file not found; ignore error if desired")
		} else {
			log.Fatalf("Config file was found but another error was produced")
		}
	}
	// 监控配置和重新获取配置
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Fatalf("Config file changed: %+v", e.Name)
	})
	return v
}
