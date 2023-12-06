package task

import (
	"github.com/madflojo/tasks"
)

type Scheduler struct {
	scheduler *tasks.Scheduler
}

/*
{
"code": 200,
"data": {
"MysqlHost": "127.0.0.1",
"MysqlPort": 3306,
"MysqlUserName": "root"
},
"msg": "success"
}
*/

type Data struct {
	MysqlHost     string `json:"MysqlHost"`
	MysqlPort     int    `json:"MysqlPort"`
	MysqlUserName string `json:"MysqlUserName"`
}

type responseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}
