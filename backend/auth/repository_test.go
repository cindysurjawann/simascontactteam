package auth

import (
	"errors"
	"testing"
	"time"

	_ "errors"

	"github.com/bagasalim/simas/model"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)
	err = db.AutoMigrate(&model.User{}, &model.UserOTP{})
	assert.NoError(t, err)
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

	NewUser := model.User{
		Username: "cindu",
		Password: string(passwordHash),
		Name:     "Cindy",
		Role:     2,
	}
	err = db.Create(&NewUser).Error
	assert.NoError(t, err)

	return db
}
func TestCreateUserService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	// repo.
	User := model.User{
		Username: "remasertu",
		Password: "123456",
		Name:     "rema",
		Role:     2,
	}
	// task := "task 1"
	res, err := repo.AddUser(User)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	User = model.User{
		Username: "remasertu",
	}
	_, err = repo.AddUser(User)
	// fmt.Println(err())
	assert.NotNil(t, err)
}
func TestFindUser(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	User := model.User{
		Username: "remasertu",
		Password: "123456",
		Name:     "rema",
		Role:     2,
	}
	repo.AddUser(User)
	res, err := repo.FindUser("remasertu")
	assert.NoError(t, err)
	assert.NotNil(t, res)
	_, err1 := repo.FindUser("remasertu1")
	assert.Equal(t, err1.Error(), errors.New("not found").Error())
}

func TestAddLastLogin(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	User := model.User{
		Username:  "cindu",
		LastLogin: time.Now(),
	}
	//success
	res, err := repo.AddLastLogin(User.Username, User.LastLogin)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}
