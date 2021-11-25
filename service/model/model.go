package model

import (
	"time"
)

type GVA_MODEL struct {
	CreatedAt time.Time `json:"-"`     //创建时间
	UpdatedAt time.Time `json:"-"`    // 更新时间
}
