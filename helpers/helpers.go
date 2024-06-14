package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/dgrijalva/jwt-go"
	"github.com/fajarcandraaa/dating_app_be/internal/presentation"
	"golang.org/x/crypto/bcrypt"
)

const (
	alphabet      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	largestuint64 = 18446744073709551615
)

// HashPassword is func to encrypt password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash is func to validate password hasher
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// IsNumber is func to validate type from input string is number or alphabet
func IsNumber(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func Base62Encode(number uint64) string {
	length := len(alphabet)
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(alphabet[(number % uint64(length))])
	}

	return encodedBuilder.String()
}

func Base62Decode(encodedString string) (uint64, error) {
	var number uint64
	length := len(alphabet)

	for i, symbol := range encodedString {
		alphabeticPosition := strings.IndexRune(alphabet, symbol)
		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), errors.New("cannot find symbol in alphabet")
		}
		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}

func GenerateUserCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 7

	rand.New(rand.NewSource(time.Now().UnixNano()))
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

func ParseToken(t string) (*presentation.ClaimTokenUser, error) {
	var claimToken presentation.ClaimTokenUser
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		b, err := json.MarshalIndent(claims, "", " ")
		if err != nil {
			return nil, err
		}

		_ = json.Unmarshal(b, &claimToken)

		return &claimToken, nil
	}

	return nil, fmt.Errorf("Error parse token")
}
