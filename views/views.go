package views

import (
	"log"
	"net/http"
	"tasks/db"
)

//ShowAllTasksFunc is used to handle the "/" URL which is the default one
func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		context := db.GetTasks() //true when you want non deleted notes
		for _, task := range context.Tasks {
			w.Write([]byte(task.Title))
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

//AddTaskFunc is used to handle the addition of new task, "/add" URL
func AddTaskFunc(w http.ResponseWriter, r *http.Request) {
	title := "random title"
	content := "random content"
	truth := db.AddTask(title, content)
	if truth != nil {
		log.Fatal("Error adding task")
	}
	w.Write([]byte("Added task"))
}
