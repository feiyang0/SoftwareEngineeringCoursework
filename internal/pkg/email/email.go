package email

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/viper"
	"math/rand"
	"sync"
	"time"

	"gopkg.in/gomail.v2"
)

type EmailParam struct {
	ServerHost string // ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	ServerPort int    // ServerPort 邮箱服务器端口，如腾讯企业邮箱为465
	FromEmail  string // FromEmail　发件人邮箱地址
	FromPasswd string // 发件人邮箱密码（注意，这里是明文形式），TODO：如果设置成密文？
}

var (
	once sync.Once
)
var emailIns *EmailParam

func getEmail() *EmailParam {
	once.Do(func() {
		emailIns = &EmailParam{
			ServerHost: viper.GetString("email.ServerHost"),
			ServerPort: viper.GetInt("email.ServerPort"),
			FromEmail:  viper.GetString("email.FromEmail"),
			FromPasswd: viper.GetString("email.FromPasswd"),
		}
		//fmt.Println("email:", emailIns)
	})
	return emailIns
}
func createCaptcha() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
func SendCaptcha(email string) (string, error) {
	body := `
    <p> 验证码:%s</p>
	`
	ep := getEmail()
	m := gomail.NewMessage()

	m.SetHeader("From", ep.FromEmail) // 发件人
	m.SetHeader("To", email)
	m.SetHeader("Subject", "荔课网练验证码")
	Captcha := createCaptcha()
	m.SetBody("text/html", fmt.Sprintf(body, Captcha))

	d := gomail.NewDialer(
		ep.ServerHost,
		ep.ServerPort,
		ep.FromEmail,
		ep.FromPasswd,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return "", err
	}
	return Captcha, nil
}
