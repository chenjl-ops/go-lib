package mysql_gorm

import (
	"fmt"
	"slices"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

var REQUIREMENTS = []string{"=", "<>", "LIKE", "IN"}

// Insert 插入数据
func (db *DbServer) Insert(data any) error {
	result := db.Engine.Create(data)
	if result.Error != nil {
		log.Error("Insert data: ", result.Error)
		fmt.Println("insert data error: ", result.Error)
		return result.Error
	}
	return nil
}

// ShowAll 查询所有数据  select * from data
func (db *DbServer) ShowAll(data any) error {
	result := db.Engine.Find(data)
	if result.Error != nil {
		log.Error("Get All Data: ", result.Error)
		fmt.Println("get all data error: ", result.Error)
		return result.Error
	}
	return nil
}

// ShowSome 根据条件查询单条数据
func (db *DbServer) ShowSome(data any, requirement string, key string, value string) error {
	if slices.Contains(REQUIREMENTS, requirement) == true {
		switch requirement {
		case "LIKE":
			result := db.Engine.Where(fmt.Sprintf("%s %s ?", key, requirement), fmt.Sprintf("%%%s%%", value)).Find(data)
			if result.Error != nil {
				log.Error("Get data: ", result.Error)
				fmt.Println("get data error: ", result.Error)
				return result.Error
			}
		default:
			result := db.Engine.Where(fmt.Sprintf("%s %s ?", key, requirement), value).Find(data)
			if result.Error != nil {
				log.Error("Get data: ", result.Error)
				fmt.Println("get data error: ", result.Error)
				return result.Error
			}
		}
	} else {
		return errors.Errorf("requirement %s currently not supported, please use in ['=', '<>', 'LIKE', 'IN']", requirement)
	}
	return nil
}

// ShowSomeByMap 根据map查询数据
func (db *DbServer) ShowSomeByMap(data any, filters map[string]any) error {
	result := db.Engine.Where(filters).Find(data)
	if result.Error != nil {
		log.Error("Get data: ", result.Error)
		fmt.Println("get data error: ", result.Error)
		return result.Error
	}
	return nil
}

// Update 根据map更新数据,未验证
func (db *DbServer) Update(m any, updateData map[string]any) error {
	result := db.Engine.Model(m).Updates(updateData)
	if result.Error != nil {
		log.Error("Update data: ", result.Error)
		fmt.Println("update data error: ", result.Error)
		return result.Error
	}
	return nil
}

// Delete 删除数据
func (db *DbServer) Delete(data any) error {
	result := db.Engine.Delete(data)
	if result.Error != nil {
		log.Error("Delete data: ", result.Error)
		fmt.Println("delete data error: ", result.Error)
		return result.Error
	}
	return nil
}

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

// ShowSomeByPage 可分页根据条件查询数据
func (db *DbServer) ShowSomeByPage(paginator *Paginator, data any, requirement string, key string, value string) error {
	if slices.Contains(REQUIREMENTS, requirement) == true {
		switch requirement {
		case "LIKE":
			result := db.Engine.Scopes(paginator.GormPagenation()).Where(fmt.Sprintf("%s %s ?", key, requirement), fmt.Sprintf("%%%s%%", value)).Find(data)
			if result.Error != nil {
				log.Error("Get data: ", result.Error)
				fmt.Println("get data error: ", result.Error)
				return result.Error
			}
		default:
			result := db.Engine.Scopes(paginator.GormPagenation()).Where(fmt.Sprintf("%s %s ?", key, requirement), value).Find(data)
			if result.Error != nil {
				log.Error("Get data: ", result.Error)
				fmt.Println("get data error: ", result.Error)
				return result.Error
			}
		}
	} else {
		return errors.Errorf("requirement %s currently not supported, please use in ['=', '<>', 'LIKE', 'IN']", requirement)
	}

	var total int64
	db.Engine.Model(data).Count(&total)
	paginator.Total = cast.ToInt(total)
	return nil
}
