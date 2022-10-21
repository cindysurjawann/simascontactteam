package manageuser

import (
	"testing"

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

func TestGetUserRepo(t *testing.T){
	db := newTestDB(t)
	repo := NewRepository(db)

	var User model.User
	res, err := repo.GetUser()
	assert.NoError(t, err)
	assert.NotNil(t, res)

	db.Migrator().DropTable(&User)
	repo = NewRepository(db)
	res, err = repo.GetUser()
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetUserReqRepo(t *testing.T){
	db := newTestDB(t)
	repo := NewRepository(db)

	var User model.User
	res, err := repo.GetUserReq("cindu")
	assert.NoError(t, err)
	assert.NotNil(t, res)

	res, err = repo.GetUserReq("candu")
	assert.Error(t, err)
	assert.Equal(t, res, User)

	db.Migrator().DropTable(&User)
	repo = NewRepository(db)
	res, err = repo.GetUserReq("test")
	assert.Error(t, err)
	assert.Equal(t, res, User)
}

func TestUpdateUserRepo(t *testing.T){
	db := newTestDB(t)
	repo := NewRepository(db)

	var User model.User
	user := model.User{
		Email: "tes",
		Role: 2,
		Name: "tes",
	}
	
	res, err:= repo.UpdateUser(user, "cindu")
	assert.NoError(t, err)
	assert.NotNil(t, res)

	res, err = repo.UpdateUser(user, "")
	assert.Error(t, err)
	assert.Equal(t, res, User)
}

func TestDeleteRepo(t *testing.T){
	db := newTestDB(t)
	repo := NewRepository(db)
	var User model.User

	res, err := repo.DeleteUser("1")
	assert.NoError(t, err)
	assert.NotNil(t, res)

	db.Migrator().DropTable(&User)
	repo = NewRepository(db)
	res, err = repo.DeleteUser("1aaaa")
	assert.Error(t, err)
	assert.NotNil(t, res)
}
