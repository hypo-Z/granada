package table

import (
	"fmt"
	"github.com/goworkeryyt/go-core/global"
	"granada/utils"
)

// CreateUser 注册初始化
func (u *User) CreateUser() {
	u.UserID = utils.GetID()
	global.DB.Create(u)
}

// CreateUserInfo 添加基本信息
func (u *User) CreateUserInfo() {
	nu:=&User{}
	nu.SelectUserByUid(u.UserID)
	if nu.UserID=="" {
		fmt.Println("uid is null")
		return
	}
	fmt.Printf("beffer addinfo user %+v\n",nu)
	nu.Phone=u.Phone
	nu.Age=u.Age
	nu.Birthday=u.Birthday
	nu.Summary=u.Summary
	nu.Gender=u.Gender
	global.DB.Updates(nu)
}

// CreateAddress 添加地址
func (u *User) CreateAddress() {
	u.Address.ID = utils.GetID()
	global.DB.Create(&u.Address)
}

// CreateCommunity 创建社群
func (u *User) CreateCommunity(uid ,cn string) {
	nu:=&User{}
	nu.SelectUserByUid(uid)
	c := &Community{}
	c.ID = utils.GetID()
	c.Name=cn
	c.CreatedBy = nu
	c.MemberSize = 1
	c.Users = append([]*User{}, nu)
	u.Communities = append(u.Communities, c)
	global.DB.Create(&u.Communities)
}

// AddCommunity 加入社群号
func (u *User) AddCommunity(uid, cid string) {
	u.SelectUserByUid(uid)
	c := &Community{}
	switch cid {
	case "number":
		c.SelectCommunityByNumber(cid)
	case "name":
		c.SelectCommunityByName(cid)
	}
	c.MemberSize++
	c.Users = append([]*User{}, u)
	global.DB.Create(&c.Users)
	u.Communities = append(u.Communities, c)
	global.DB.Create(&u.Communities)
}

// CreateRelation 创建关系
func (u *User) CreateRelation(uid string, t int) {
	nu:=&User{}
	nu.SelectUserByUid(uid)
	r := &Relation{}
	r.Users = append([]*User{}, nu)
	r.ID = utils.GetID()
	r.Type = t
	u.Relations = append(u.Relations, r)
	global.DB.Create(&u.Relations)
}

// CreateHeadImage 添加头像
func (u *User) CreateHeadImage(uid, t string, s int64) {
	u.SelectUserByUid(uid)
	h := HeadImage{}
	h.HeadImageID = utils.GetID()
	h.Type = t
	h.Size = s
	u.HeadImages = append(u.HeadImages, h)
	global.DB.Create(&u.HeadImages)
}
