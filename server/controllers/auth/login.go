package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"granada1/model/auth"
	"granada1/model/table"
	"net/http"
)

// Login 登陆初始化
func Login(c *gin.Context) {
	d := auth.Data{}
	u := table.User{}
	err := c.BindJSON(&d)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "数据格式错误",
		})
		return
	}
	u.SelectUserByEmail(d.Email)
	fmt.Println("password",u.Password)
	if d.Password == u.Password {
		// 设置cookie
		c.SetCookie("auth_cookie", "123", 600, "/",
			"localhost", false, true)
		// 返回信息
		c.String(200, "login success!")
	} else {
		// 返回信息
		c.String(400, "password error！")
	}
}
