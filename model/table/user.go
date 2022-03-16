package table

import (
	"time"
)

// User 用户信息
type User struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
	UserID      string     `gorm:"type:varchar(20);primary_key"`     // 关联外键
	UserName    string     `gorm:"type:varchar(20);index:user_name"` // 账号名
	Phone       string     `gorm:"type:varchar(20);index:phone"`     // 手机号
	Email       string     `gorm:"type:varchar(50);index:email"`     // 邮箱
	Password    string     `gorm:"type:varchar(20);not null"`        // 密码
	Summary     string     `gorm:"type:varchar(150)"`                // 设置简介大小为255
	Gender      bool       `gorm:"type:boolean;"`                    // 性别
	Age         int        `gorm:"type:varchar(2);"`                 // 年龄
	Birthday    *time.Time  // 生日
	AddressID   int
	Address     Address      `gorm:"foreignKey:AddressID"`     // 地址
	Communities []*Community `gorm:"many2many:user_community"` // 加入的社群号
	HeadImages  []HeadImage  `gorm:"foreignKey:HeadImageID"`   // 用户头像
	Relations   []*Relation  `gorm:"many2many:user_relation"`  // 所从关系
}

// HeadImage 可实现切换头像功能
type HeadImage struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
	HeadImageID string     `gorm:"type:varchar(20);primary_key"` // 关联外键
	Type        string     `gorm:"type:varchar(20)"`             // 图片类型
	Size        int64      `gorm:"type:varchar(150)"`            // 图片大小
}

// Community 社群信息
type Community struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	ID         string     `gorm:"type:varchar(20);primary_key"`          // 关联外键
	Name       string     `gorm:"type:varchar(20);index:community_name"` // 社群名
	Number     string     `gorm:"type:varchar(20);index:number"`         // 社群号
	CreatedBy  *User      `gorm:"foreignKey:UserID"`                     // 社群创建者
	MemberSize int        `gorm:"type:integer(255);not null"`            // 社群成员数
	Users      []*User    `gorm:"many2many:user_community"`              // 成员
}

// Address 用户地址
type Address struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	ID        string     `gorm:"type:varchar(20);primary_key"` // 关联外键
	Street    string     `gorm:"type:varchar(20)"`             // 街道
	Unit      string     `gorm:"type:varchar(20)"`             // 单元（可以不存在）
	City      string     `gorm:"type:varchar(20)"`             // 城市
	State     string     `gorm:"type:varchar(20)"`             // 州/省
	Zipcode   string     `gorm:"type:varchar(20)"`             // 邮编
}

// Relation 人员关系表
type Relation struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	ID        string     `gorm:"type:varchar(20);primary_key"` // 关联外键
	Type      int        `gorm:"type:varchar(2)"`              // 关系类型
	Users     []*User    `gorm:"many2many:user_relation;"`     // 成员
}

// RelationType 人际关系类型
type RelationType struct {
	Friend    int
	Relatives int
	Classmate int
	Teachers  int
	Employ    int
	Lover     int
	Hobby     int
}
