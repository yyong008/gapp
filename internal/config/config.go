// config/config.go
package config

import (
	"github.com/spf13/viper"
	"log"
)

func LoadConfig() {
	viper.SetConfigName("config")   // 配置文件名
	viper.SetConfigType("yaml")     // 配置文件类型
	viper.AddConfigPath("./config") // 配置文件所在路径

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

// 获取 JWT 配置
func GetJWTSecretKey() string {
	return viper.GetString("jwt.secret_key")
}

func GetJWTIssuer() string {
	return viper.GetString("jwt.issuer")
}

func GetJWTExpiration() int {
	return viper.GetInt("jwt.expiration")
}
