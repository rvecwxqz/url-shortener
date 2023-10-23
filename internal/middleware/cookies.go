package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

var secretKey = "my key"

type ContextKey string

func verifyCookie(cookie *http.Cookie) bool {
	cValue := cookie.Value
	if len(cValue) < 64 {
		return false
	}
	id := cValue[:len(cValue)-64]
	sign, err := hex.DecodeString(cValue[len(cValue)-64:])
	if err != nil {
		fmt.Println(err)
		return false
	}
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(id))
	hash := h.Sum(nil)
	if hmac.Equal([]byte(sign), hash) {
		return true
	} else {
		return false
	}
}

func createCookie() *http.Cookie {
	id := uuid.New().String()
	sign := createSign(id)
	return &http.Cookie{
		Name:  "user_id",
		Value: id + sign,
	}
}

func createSign(id string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(id))
	return hex.EncodeToString(h.Sum(nil))
}

func CookieHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usedID, err := r.Cookie("user_id")
		if err != nil {
			fmt.Println(err, "err")
			usedID = createCookie()
			http.SetCookie(w, usedID)
		} else {
			if !verifyCookie(usedID) {
				usedID = createCookie()
				http.SetCookie(w, usedID)
			}
		}
		uIDKey := ContextKey("user_id")
		ctx := context.WithValue(r.Context(), uIDKey, usedID.Value)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
