package models

import (
	u "ServerApp/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name string `json:"name"`
	Priority string `json:"priority"`
	Status string `json:"status"`
	UserId uint `json:"user_id"`
}

func (task *Task) Validate() (map[string]interface{}, bool) {
	if task.Name == "" {
		return u.Message(false, "Task name is empty"), false
	}
	if task.Priority == "" {
		return u.Message(false, "Task prio is empty"), false
	}
	if task.Status == "" {
		return u.Message(false, "Task satus is empty"), false
	}
	if task.UserId <= 0 {
		return u.Message(false, "Unknown user"), false
	}
	return u.Message(true, "success"), true
}

func (task *Task)  Create() (map[string]interface{}) {
	if resp, ok := task.Validate(); !ok {
		return resp
	}
	GetDB().Create(task)
	resp := u.Message(true, "success")
	resp["task"] = task
	return resp
}

func GetTask(id uint) (*Task) {
	task := &Task{}
	err := GetDB().Table("task").Where("id = ?", id).First(task).Error
	if err != nil {
		return nil
	}
	return task
}

func GetTasksFor(user uint) ([]*Task) {
	tasks := make([]*Task, 0)
	err := GetDB().Table("tasks").Where("user_id = ?", user).Find(&tasks).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tasks
}

func GetTaskByFilterFor(user uint, status string, prio string) ([]*Task) {
	tasks := make([]*Task, 0)
	err := GetDB().Table("tasks").Where("user_id = ?", user).Where("status = ?", status).Where("priority = ?", prio).Find(&tasks).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tasks
}