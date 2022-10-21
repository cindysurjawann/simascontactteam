package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoginService(t *testing.T) {
	// os.Setenv("testing", "y")
	// db := newTestDB(t)
	// repo := NewRepository(db)
	// service := NewService(repo)
	// passwordHash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	// if err != nil {
	// 	return
	// }
	// User := model.User{
	// 	Username: "remasertu",
	// 	Password: string(passwordHash),
	// 	Name:     "rema",
	// 	Role:     2,
	// }

	// repo.AddUser(User)
	// data := LoginRequest{
	// 	Username: "remasertu",
	// 	Password: "123456",
	// }
	// res, _, err := service.Login(data)
	// fmt.Println("test", data.Username, res, err)
	// assert.NoError(t, err)
	// assert.NotNil(t, res)

	// data = LoginRequest{
	// 	Username: "remasertu2",
	// 	Password: "123456",
	// }
	// _, _, err = service.Login(data)
	// assert.Equal(t, err.Error(), "username or password is wrong")

	// data = LoginRequest{
	// 	Username: "remasertu",
	// 	Password: "12345",
	// }
	// res, _, err = service.Login(data)
	// assert.NotNil(t, err)

}
func TestCreateAccountService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	data := RegisterRequest{
		Username: "remasertu",
		Password: "123456",
		Name:     "rema",
	}
	res, _, err := service.CreateAccount(data)
	fmt.Println("test", data.Username, res, err)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	data = RegisterRequest{
		Username: "remasertu",
		Password: "123456",
		Name:     "rema",
	}
	_, _, err = service.CreateAccount(data)
	assert.Equal(t, err.Error(), "duplicate data")

	// data = RegisterRequest{
	// 	Username: "remasertu",
	// 	Password: "123456",
	// 	Name:     "rema",
	// }
	// res, _, err = service.Login(data)
	// assert.NotNil(t, err)

}

func TestUpdateLastLoginService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	data := LastLoginRequest{
		Username:  "cindu",
		LastLogin: time.Now(),
	}
	//sukses
	res, _, err := service.UpdateLastLogin(data)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
