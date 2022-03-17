package table

import (
	"fmt"
	"github.com/goworkeryyt/go-core/global"
	"granada/utils"
)

// UpdateUserByUid 更新用户
func (u *User)UpdateUserByUid(uid string) {
	global.DB.Updates(u).Where("UserID=?",uid)
}

// UpdateAddressByUid 更新地址
func (u *User)UpdateAddressByUid(addressId string) {
	global.DB.Updates(u.Address).Where("AddressID=?",addressId)
}

// UpdateCommunityByUid 更新社群
func (u *User)UpdateCommunityByUid(communityId string) {
	global.DB.Updates(u.Communities[0]).Where("CommunityID=?",communityId)
}

// UpdateRelationByUid 更新关系
func (u *User)UpdateRelationByUid(relationId string) {
	global.DB.Updates(u.Relations[0]).Where("RelationID=?",relationId)
}

// UpdateUserInfo  添加基本信息
func (u *User) UpdateUserInfo() {
	nu:=&User{}
	nu.SelectUserByUid(u.UserID)
	fmt.Printf("beffer addinfo user %+v\n",nu)
	nu.Phone=u.Phone
	nu.Age=u.Age
	nu.Birthday=u.Birthday
	nu.Summary=u.Summary
	nu.Gender=u.Gender
	global.DB.Updates(nu)
}

// UpdateAddress  添加地址
func (u *User) UpdateAddress() {
	nu:=&User{}
	nu.SelectUserByUid(u.UserID)
	nu.Address=u.Address
	nu.Address.ID = utils.GetID()
	global.DB.Updates(nu)
}