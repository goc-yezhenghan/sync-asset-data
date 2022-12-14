package initServe

import "github.com/spf13/viper"

// 初始化config，获取配置
func InitConfig() {
	viper.SetConfigFile("./configs/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("find config file failed")
		} else {
			panic("read config failed")
		}
	}
}
