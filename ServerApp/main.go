package main

import (
	"ServerApp/app"
	"ServerApp/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/me/tasks", controllers.GetTasksFor).Methods("GET")
	router.HandleFunc("/api/task/new", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/me/taskByFilter/", controllers.GetTaskByFilter).Methods("GET")
	router.Use(app.JwtAuthentication) // middleware for JWT-token check

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}
	fmt.Println(port)
	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}
}