package table

import (
	"github.com/goworkeryyt/go-core/global"
)

// DeleteUserByUid 删除用户
func (u *User)DeleteUserByUid(uid string) {
	global.DB.Delete(u).Where("UserID=?",uid)
}

// DeleteAddressByUid 删除地址
func (u *User)DeleteAddressByUid(addressId string) {
	global.DB.Delete(u.Address).Where("Address=?",addressId)
}

// DeleteCommunityByUid 删除社群
func (u *User)DeleteCommunityByUid(communityId string) {
	global.DB.Delete(u.Communities[0].ID).Where("CommunityID=?",communityId)
}

// DeleteRelationByUid 删除关系
func (u *User)DeleteRelationByUid(relationId string) {
	global.DB.Delete(u.Relations[0].ID).Where("RelationID=?",relationId)
}