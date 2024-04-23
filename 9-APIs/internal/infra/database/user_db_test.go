package database

import (
	"testing"

	"github.com/henrymoreirasilva/go-expert/9-APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Henry", "h@h.com.br", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)

}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Henry", "h@h.com.br", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)

}
