package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Student struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone uint64 `json:"phone"`
}

// type Mentor struct {
// 	Name       string `json:"name"`
// 	Spatiality string `json:"expart_in"`
// }

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Ghorai@2001"
	dbname   = "stdb"
)

func openConnection() *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connected successfully...")

	return db
}

func studentTable() {
	db := openConnection()

	defer db.Close()

	queryStr := `create table if not exists Student (student_id serial primary key,  name varchar(255), email varchar(255), phone int)`

	res, err := db.Exec(queryStr)

	if err != nil {
		log.Fatal(err)
	}

	resByte, _ := json.Marshal(res)
	log.Writer().Write(append([]byte("Student table successfully created: "), resByte...))

}

// func mentorTable() {
// 	db := openConnection()

// 	defer db.Close()

// 	queryStr := `create table if not exists Mentor (mentor_id serial primary key, name varchar(255), expart_in varchar(255) )`

// 	res, err := db.Exec(queryStr)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	resByte, _ := json.Marshal(res)
// 	log.Writer().Write(append([]byte("Mentor table successfully created: "), resByte...))

// }

// CURD Operations

func addStudent(w http.ResponseWriter, r *http.Request) {
	db := openConnection()

	defer db.Close()
	w.Header().Set("Content-Type", "application-json")

	var student Student

	json.NewDecoder(r.Body).Decode(&student)

	sqlQuery := `Insert into student (name, email, phone) values($1, $2, $3)`

	_, err := db.Exec(sqlQuery, &student.Name, &student.Email, &student.Phone)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("student added!")

	json.NewEncoder(w).Encode("Student successfully added with the name of " + student.Name)

}

// func addMentor(w http.ResponseWriter, r *http.Request) {
// 	db := openConnection()

// 	defer db.Close()
// 	w.Header().Set("Content-Type", "application-json")

// 	var mentor Mentor

// 	json.NewDecoder(r.Body).Decode(&mentor)

// 	sqlQuery := `Insert into mentor (name, expert_in) values($1, $2)`

// 	_, err := db.Exec(sqlQuery, &mentor.Name, &mentor.Spatiality)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	json.NewEncoder(w).Encode("Mentor successfully added with the name of " + mentor.Name)

// }

func getAllStudent(w http.ResponseWriter, r *http.Request) {
	db := openConnection()

	defer db.Close()
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query("select * from student")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}

	var students []Student

	for rows.Next() {
		var student Student
		rows.Scan(&student.Email, &student.Name, &student.Email, &student.Phone)

		fmt.Println(student)
		students = append(students, student)
	}

	resByte, _ := json.MarshalIndent(students, "", "\t")

	w.Write(resByte)

	fmt.Println("get all student")

}

func getStudent(w http.ResponseWriter, r *http.Request) {
	db := openConnection()
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	urlPath := r.URL.Path

	fmt.Println(urlPath)

	pathVals := strings.Split(urlPath, "/")
	id := pathVals[len(pathVals)-1]

	fmt.Println(id)

	myId, _ := strconv.Atoi(id)

	row, err := db.Query(`select * from student where student_id=$1`, myId)

	var student Student
	row.Columns()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}

	fmt.Println("get student")
	json.NewEncoder(w).Encode(&student)
}

// func getAllMentor(w http.ResponseWriter, r *http.Request) {

// 	fmt.Println(r.Method)
// 	db := openConnection()

// 	defer db.Close()

// 	queryStr := `Select * from mentor`

// 	rows, err := db.Query(queryStr)

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Fatal(err)
// 	}

// 	json.NewEncoder(w).Encode(rows)

// }

// func getMentor(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Hi This is supriya")
// 	db := openConnection()
// 	defer db.Close()

// 	params := url.Values{}

// 	row, err := db.Query(`select * from mentor where mentor_id=$1`, params["id"])

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Fatal(err)
// 	}

// 	resByte, _ := json.MarshalIndent(row, "", "\t")
// 	w.Write(resByte)
// }

func updateStudentEmail(w http.ResponseWriter, r *http.Request) {
	db := openConnection()

	defer db.Close()

	params := r.URL.Query()

	id := params.Get("id")
	params2, _ := strconv.Atoi(id)

	fmt.Println(params2)

	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	queryStr := `update  student set email=$1 where student_id=$2`

	_, err := db.Exec(queryStr, student.Email, params2)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}

	fmt.Println("student email updated")

	json.NewEncoder(w).Encode("Student successfully updated")

}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	db := openConnection()

	defer db.Close()

	w.Header().Set("Content-Type", "application-json")

	params := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(params)

	queryStr := `delete from student where student_id=$1`

	res, err := db.Exec(queryStr, id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("delete student")

	json.NewEncoder(w).Encode(res)
}

func deleteMentor(w http.ResponseWriter, r *http.Request) {
	db := openConnection()

	defer db.Close()

	params := url.Values{}

	queryStr := `delete from mentor where mentor_id=$1`

	_, err := db.Exec(queryStr, params["id"])

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("Mentor deleted")
}

func main() {
	fmt.Println("Building api using http...")
	// mentorTable()
	studentTable()

	http.HandleFunc("/student", student_dispatcher)        // post
	http.HandleFunc("/students", student_dispatcher)       // get
	http.HandleFunc("/student", student_dispatcher)        // get
	http.HandleFunc("/student-email/", student_dispatcher) // put
	http.HandleFunc("/student-one", student_dispatcher)    // delete
	// http.HandleFunc("/mentor", mentor_dispatcher)
	// http.HandleFunc("/mentors", mentor_dispatcher)
	// http.HandleFunc("/mentor/{id}", mentor_dispatcher)
	// http.HandleFunc("/mentor/{id}", mentor_dispatcher)
	defer log.Fatal(http.ListenAndServe(":8000", nil))
}

func student_dispatcher(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	fmt.Println(urlPath)
	switch method := r.Method; method {
	case "GET":
		path := strings.Split(urlPath, "/")[1]
		fmt.Println(path)
		if path == "students" {
			getAllStudent(w, r)
		} else {
			getStudent(w, r)
		}
	case "POST":
		addStudent(w, r)
	case "PUT":
		updateStudentEmail(w, r)
	case "DELETE":
		deleteStudent(w, r)

	}
}

// func mentor_dispatcher(w http.ResponseWriter, r *http.Request) {
// 	switch method := r.Method; method {
// 	case "GET":
// 		getAllMentor(w, r)
// 		getMentor(w, r)
// 	case "POST":
// 		addMentor(w, r)
// 	case "DELETE":
// 		deleteMentor(w, r)

// 	}
// }
