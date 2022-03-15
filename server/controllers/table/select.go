package table

import (
	"github.com/gin-gonic/gin"
	"granada1/model/table"
	"net/http"
)

func SelectAllUserID(c *gin.Context) {
	u := table.User{}
	if err := c.BindJSON(&u);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "数据格式错误",
		})
		return
	}
	u.SelectAllUserID()
	if len(u.UserID)==0{
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"userId": u.UserID,
	})
}

func SelectAddress(c *gin.Context) {
	u := table.User{}
	if err := c.BindJSON(&u);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "数据格式错误",
		})
		return
	}
	u.SelectUserByUid(u.UserID)
	if u.Address.AddressID==""{
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"address": u.Address,
	})
}
