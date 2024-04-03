package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title          string    `db:"title"`
	ISBN           string    `db:"isbn"`
	FirstPublished time.Time `db:"first_published"`
	Authors        []Author  `gorm:"many2many:book_author;"`
}

type Author struct {
	gorm.Model
	Name       string `db:"name"`
	ZodiacSign string `db:"zodiac_sign"`
	Books      []Book `gorm:"many2many:book_author;"`
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
	// ghostStory := Book{Title: "Ghost Story", ISBN: "0-698-10959-7", FirstPublished: datePublished}

	// db.Create(&ghostStory)
	// log.Printf("Insered item id: %d\n", ghostStory.ID)

	// author := Author{Name: "Peter Straub", ZodiacSign: "Pisces", Books: []Book{ghostStory}}
	// db.Create(&author)
	// log.Printf("Insered item id: %d\n", author.ID)

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
