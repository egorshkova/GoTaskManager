package main

import (
	"ClientApplication/Lib"
)

func main() {
	user := &Lib.Account{Email: "Elena@mail.com", Password: "abcdefabc"}
	//user := &Lib.Account{Email: "1234567@mail.com", Password: "test12345"}
	//user := &Lib.Account{Email: "annaAnt@gmail.com", Password: "abcabcabc"}
	//task1 := &Lib.Task{Name: "Test server", Priority: "High", Status: "New"}
	//task2 := &Lib.Task{Name: "Test client", Priority: "High", Status: "In Progress"}
	//task := &Lib.Task{Name: "AnnaTask", Priority: "Medium", Status: "New"}
	//task := &Lib.Task{Name: "TaskOctober", Priority: "High", Status: "Done"}
	//Lib.CreateNewUser(user)
	Lib.LoginUser(user)
	//Lib.CreateNewTask(task1)
	//Lib.CreateNewTask(task2)
	Lib.GetMeTasks()
    Lib.GetMeTasksByFilter("New", "High")
	Lib.GetMeTasksByFilter("In Progress", "High")
}
