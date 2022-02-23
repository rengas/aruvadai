package service

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID           string    `db:"id"`
	Name         string    `db:"name"`
	FiveKg       int       `db:"five_kg"`
	TenKg        int       `db:"ten_kg"`
	TwentyFiveKg int       `db:"twenty_five_kg"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type Aruvadai struct {
	db *sqlx.DB
}

func NewAruvadai(d *sqlx.DB) *Aruvadai {
	return &Aruvadai{
		db: d,
	}
}

func (a *Aruvadai) Index(c *fiber.Ctx) error {

	var products []Product
	err := a.db.Select(&products, "SELECT * FROM product")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(products)
	return c.Render("index", fiber.Map{
		"Title": "அறுவடை ...",
	})
}
