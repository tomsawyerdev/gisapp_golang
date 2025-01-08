package controllers

import (
	//"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	//svc "gisapi/database"
	"gisapi/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
	"net/http"
	//"reflect"
	//"log"
	"strings"
	"time"
)

func SessionGet(c *gin.Context) {
	fmt.Println("getSession:")
	c.IndentedJSON(http.StatusOK, gin.H{"message": "getSession"})

}

type Claims struct {
	UUID string `json:"uuid" `
	jwt.RegisteredClaims
}

func SessionCreate(c *gin.Context) {
	fmt.Println("newSession:")
	// puede ser var o type
	var creds struct {
		Username string `json:"username"  binding:"required"`
		Password string `json:"password"  binding:"required"`
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		//fmt.Println("err:", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid request"})
		return
	}
	// Check for User and Password Hash

	user, err := models.GetUserFromUsername(creds.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 401, "message": "Cant find User"})
		return
	}
	//fmt.Println("User:", user)

	// -----------------------------------

	// Verify HASH Argon
	// demond@mail.com, password
	// hash stored: $2a$12$Q0LciK2AG8x8bgJMe11VGeMMX17HoKO.CgAEhCTnYRkH3umMOMw9i
	// TODA update stored hash to new version

	/*
		match, err := comparePasswordAndHash(creds.Password, user.Hash)
		if err != nil {

			fmt.Printf("Hash error: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"status": 401, "message": "Unauthorized"})
			return
		}
	*/

	//fmt.Printf("Match: %v\n", match)

	match := true // Now is a demo so jump password verification

	if match {

		//exp := jwt.NewNumericDate(time.Now().Add(3 * time.Minute))
		exp := jwt.NewNumericDate(time.Now().Add(24 * time.Hour))

		claims := Claims{
			user.UUID,
			jwt.RegisteredClaims{ExpiresAt: exp}}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Sign and get the complete encoded token as a string using the secret
		// 'Bearer ' +
		tokenString, _ := token.SignedString([]byte("supersecret"))

		//fmt.Println(tokenString, err)
		//c.JSON(http.StatusOK, gin.H{"status": 200, "username": user.UserName, "token": tokenString})
		c.JSON(http.StatusOK, gin.H{"status": "success", "isloged": true, "name": user.UserName, "token": tokenString})
		return

	}
	c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Authentication Fail"})

}

// Argon Helpers

func comparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	fmt.Println("comparePasswordAndHash:--")
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	p, salt, hash, err := decodeHash(encodedHash)

	if err != nil {
		return false, err
	}

	// Derive the key from the other password using the same parameters.
	// Regenerate the hash
	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.

	//fmt.Println("Compare:", hash, otherHash)
	//fmt.Println("Compare:", hash, otherHash) //Ascii
	//fmt.Printf("Compare: %x, %x \n", string(hash[:]), string(otherHash[:])) //Strings
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
