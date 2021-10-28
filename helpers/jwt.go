package helpers

import (
	"bitbucket.org/friendsonly/core/config"
	"bitbucket.org/friendsonly/core/models"
	"github.com/dgrijalva/jwt-go"
)

func NewToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.ID

	return token.SignedString([]byte(config.Config("JWT_SECRET")))
}
