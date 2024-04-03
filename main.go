package main

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID             uuid.UUID `sql:"type:uuid;primary_key"`
	Title          string    `db:"title"`
	ISBN           string    `db:"isbn"`
	FirstPublished time.Time `db:"first_published"`
	Authors        []Author  `gorm:"many2many:book_author;"`
}

type Author struct {
	gorm.Model
	ID         uuid.UUID `sql:"type:uuid;primary_key"`
	Name       string    `db:"name"`
	ZodiacSign string    `db:"zodiac_sign"`
	Books      []Book    `gorm:"many2many:book_author;"`
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=playground port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Ouw woes!")
	}

	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Author{})

	// many2many insert example
	// datePublished, _ := time.Parse("2006-01-02", "1979-01-01")
	// ghostStory := Book{ID: uuid.New(), Title: "Ghost Story", ISBN: "0-698-10959-7", FirstPublished: datePublished}

	// db.Create(&ghostStory)
	// log.Printf("Insered item id: %s\n", ghostStory.ID)

	// author := Author{ID: uuid.New(), Name: "Peter Straub", ZodiacSign: "Pisces", Books: []Book{ghostStory}}
	// db.Create(&author)
	// log.Printf("Insered item id: %s\n", author.ID)

	// many2many query example
	var book Book
	err = db.Model(&Book{Title: "Ghost Story"}).Preload("Authors").Find(&book).Error
	if err != nil {
		log.Println("Ouw woes again!")
	}

	log.Println(book.ISBN)
	log.Println(book.Authors[0].Name)
	log.Println("...and out")
}
