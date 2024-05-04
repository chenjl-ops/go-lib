package mysql_gorm

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"slices"
)

var REQUIREMENTS = []string{"=", "<>", "LIKE", "IN"}

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
		errors.Errorf("requirement %s currently not supported, please use in ['=', '<>', 'LIKE', 'IN']", requirement)
	}
	return nil
}

func (db *DbServer) Update(m any, updateData map[string]any) error {
	result := db.Engine.Model(m).Updates(updateData)
	if result.Error != nil {
		log.Error("Update data: ", result.Error)
		fmt.Println("update data error: ", result.Error)
		return result.Error
	}
	return nil
}

func (db *DbServer) Delete(data any) error {
	result := db.Engine.Delete(data)
	if result.Error != nil {
		log.Error("Delete data: ", result.Error)
		fmt.Println("delete data error: ", result.Error)
		return result.Error
	}
	return nil
}
