package Lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"strings"
)

type Account struct {
	Email string
	Password string
	Token string
}
type Task struct {
	Name string
	Priority string
	Status string
	UserId uint
}

var currentJWTtoken = ""

var GetMeTasksByFilter = func(status string, priority string) {
	url := "http://127.0.0.1:8000/api/me/taskByFilter/"
	url += "?status=" + url2.QueryEscape(status) + "&priority=" + url2.QueryEscape(priority)
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer "+string(currentJWTtoken))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("ReadAll has been failed, %s", err)
	} else {
		str := strings.Split(string(body), "},")
		log.Printf("Filtered tasks")
		log.Printf("%-20s %-8s %-8s", "Task:", "Prio:", "Status:")
		for i := 0; i < len(str); i++ {
			if strings.Contains(str[i], "priority") {
				strTask := strings.Split(str[i], ",\"")
				var currentTask Task
				for j := 0; j < len(strTask); j++ {
					if strings.Contains(strTask[j], "name") {
						taskName := strings.Split(strTask[j], "\"")
						currentTask.Name = taskName[2]
					} else if strings.Contains(strTask[j], "priority") {
						taskPrio := strings.Split(strTask[j], "\"")
						currentTask.Priority = taskPrio[2]
					} else if strings.Contains(strTask[j], "status") {
						taskStatus := strings.Split(strTask[j], "\"")
						currentTask.Status = taskStatus[2]
						break
					}
				}
				log.Printf("%-20s %-8s %-8s", currentTask.Name, currentTask.Priority, currentTask.Status)
			}
		}
	}
}

var GetMeTasks = func() {
	url := "http://127.0.0.1:8000/api/me/tasks"
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer " + string(currentJWTtoken))
	request.Header.Set("Content-Type", "application/json")
	client :=&http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("ReadAll has been failed, %s", err)
	} else {
		str := strings.Split(string(body), "},")
		log.Printf("All my tasks")
		log.Printf("%-20s %-8s %-8s", "Task:", "Prio:", "Status:")
		for i := 0; i < len(str); i++ {
			if strings.Contains(str[i], "priority") {
				strTask := strings.Split(str[i], ",\"")
				//log.Println(strTask)
				var currentTask Task
				for j := 0; j < len(strTask); j++ {
					if strings.Contains(strTask[j], "name") {
						taskName := strings.Split(strTask[j], "\"")
						currentTask.Name = taskName[2]
					} else if strings.Contains(strTask[j], "priority") {
						taskPrio := strings.Split(strTask[j], "\"")
						currentTask.Priority = taskPrio[2]
					} else if strings.Contains(strTask[j], "status") {
						taskStatus := strings.Split(strTask[j], "\"")
						currentTask.Status = taskStatus[2]
						break
					}
				}
				log.Printf("%-20s %-8s %-8s", currentTask.Name, currentTask.Priority, currentTask.Status)
			}
		}
	}
}

var CreateNewTask = func(task *Task){
	url := "http://127.0.0.1:8000/api/task/new"
	requestBody, err := json.Marshal(task)

	if err != nil {
		fmt.Printf("Can't parse account info, %s", err)
	} else {
		request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		request.Header.Set("Authorization", "Bearer " + string(currentJWTtoken))
		request.Header.Set("Content-Type", "application/json")
		client :=&http.Client{}
		response, err := client.Do(request)

		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		}
}

var LoginUser = func(account *Account) {
	requestBody, err := json.Marshal(account)

	if err == nil {
		resp, err := http.Post("http://127.0.0.1:8000/api/user/login", "application/json", bytes.NewBuffer(requestBody))

		if err != nil {
			fmt.Printf("Http post has been failed, %s", err)
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("ReadAll has been failed, %s", err)
			} else {
				str := strings.Split(string(body), ",")
				for i := 0; i < len(str); i++ {
					if strings.Contains(str[i], "token") {
						strWithToken := strings.Split(str[i], "\"")
						currentJWTtoken = strWithToken[3]
					}
				}
			}
		}

	} else {
		fmt.Printf("Can't parse account info, %s", err)
	}
}

var CreateNewUser = func(account *Account) {
	requestBody, err := json.Marshal(account)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post("http://127.0.0.1:8000/api/user/new", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
