package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Ghorai@2001"
	dbname   = "mydb"
)

func main() {
	fmt.Println("Building apis in Golang with Postgres")

	createPersonTable()
	http.HandleFunc("/persons", getHandler)
	http.HandleFunc("/person", insertPerson)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func openConnection() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Postgres is not able to connect: %s", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("Not able to connect with DB: %s", err)
	}

	fmt.Println("Successfully Connected...")

	return db
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all person")
	db := openConnection()

	w.Header().Set("Content-Type", "application/json")

	var persons []Person

	rows, err := db.Query("select * from person")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var person Person

		rows.Scan(&person.Name, &person.Email)
		persons = append(persons, person)

	}

	personInByte, _ := json.MarshalIndent(persons, "", "\t")

	w.Write(personInByte)
	defer db.Close()
	defer rows.Close()

}

func insertPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding person")
	db := openConnection()

	w.Header().Set("Content-Type", "application/json")

	var person Person

	json.NewDecoder(r.Body).Decode(&person)

	sqlstatement := `insert into person (name, email) values($1, $2)`

	_, err := db.Exec(sqlstatement, &person.Name, &person.Email)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Person successfully added!")

	defer db.Close()

}

func createPersonTable() {
	db := openConnection()

	query := `create table if not exists person ( name varchar(255), email varchar(255) )`

	res, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Person table successfully created!")

	fmt.Println(res)
	defer db.Close()
}
