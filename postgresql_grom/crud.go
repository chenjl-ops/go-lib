package postgresql_grom

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// ShowAllByPage 可分页数据查询所有数据
func (db *DbServer) ShowAllByPage(paginator *Paginator, data any) error {
	result := db.Engine.Scopes(paginator.GormPagenation()).Find(data)
	if result.Error != nil {
		log.Error("Get All Data: ", result.Error)
		fmt.Println("get all data error: ", result.Error)
		return result.Error
	}
	var total int64
	db.Engine.Model(data).Count(&total)
	paginator.Total = cast.ToInt(total)
	return nil
}
