package global

import (
	"gorm.io/gorm"
	"time"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"id"` // 主键ID
	CreatedDate time.Time      `json:"create_date"`// 创建时间
	UpdatedDate time.Time      `json:"update_date"`// 更新时间
	DeletedDate gorm.DeletedAt `gorm:"index" json:"delete_date"` // 删除时间
}
