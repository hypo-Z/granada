package table

import "github.com/goworkeryyt/go-core/global"

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
	global.DB.Updates(u.Communitys[0]).Where("CommunityID=?",communityId)
}

// UpdateRelationByUid 更新关系
func (u *User)UpdateRelationByUid(relationId string) {
	global.DB.Updates(u.Relations[0]).Where("RelationID=?",relationId)
}
