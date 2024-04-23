package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/henrymoreirasilva/go-expert/9-APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(entity.Product{})

	product, err := entity.NewProduct("Mouse", 104.12)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, product.Name, "Mouse")

}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory2"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)

		err = db.Create(product).Error
		assert.NoError(t, err)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 15", products[4].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 4)
	assert.Equal(t, "Product 22", products[1].Name)

}

func TestFindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory3"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Teclado", rand.Float64()*100)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	findProduct, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.NotEmpty(t, findProduct.ID)
	assert.Equal(t, findProduct.ID, product.ID)

}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory3"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Teclado", rand.Float64()*100)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	product.Name = "Teclado 60%"
	err = productDB.Update(product)
	assert.NoError(t, err)

	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, "Teclado 60%")

}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory3"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Teclado", rand.Float64()*100)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	productDelete, err := productDB.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Equal(t, productDelete.Name, "")
}
