package db

import (
	"database/sql"
	"fmt"
	"log"
	"tasks/types"
	"time"

	_ "github.com/lib/pq" // we want to use pg
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

// GetTasks returns tasks from the database
func GetTasks() types.Context {
	var tasks []types.Task
	var context types.Context
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
		task := types.Task{
			ID:      TaskID,
			Title:   TaskTitle,
			Content: TaskContent,
			Created: TaskCreated.Format(time.UnixDate)[0:20],
		}

		tasks = append(tasks, task)
	}
	context = types.Context{Tasks: tasks}
	return context
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
