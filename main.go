package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "dream"
	dbName   = "gogo"
)

type Task struct {
	ID        int
	Name      string
	Completed bool
}

// connect to the database
func connectDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", dbUser, dbPass, dbName)
	db, err := sql.Open(dbDriver, connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// create a new task
func createTask(name string) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM tasks WHERE name = ?", name).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("task with name '%s' already exists", name)
	}

	_, err = db.Exec("INSERT INTO tasks (name) VALUES (?)", name)
	if err != nil {
		return err
	}
	fmt.Println("Task created successfully")
	return nil
}

// fetch and display tasks
func listTasks() ([]Task, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Name, &task.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// task completed
func completeTask(id int) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE tasks SET completed = true WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("Task marked as completed")
	return nil
}

// delete a task
func deleteTask(id int) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("Task deleted successfully")
	return nil
}

func main() {
	err := createTask("Task 1")
	if err != nil {
		fmt.Println("Error:", err)
	}
	tasks, err := listTasks()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Tasks:", tasks)
	err = completeTask(1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = deleteTask(1)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
