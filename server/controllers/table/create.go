package table

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"granada/model/table"
	"net/http"
)

/*{
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
	u.CreateUserInfo(u.UserID)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "添加信息成功",
	})
	fmt.Printf("user:%+v", u)
}

/* CreateAddress 地址初始化 json格式：
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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "数据格式错误",
		})
		return
	}
	u.CreateAddress(u.UserID)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "注册成功",
	})
	fmt.Printf("user:%+v", u)
}

/* CreateRelation 关系初始化 json格式：
{
	"UserID":"",
	"Relation":
	[{
		"relationType":"",            // 关系类型
		"users":""          // 成员
	}]
}*/
func CreateRelation(c *gin.Context) {
	u := table.User{}
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "数据格式错误",
		})
		return
	}
	u.CreateRelation(u.UserID, u.Relations[0].RelationType)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "建立关系成功",
	})
	fmt.Printf("user:%+v", u)
}

/*社群初始化

{
	"user_id":"",
	"community":
	{
		"name":"",
	}

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
	u.CreateCommunity(u.UserID)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "创建社群成功",
	})
}

/*加入社群
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
	u.AddCommunity(u.UserID, u.Communitys[0].CommunityID)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "创建社群成功",
	})
}

/*头像初始化
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
	fmt.Printf("user:%+v", u)
}
