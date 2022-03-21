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

// CreateCommunity 创建社群
func (u *User) CreateCommunity() {
	nu := User{}
	nu.SelectUserByUid(u.UserID)
	fmt.Printf("beffer user %+v\n", nu)
	c := Community{}
	c.ID = utils.GetID()
	c.Name = u.Communities[0].Name
	c.CreatedBy = nu
	c.MemberSize = 1
	c.Users = append([]User{}, nu)
	nu.Communities = append([]Community{}, c)
	global.DB.Updates(nu)
}

// AddCommunity 加入社群号
func (u *User) AddCommunity(uid, cid string) {
	nu := User{}
	nu.SelectUserByUid(uid)
	c := Community{}
	switch cid {
	case "number":
		c.SelectCommunityByNumber(cid)
	case "name":
		c.SelectCommunityByName(cid)
	}
	c.MemberSize++
	c.Users = append([]User{}, nu)
	global.DB.Create(&c.Users)
	u.Communities = append([]Community{}, c)
	global.DB.Create(&u.Communities)
}

// CreateRelation 创建关系
func (u *User) CreateRelation() {
	nu := User{}
	nu.SelectUserByUid(u.Relations[0].Users[0].UserID)
	r := Relation{}
	r.Users = append([]User{}, nu)
	r.ID = utils.GetID()
	r.Type = u.Relations[0].Type
	u.SelectUserByUid(u.UserID)
	u.Relations = append([]Relation{}, r)
	global.DB.Updates(u)
}

// CreateHeadImage 添加头像
func (u *User) CreateHeadImage() {
	nu := User{}
	nu.SelectUserByUid(u.UserID)
	h := HeadImage{}
	h.HeadImageID = utils.GetID()
	h.Type = u.HeadImages[0].Type
	h.Size = u.HeadImages[0].Size
	nu.HeadImages = append([]HeadImage{}, h)
	global.DB.Updates(nu)
}
