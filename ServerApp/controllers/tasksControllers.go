package controllers

import (
	u "ServerApp/utils"
	"ServerApp/models"
	"encoding/json"
	"net/http"
)

var CreateTask = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	task := &models.Task{}
	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	task.UserId = user
	resp := task.Create()
	u.Respond(w, resp)
}

var GetTasksFor = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.GetTasksFor(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetTaskByFilter = func (w http.ResponseWriter, r * http.Request) {
	status := r.URL.Query().Get("status")
	priority := r.URL.Query().Get("priority")
	id := r.Context().Value("user").(uint)
	data := models.GetTaskByFilterFor(id, status, priority)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}