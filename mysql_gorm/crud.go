package mysql_gorm

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"slices"
)

var REQUIREMENTS = []string{"=", "<>", "LIKE", "IN"}

func Insert(data any) error {
	result := Engine.Create(data)
	if result.Error != nil {
		log.Error("Insert data: ", result.Error)
		fmt.Sprintln("insert data error: ", result.Error)
		return result.Error
	}
	return nil
}

// ShowAll 查询所有数据  select * from data
func ShowAll(data any) error {
	result := Engine.Find(data)
	if result.Error != nil {
		log.Error("Get All Data: ", result.Error)
		fmt.Sprintln("get all data error: ", result.Error)
		return result.Error
	}
	return nil
}

// ShowSome 根据条件查询单条数据
func ShowSome(data any, requirement string, key string, value string) error {
	if slices.Contains(REQUIREMENTS, requirement) == true {
		switch requirement {
		case "LIKE":
			result := Engine.Where(fmt.Sprintf("%s %s ?", key, requirement), fmt.Sprintf("%%%s%%", value)).Find(data)
			if result.Error != nil {
				log.Error("Get data: ", result.Error)
				fmt.Sprintln("get data error: ", result.Error)
				return result.Error
			}
		default:
			result := Engine.Where(fmt.Sprintf("%s %s ?", key, requirement), value).Find(data)
			if result.Error != nil {
				log.Error("Get data: ", result.Error)
				fmt.Sprintln("get data error: ", result.Error)
				return result.Error
			}
		}
	} else {
		errors.Errorf("requirement %s currently not supported, please use in ['=', '<>', 'LIKE', 'IN']", requirement)
	}
	return nil
}

func Update(m any, updateData map[string]any) error {
	result := Engine.Model(m).Updates(updateData)
	if result.Error != nil {
		log.Error("Update data: ", result.Error)
		fmt.Sprintln("update data error: ", result.Error)
		return result.Error
	}
	return nil
}

func Delete(data any) error {
	result := Engine.Delete(data)
	if result.Error != nil {
		log.Error("Delete data: ", result.Error)
		fmt.Sprintln("delete data error: ", result.Error)
		return result.Error
	}
	return nil
}
