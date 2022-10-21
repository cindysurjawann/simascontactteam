package custom

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type DataJWT struct {
	Username  string
	Name      string
	Role      int8
	LastLogin int64
	jwt.StandardClaims
}

func GenerateJWT(username string, name string, role int8, LastLogin time.Time) (string, error) {
	claim := DataJWT{
		Username:  username,
		Name:      name,
		Role:      role,
		LastLogin: LastLogin.Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte("token"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func MiddlewareAuth(c *gin.Context) {
	auth := c.Request.Header["Authorization"]
	if len(auth) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You not have access",
		})
		fmt.Println("no auth")
		return
	}
	token := auth[0]
	dataUser, err := ClaimToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You not have access",
		})
		return
	}
	c.Set("user", dataUser)
	c.Next()
}

func ClaimToken(tokenString string) (DataJWT, error) {
	claims := DataJWT{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte("token"), nil
	})
	if err != nil {
		fmt.Println("err claimToken", err)
		return DataJWT{}, err
	}
	if !token.Valid {
		fmt.Println("err claimToken", "novalid")
		return DataJWT{}, errors.New("no valid")
	}
	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()){
		return DataJWT{}, errors.New("expire")
	}
	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 0*time.Second {
	// 	return DataJWT{}, errors.New("expire")
	// }
	return claims, nil
}
