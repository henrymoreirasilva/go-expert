package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primatyKey"`
	Name  string
	Price float64
	gorm.Model			// inclui funcionalidades do GORM
}

func main() {
	//dsn := "root:root@tcp(172.17.0.3:3306)/goexpert"
	dsn := "root:root@tcp(172.17.0.3:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{
	// 	Name:  "Teclado",
	// 	Price: 185.70,
	// })

	// Create batch
	// db.Create(&[]Product{
	// 	{Name: "Teclado", Price: 150},
	// 	{Name: "Mouse", Price: 75},
	// 	{Name: "Modem", Price: 450.90},
	// })

	// Select one
	// var product Product
	// db.First(&product, 1)
	// println(product.Name)

	// db.First(&product, "name = ?", "Modem")
	// println(product.Price)

	// Select all
	// var products []Product
	// db.Find(&products)
	// for _, p := range products {
	// 	println(p.Name)
	// }

	// var products []Product
	// db.Limit(1).Offset(2).Find(&products)
	// for _, p := range products {
	// 	println(p.Name)
	// }

	// var products []Product
	// db.Where("price < ?", 100).Find(&products)
	// for _, p := range products {
	// 	println(p.Name, p.Price)
	// }

	// Update
	// var p Product
	// db.First(&p, 1)
	// p.Name = "Teclado RedDragon"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// println(p2.Name)

	// Soft delete
	// db.Delete(&p2)


	// Restore soft deleted
	var p Product
	db.Unscoped().First(&p, 1)
	db.Unscoped().Model(&p).Update("deleted_at", nil)

	var p2 Product
	db.First(&p2, 1)
	println(p2.Name)


}
