package auth

import (
	"github.com/gin-gonic/gin"
	"granada1/model/auth"
	"granada1/model/table"
	"net/http"
)

// Register 注册初始化{"name":"","password":"","email":"","code":""}
func Register(c *gin.Context) {
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

	if s.Verify(d.Email, d.Code, true) {
		u.Password = d.Password
		u.Email = d.Email
		u.Name = d.Name
		u.CreateUser()
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "注册成功",
		})
	} else {
		c.String(400, "注册失败")
	}
}
