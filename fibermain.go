package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Personel struct {
	ID         int     `json:"id"`
	First_Name string  `json:"first_name"`
	Last_Name  string  `json:"last_name"`
	Age        float32 `json:"age"`
	Work       Work
}
type Work struct {
	Location   string  `json:"location"`
	Department string  `json:"department"`
	Experience float32 `json:"experience"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "172754"
	dbname   = "postgres"
)

func OpenConnention() *sql.DB {

	psq := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psq)

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
func GetHandler(c *fiber.Ctx) error {

	db := OpenConnention()

	var people []Personel
	var person Personel

	rows, err := db.Query("SELECT * FROM personel")
	for rows.Next() {
		rows.Scan(&person.ID, &person.First_Name, &person.Last_Name, &person.Age, &person.Work.Location, &person.Work.Department, &person.Work.Experience)

		people = append(people, person)

	}
	peopleByte, _ := json.MarshalIndent(people, "", "\t")

	c.Write(peopleByte)

	defer rows.Close()
	defer db.Close()

	return err
}
func main() {

	app := fiber.New()
	app.Get("/", GetHandler)

	http.ListenAndServe(":8080", nil)
}