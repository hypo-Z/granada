package table

import "github.com/goworkeryyt/go-core/global"

// SelectAllUserID 查找所有用户id
func (u *User)SelectAllUserID() []string {
	var U []string
	global.DB.Table("user").Pluck("user_id",&U)
	return U
}

// SelectUserByUid 查询用户
func (u *User)SelectUserByUid(uid string) {
	global.DB.Find(u,uid)
}

// SelectUserByName 查询用户
func (u *User)SelectUserByName(name string) {
	global.DB.Where("name=?",name).First(u)
}

// SelectUserByPhone 通过手机号查询用户
func (u *User)SelectUserByPhone(phone string) {
	global.DB.Where("phone=?",phone).First(u)
}

// SelectUserByEmail 通过邮箱号查询用户
func (u *User)SelectUserByEmail(email string) {
	global.DB.Where("email=?",email).First(u)
}

// SelectUserByAddress 通过地址查询用户
func (u *User)SelectUserByAddress(){
	global.DB.Association("Address").Find(&u.Address)
}

// SelectAllCommunity 查询所有的社群号
func (c *Community)SelectAllCommunity() {
	var C []string
	global.DB.Pluck("community_id",C).Find(c)
}

// SelectCommunityByName  通过名字查找社群
func (c *Community)SelectCommunityByName(n string) {
	global.DB.Where("name=?",n).First(c)
}

// SelectCommunityByNumber   通过群号查找社群
func (c *Community)SelectCommunityByNumber(n string) {
	global.DB.Where("number=?",n).First(c)
}

// 查找地址
func (u *User)SelectAddress()  {

}