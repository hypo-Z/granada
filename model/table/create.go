package table

import (
	"github.com/goworkeryyt/go-core/global"
	"granada/utils"
)

// CreateUser 注册初始化
func (u *User) CreateUser() {
	u.UserID = utils.GetID()
	global.DB.Create(u)
}

// CreateUserInfo 添加基本信息
func (u *User) CreateUserInfo(uid string) {
	u.SelectUserByUid(uid)
	global.DB.Create(u)
}

// CreateAddress 添加地址
func (u *User) CreateAddress(uid string) {
	u.SelectUserByUid(uid)
	u.Address.AddressID = utils.GetID()
	global.DB.Create(&u.Address)
}

// CreateCommunity 创建社群
func (u *User) CreateCommunity(uid string) {
	u.SelectUserByUid(uid)
	c := &Community{}
	c.CommunityID = utils.GetID()
	c.CreatedBy = u
	c.MemberSize = 1
	c.Users = append([]*User{}, u)
	u.Communitys = append(u.Communitys, c)
	global.DB.Create(u)
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
	u.Communitys = append(u.Communitys, c)
	global.DB.Create(u)
}

// CreateRelation 创建关系
func (u *User) CreateRelation(uid string, t int) {
	u.SelectUserByUid(uid)
	r := &Relation{}
	r.Users = append([]*User{}, u)
	r.RelationID = utils.GetID()
	r.RelationType = t
	u.Relations = append(u.Relations, r)
	global.DB.Create(u)
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
