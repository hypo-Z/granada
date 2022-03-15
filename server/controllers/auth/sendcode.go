package auth

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goworkeryyt/go-core/global"
	"go.uber.org/zap"
	"granada/assets"
	"granada/model/redis"
	"granada/pkg"
	"granada/utils"
	"net/http"
	"time"
)

var s = redis.Store{
	Expiration: time.Minute * 5,
	PreKey:     "@",
	Context:    context.Background(),
}

// SendCode 发送验证码
func SendCode(c *gin.Context) {
	// 保存验证码错误
	acc := c.Param("account")
	code := utils.GenerateRandNum(6)
	if err := s.Set(acc, code); err != nil {
		global.LOG.Error("保存验证码错误", zap.Error(err))
		return
	}
	key := s.PreKey + acc
	fmt.Println("code:", s.Get(key, true))

	//发送验证码
	if err := SendEmail(acc, code); err != nil {
		global.LOG.Error("发送验证码错误", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "成功",
	})
}

func SendEmail(email, code string) (err error) {
	smtpOpt := global.CONFIG.Email
	Client := pkg.NewSMTPClient(smtpOpt.From, smtpOpt.Secret, smtpOpt.Host, smtpOpt.Port)

	body := assets.NewEmailBody(code, "5", "1")
	Client.Subject = code + "是你的验证码"
	Client.Body = body
	Client.From = smtpOpt.From
	Client.To = []string{email}
	if err := Client.SendHTML(); err != nil {
		return err
	} else {
		//TODO：测试
		fmt.Println(Client.To[0] + "-" + code)
	}
	return err
}
