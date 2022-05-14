package router

import (
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"time"
)

type JWTClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

func JwtEncryption(userName string) string {
	var (
		err          error
		signedString string
	)
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	encryptionKey := []byte("nft-manager-key-encryption")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":      time.Now().Unix() - 5,
		"exp":      time.Now().Unix() + 60*60*48,
		"iss":      "nft-manager",
		"username": userName,
	})
	// Sign and get the complete encoded token as a string using the secret
	if signedString, err = t.SignedString(encryptionKey); err != nil {
		log.Errorf("JWT signed string error: %v", err)
		return ""
	}
	return signedString
}

func JwtDecryption(signedString string) (bool, string) {
	var (
		err              error
		decryptionString *jwt.Token
	)
	if decryptionString, err = jwt.ParseWithClaims(signedString, &JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("nft-manager-key-encryption"), nil
		}); err != nil {
		log.Errorf("JWT decryption error: %v", err)
		return false, ""
	}
	if !decryptionString.Valid {
		return false, ""
	}
	return true, decryptionString.Claims.(*JWTClaims).UserName
}
