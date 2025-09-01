package postgresql_grom

import (
	"errors"

	"gorm.io/gorm"
)

func PaginatorHandler(paginator *Paginator) error {
	if paginator == nil {
		return errors.New("paginator is nil")
	}

	// 初始化默认分页数据
	if paginator.PageSize <= 0 {
		paginator.PageSize = 50
	}
	if paginator.CurrentPage <= 0 {
		paginator.CurrentPage = 1
	} else if paginator.CurrentPage > 0 {
		// TODO offset set
		paginator.Offset = (paginator.CurrentPage - 1) * paginator.PageSize
	}
	return nil
}

func (p *Paginator) GormPagenation() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.Offset).Limit(p.PageSize)
	}
}
