package task

import (
	"fmt"
	"github.com/chenjl-ops/go-lib/requests"
	"github.com/madflojo/tasks"
	log "github.com/sirupsen/logrus"
	"time"
)

func NewTask() *Scheduler {
	scheduler := tasks.New()
	defer scheduler.Stop()
	return &Scheduler{scheduler: scheduler}
}

func CheckHealth() error {
	//url := fmt.Sprintf("https://kyc.k8s.prod.bknws.com/health")
	url := fmt.Sprintf("http://localhost:8080/test/v1/test1")
	data := new(responseData)
	headers := map[string]string{"Content-type": "application/json"}
	err := requests.Request(url, "GET", headers, nil, data)

	log.Infof("health request: %+v", data)

	if err != nil {
		log.Error("checkHealth request error: ", err)
	}
	return nil
}

func InitTask() {
	s := NewTask()
	s.addCheck()
}

func (scheduler Scheduler) addCheck() (string, error) {
	id, err := scheduler.scheduler.Add(&tasks.Task{
		Interval: time.Duration(30) * time.Second,
		TaskFunc: CheckHealth,
	})

	if err != nil {
		log.Error("Add task error: ", err)
	}
	return id, err
}
