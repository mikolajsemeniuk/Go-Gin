package configuration

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Config *viper.Viper
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
		message := fmt.Sprintf("\nenv environment variable is not set, setting to: %s\n", env)
		fmt.Println(Yellow + message + Reset)
	}

	Config = viper.New()
	Config.SetConfigType("yaml")
	Config.SetConfigName(env)
	Config.AddConfigPath("../configuration/")
	Config.AddConfigPath("configuration/")
	err := Config.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
