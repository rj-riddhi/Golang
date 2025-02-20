package helper

import (
	"fmt"
	"log"
	"os"
	"time"

	"githlab.com/radhika.parmar/go-jwt-auth-project/models"
	jwt "github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	User_type  string
	User_id    string
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, first_name string, last_name string, user_type string, user_id string) (token string, refreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: first_name,
		Last_name:  last_name,
		User_type:  user_type,
		User_id:    user_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err.Error())
		return
	}
	return token, refreshToken, err
}

func UpdateAllTokens(db *gorm.DB, token, refreshToken string, userID int) error {
	// Find the user by ID
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}

	// Update the user's token and refresh token
	user.Token = token
	user.Refresh_token = refreshToken

	// Save the updated user
	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("token is not valid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}
