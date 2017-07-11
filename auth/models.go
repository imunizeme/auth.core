package auth

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/imunizeme/config.core"
	"github.com/nuveo/prest/adapters/postgres"
)

// ImunizemeClaims JWT
type ImunizemeClaims struct {
	jwt.StandardClaims
}

// LoggedUser representation
type LoggedUser struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	CPF  string `json:"cpf,omitempty"`
}

// Token for user
func Token(u LoggedUser) string {
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	claims := ImunizemeClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
			Id:        strconv.Itoa(u.ID),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "imunizeme",
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := tok.SignedString([]byte(config.Get.JWTKey))
	return signedToken
}

// Authenticate user
func Authenticate(email, password string) (user LoggedUser, err error) {
	users := make([]LoggedUser, 0)
	sqlQuery := `SELECT u.id, u.name, u.cpf  
	 FROM users u 
	 WHERE u.cpf = $1 AND
	 u.password = $2 LIMIT 1`
	jsonData, err := postgres.Query(sqlQuery, email, hashPassword(password))
	if err != nil {
		return
	}
	if err = json.Unmarshal(jsonData, &users); err != nil {
		return
	}
	if len(users) > 0 {
		user = users[0]
	}
	return
}

func hashPassword(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
