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
func initDB() *sql.DB{
	db, err := sql.Open("sqlite3", "students.db")
	if err != nil{
		log.Fatal(err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		grade TEXT NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil{
		log.Fatal(err)
	}

	return db
}

/*Add a new student */
func addStudent(db *sql.DB, name, grade string){
	insertSQL := `INSERT INTO students (name, grade) VALUES (?, ?)`
	_, err := db.Exec(insertSQL, name, grade)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("Students %s with grade %s added successfully!\n", name, grade)
}

/*List all students */
func listStudents(db *sql.DB){
	rows, err := db.Query("SELECT id, name, grade FROM students")

	if err != nil{
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("Students List:")
	for rows.Next(){
		var student Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Grade); err != nil{
			log.Fatal(err)
		}

		fmt.Printf("ID: %d, Name: %s, Grade: %s\n", student.ID, student.Name, student.Grade)
	}
}

/*Update student information */
func updateStudent(db *sql.DB, id int, name, grade string){
	updateSQL := `UPDATE students SET name = ?, grade = ? WHERE id = ?`
	_, err := db.Exec(updateSQL, name, grade, id)

	if err != nil{
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
func main(){
	db := initDB()
	defer db.Close()

	// Add some students
	addStudent(db, "Alice", "A")
	addStudent(db, "Bob", "B+")
	addStudent(db, "Charlie", "C")

	// List all students
	listStudents(db)

	// Update a student's information
	updateStudent(db, 2, "Bobby", "A-")

	// List all students after update
	listStudents(db)

	// Delete a student
	deleteStudent(db, 1)

	// List all students after deletion
	listStudents(db)
}