package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"ID"`
	Username string `json:"username"`
}

var secretKey = []byte("go-react-chatroom")

func createToken(username string, id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       id,
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func hashAndSaltPassword(password string) (hashedPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func verifyToken(tokenString string) (float64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return -1, err
	}

	if !token.Valid {
		return -1, fmt.Errorf("invalid token")
	}

	idField := token.Claims.(jwt.MapClaims)["id"]
	id, ok := idField.(float64)

	if !ok {
		return -1, fmt.Errorf("invalid id")
	}

	return id, nil
}

func getUserId(c *gin.Context) float64 {
	token := c.Request.Header.Get("Authorization")[len("Bearer "):]
	id, err := verifyToken(token)
	if err != nil {
		fmt.Print("Invalid token")
		c.JSON(http.StatusBadRequest, err.Error())
		return -1
	}

	return id
}
