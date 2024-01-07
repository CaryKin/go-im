/*
*

	@author: jixiaogang
	@since: 2024/1/6
	@desc: //TODO

*
*/
package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseStruct struct {
	ID        uint           `json:"id" gorm:"primarykey"` // 主键ID
	CreatedAt time.Duration  `json:"createdAt"`            // 创建时间
	UpdatedAt time.Duration  `json:"updatedAt"`            // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
}
