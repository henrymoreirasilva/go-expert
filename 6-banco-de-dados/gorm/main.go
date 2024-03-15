package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primatyKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(172.17.0.3:3306)/goexpert"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// db.Create(&Product{
	// 	Name:  "Teclado",
	// 	Price: 185.70,
	// })

	db.Create(&[]Product{
		{Name: "Teclado", Price: 150},
		{Name: "Mouse", Price: 75},
		{Name: "Modem", Price: 450.90},
	})

}
