package main

import (
	"fmt"
	"log"

	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rengas/aruvadai/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test123"
	dbname   = "aruvadai"
)

func main() {

	//Database migarations
	initMigrate()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}

	a := service.NewAruvadai(db)

	app.Get("/", a.Index)

	log.Fatal(app.Listen(":3000"))

}

func initDatabase() (db *sql.DB) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func initMigrate() {

	db := initDatabase()
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}

	m.Down()

	m.Up()
}
