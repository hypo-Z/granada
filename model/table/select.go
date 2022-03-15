package table

import "github.com/goworkeryyt/go-core/global"

// SelectAllUserID 查找所有用户id
func (u *User)SelectAllUserID()  {
	var U []string
	global.DB.Find(u).Pluck("user_id",U)
}

// SelectUserByUid 查询用户
func (u *User)SelectUserByUid(uid string) {
	global.DB.First(u).Where("UserID=?",uid)
}

// SelectUserByPhone 通过手机号查询用户
func (u *User)SelectUserByPhone(phone string) {
	global.DB.Find(u).Where("Phone=?",phone)
}

// SelectUserByEmail 通过邮箱号查询用户
func (u *User)SelectUserByEmail(email string) {
	global.DB.Find(u).Where("Email=?",email)
}

// SelectUserByAddress 通过地址查询用户
func (u *User)SelectUserByAddress() (err error) {
	if err = global.DB.Association("Address").Find(&u.Address);err != nil {
		return err
	}
	return nil
}

// SelectAllCommunity 查询所有的社群号
func (c *Community)SelectAllCommunity() {
	var C []string
	global.DB.Find(c).Pluck("community_id",C)
}

// SelectCommunityByName  通过名字查找社群
func (c *Community)SelectCommunityByName(n string) {
	global.DB.First(c).Where("name=?",n)
}

// SelectCommunityByNumber   通过群号查找社群
func (c *Community)SelectCommunityByNumber(n string) {
	global.DB.First(c).Where("number=?",n)
}