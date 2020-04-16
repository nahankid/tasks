package main

import (
	"fmt"
	"log"
	"net/http"
	"tasks/views"
)

func main() {
	http.HandleFunc("/", views.ShowAllTasksFunc)
	http.HandleFunc("/add/", views.AddTaskFunc)
	fmt.Println("running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
