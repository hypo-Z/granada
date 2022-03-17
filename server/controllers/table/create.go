package table

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"granada/model/table"
	"net/http"
)

/*AddInfo
{
	"user_id":""
	"phone":"",
	"summary":"",
	"gender":"",
	"age":"",
	"birthday":""
}*/
func AddInfo(c *gin.Context) {
	u := table.User{}
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "数据格式错误",
		})
		return
	}
	u.UpdateUserInfo()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "添加信息成功",
	})
	fmt.Printf("after addinfo user:%+v\n", u)
}

/*AddAddress
地址初始化 json格式：
{
"user_id":"",
"Address":{
    "Street" :"" ,
    "Unit"  :"",
    "City" :"" ,
    "State"  :"" ,
    "Zipcode" :""
	}
}*/
func AddAddress(c *gin.Context) {
	u := table.User{}
	err := c.BindJSON(&u)
	if u.UserID == "" {
		c.String(400, "id is null")
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "数据格式错误",
		})
		return
	}
	u.UpdateAddress()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "添加地址成功",
	})
	fmt.Printf("after add address user:%+v\n", &u)
}

/*CreateRelation 关系初始化 json格式：
{
	"UserID":"",
	"Relation":
	[{
		"Type":"",            // 关系类型
		"users":[{}]         // 成员
	}]
}*/
func CreateRelation(c *gin.Context) {
	u := table.User{}
	err := c.BindJSON(&u)
	fmt.Printf("before relation %+v",u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "数据格式错误",
		})
		return
	}
	u.CreateRelation()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "建立关系成功",
	})
	fmt.Printf("after create relation user:%+v\n", u)
}

/*CreateCommunity
社群初始化

{
	"user_id":"",
	"communities":
	[{
		"name":"",
	}]

}*/
func CreateCommunity(c *gin.Context) {
	u := table.User{}
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "数据格式错误",
		})
		return
	}
	u.CreateCommunity()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "创建社群成功",
	})
	fmt.Printf("after create community user:%+v\n", u)
}

/*AddCommunity
加入社群
{
	"user_id":"",
	"community":
	{
		"community_id":"",
	}

}*/
func AddCommunity(c *gin.Context) {
	u := table.User{}
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "数据格式错误",
		})
		return
	}
	u.AddCommunity(u.UserID, u.Communities[0].ID)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "创建社群成功",
	})
}

/*AddHeadImage
头像初始化
{
	"user_id":"",
	"head_images":
	[{
		"type":"",
		"size":""
	}]
}
*/
func AddHeadImage(c *gin.Context) {
	u := table.User{}
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "数据格式错误",
		})
		return
	}
	u.CreateHeadImage(u.UserID, u.HeadImages[0].Type, u.HeadImages[0].Size)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "新增头像成功",
	})
	fmt.Printf("after add headimage user:%+v\n", u)
}
