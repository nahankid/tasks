package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var database *sql.DB
var err error

func init() {
	connStr := "user=bracket dbname=tasks sslmode=disable"
	database, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", ShowAllTasksFunc)
	http.HandleFunc("/add/", AddTaskFunc)
	fmt.Println("running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//ShowAllTasksFunc is used to handle the "/" URL which is the default one
func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		context := GetTasks() //true when you want non deleted notes
		for _, task := range context.Tasks {
			w.Write([]byte(task.Title))
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// GetTasks returns tasks from the database
func GetTasks() Context {
	var tasks []Task
	var context Context
	var TaskID int
	var TaskTitle string
	var TaskContent string
	var TaskCreated time.Time
	var getTasksql string

	getTasksql = "SELECT id, title, content, created_date FROM task;"

	rows, err := database.Query(getTasksql)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&TaskID, &TaskTitle, &TaskContent, &TaskCreated)
		if err != nil {
			fmt.Println(err)
		}

		TaskCreated = TaskCreated.Local()
		task := Task{
			ID:      TaskID,
			Title:   TaskTitle,
			Content: TaskContent,
			Created: TaskCreated.Format(time.UnixDate)[0:20],
		}

		tasks = append(tasks, task)
	}
	context = Context{Tasks: tasks}
	return context
}

//AddTaskFunc is used to handle the addition of new task, "/add" URL
func AddTaskFunc(w http.ResponseWriter, r *http.Request) {
	title := "random title"
	content := "random content"
	truth := AddTask(title, content)
	if truth != nil {
		log.Fatal("Error adding task")
	}
	w.Write([]byte("Added task"))
}

//AddTask is used to add the task in the database
func AddTask(title, content string) error {
	query := `INSERT INTO task(title, content, created_date, last_modified_at)
                VALUES(?,?, NOW(), NOW())`
	restoreSQL, err := database.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}

	tx, err := database.Begin()

	_, err = tx.Stmt(restoreSQL).Exec(title, content)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	} else {
		log.Print("insert successful")
		tx.Commit()
	}
	return err
}
