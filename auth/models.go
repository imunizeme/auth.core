package auth

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/imunizeme/config.core"
	"github.com/prest/adapters/postgres"
)

// ImunizemeClaims JWT
type ImunizemeClaims struct {
	jwt.StandardClaims
}

// LoggedUser representation
type LoggedUser struct {
	ID        int    `json:"id,omitempty"`
	Login     string `json:"login,omitempty"`
	ProfileID int    `json:"profile_id,omitempty"`
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
	sqlQuery := `SELECT u.id, u.cpf_cnpj as login, p.id as profile_id
	 FROM users u JOIN profile p ON (u.id = p.user_id)
	 WHERE u.cpf_cnpj = $1 AND
	 u.password = $2 LIMIT 1`
	sc := postgres.Query(sqlQuery, email, hashPassword(password))
	if sc.Err() != nil {
		return
	}
	_, err = sc.Scan(&user)
	return
}

func hashPassword(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
