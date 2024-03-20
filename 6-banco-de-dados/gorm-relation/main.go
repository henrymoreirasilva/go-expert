package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product // has many
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey`
	Value     string
	ProductID int // has one Product
}

type Tag struct {
	ID       int `gorm:"primaryKey`
	Name     string
	Products []Product `gorm:"many2many:products_tags;"` // many to many
}

type Product struct {
	ID           int `gorm:"primatyKey"`
	Name         string
	Price        float64
	CategoryID   int // belongs to Category
	Category     Category
	SerialNumber SerialNumber // has one SerialNumber
	Tags         []Tag        `gorm:"many2many:products_tags;"` // many to many
	gorm.Model                // inclui funcionalidades do GORM
}

func main() {
	//dsn := "root:root@tcp(172.17.0.3:3306)/goexpert"
	dsn := "root:root@tcp(172.17.0.3:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{}, &Tag{})

	// Create category
	db.Create(&[]Category{
		{Name: "Eletrônicos"},
		{Name: "Livros"},
	})

	tag1 := Tag{Name: "Promoção"}
	db.Create(&tag1)

	tag2 := Tag{Name: "Lançamento"}
	db.Create(&tag2)

	// Create product
	db.Create(&[]Product{
		{Name: "Smartphone", Price: 2000, CategoryID: 1, Tags: []Tag{tag1, tag2}},
		{Name: "Notebook", Price: 4999, CategoryID: 1, Tags: []Tag{}},
		{Name: "Teclado", Price: 200, CategoryID: 1, Tags: []Tag{}},
		{Name: "PHP para iniciantes", Price: 50, CategoryID: 2, Tags: []Tag{tag1, tag2}},
		{Name: "Laravel V1", Price: 50, CategoryID: 2, Tags: []Tag{tag1, tag2}},
		{Name: "Java", Price: 120, CategoryID: 2, Tags: []Tag{}},
	})

	// Create serial number
	db.Create(&[]SerialNumber{
		{Value: "123", ProductID: 1},
		{Value: "840", ProductID: 2},
		{Value: "555", ProductID: 4},
	})

	// Listando relacionamento BELONGS TO e HAS ONE
	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, p := range products {
	// 	println(p.Name, p.Category.Name, p.SerialNumber.Value)
	// }

	// Listando relacionamentos HAS MANY
	// var categories []Category
	// err = db.Model(&Category{}).Preload("Products").Find(&categories).Error      // sem SerialNumber
	// err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error     // com SerialNumber
	// if err != nil {
	// 	panic(err)
	// }

	// for _, c := range categories {
	// 	println(c.Name)
	// 	for _, p := range c.Products {
	// 		println("- ", p.Name, p.SerialNumber.Value)
	// 	}
	// }
}
