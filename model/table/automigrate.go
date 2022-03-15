package table

import (
	"github.com/goworkeryyt/go-core/global"
	"go.uber.org/zap"
)

// AutoMigrate 自动迁移数据表
func AutoMigrate() {
	err := global.DB.AutoMigrate(&User{},&HeadImage{},&Address{})
	if err != nil {
		global.LOG.Error("数据表自动迁移失败",zap.Error(err))
		return
	}
}

