package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

/*Student Information*/
type Student struct {
	ID      int
	Name    string
	Grade   string
	Age     int
	Address string
	Email   string
	Major   string
}

/*Database for student system*/ 
func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "students.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		grade TEXT NOT NULL,
		age INTEGER,
		address TEXT,
		email TEXT,
		major TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}


/*Add a new student */
func addStudent(db *sql.DB, name, grade string, age int, address, email, major string) {
	insertSQL := `INSERT INTO students (name, grade, age, address, email, major) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(insertSQL, name, grade, age, address, email, major)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Student %s added successfully!\n", name)
}

/*List all students */
func listStudents(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, grade, age, address, email, major FROM students")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Student List:")
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.ID, &student.Name, &student.Grade, &student.Age, &student.Address, &student.Email, &student.Major)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Grade: %s, Age: %d, Address: %s, Email: %s, Major: %s\n", student.ID, student.Name, student.Grade, student.Age, student.Address, student.Email, student.Major)
	}
}


/*Update student information */
func updateStudent(db *sql.DB, id int, name, grade string, age int, address, email, major string) {
	updateSQL := `UPDATE students SET name = ?, grade = ?, age = ?, address = ?, email = ?, major = ? WHERE id = ?`
	_, err := db.Exec(updateSQL, name, grade, age, address, email, major, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Student ID %d updated to %s with grade %s.\n", id, name, grade)
}


/*Delete Student*/
func deleteStudent(db *sql.DB, id int){
	deleteSQL := `DELETE FROM students WHERE id = ?`
	_, err := db.Exec(deleteSQL, id)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("Student ID %d deleted successfully.\n", id)
}

/*Main*/
func main() {
	db := initDB()
	defer db.Close()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [add|list|update|delete]")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 8 {
			fmt.Println("Usage: go run main.go add <name> <grade> <age> <address> <email> <major>")
			return
		}
		age, _ := strconv.Atoi(os.Args[4])
		addStudent(db, os.Args[2], os.Args[3], age, os.Args[5], os.Args[6], os.Args[7])

	case "list":
		listStudents(db)

	case "update":
		if len(os.Args) < 9 {
			fmt.Println("Usage: go run main.go update <id> <name> <grade> <age> <address> <email> <major>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		age, _ := strconv.Atoi(os.Args[5])
		updateStudent(db, id, os.Args[3], os.Args[4], age, os.Args[6], os.Args[7], os.Args[8])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go delete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		deleteStudent(db, id)

	default:
		fmt.Println("Unknown command. Use add, list, update, or delete.")
	}
}