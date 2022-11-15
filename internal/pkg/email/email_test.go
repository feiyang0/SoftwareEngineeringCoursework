package email

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestGetEmail(t *testing.T) {
	viper.SetConfigFile("../../../config/server.yaml")
	viper.ReadInConfig()
	_, err := SendCaptcha("yang17898441572@163.com")
	if err != nil {
		fmt.Println("err:", err)
	}
}
