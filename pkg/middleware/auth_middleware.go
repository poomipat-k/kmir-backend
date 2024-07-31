package appMiddleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/poomipat-k/kmir-backend/pkg/utils"
)

func IsLoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := getAccessToken(r)
		if err != nil {
			utils.ErrorJSON(w, err, "authToken", http.StatusUnauthorized)
			return
		}

		claims, ok := accessToken.Claims.(jwt.MapClaims)
		if ok {
			userId := fmt.Sprintf("%v", claims["userId"])
			username := fmt.Sprintf("%v", claims["username"])
			userRole := fmt.Sprintf("%v", claims["userRole"])
			r.Header.Set("userId", userId)
			r.Header.Set("username", username)
			r.Header.Set("userRole", userRole)
			next(w, r)
		} else {
			utils.ErrorJSON(w, errors.New("corrupt token"), "authToken", http.StatusUnauthorized)
			return
		}
	})
}

func IsUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := getAccessToken(r)
		if err != nil {
			utils.ErrorJSON(w, err, "authToken", http.StatusUnauthorized)
			return
		}

		claims, ok := accessToken.Claims.(jwt.MapClaims)
		if ok {
			userId := fmt.Sprintf("%v", claims["userId"])
			username := fmt.Sprintf("%v", claims["username"])
			userRole := fmt.Sprintf("%v", claims["userRole"])
			if userRole != "user" {
				utils.ErrorJSON(w, errors.New("user permission denied"), "authToken", http.StatusForbidden)
				return
			}
			r.Header.Set("userId", userId)
			r.Header.Set("username", username)
			r.Header.Set("userRole", userRole)
			next(w, r)
		} else {
			utils.ErrorJSON(w, errors.New("corrupt token"), "authToken", http.StatusUnauthorized)
			return

		}
	})
}

func IsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := getAccessToken(r)
		if err != nil {
			utils.ErrorJSON(w, err, "authToken", http.StatusUnauthorized)
			return
		}

		claims, ok := accessToken.Claims.(jwt.MapClaims)
		if ok {
			userId := fmt.Sprintf("%v", claims["userId"])
			username := fmt.Sprintf("%v", claims["username"])
			userRole := fmt.Sprintf("%v", claims["userRole"])
			if userRole != "admin" {
				utils.ErrorJSON(w, errors.New("admin permission denied"), "authToken", http.StatusForbidden)
				return
			}
			r.Header.Set("userId", userId)
			r.Header.Set("username", username)
			r.Header.Set("userRole", userRole)
			next(w, r)
		} else {
			utils.ErrorJSON(w, errors.New("corrupt token"), "authToken", http.StatusUnauthorized)
			return

		}
	})
}

func IsAdminOrViewer(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := getAccessToken(r)
		if err != nil {
			utils.ErrorJSON(w, err, "authToken", http.StatusUnauthorized)
			return
		}

		claims, ok := accessToken.Claims.(jwt.MapClaims)
		if ok {
			userId := fmt.Sprintf("%v", claims["userId"])
			username := fmt.Sprintf("%v", claims["username"])
			userRole := fmt.Sprintf("%v", claims["userRole"])
			if userRole == "admin" || userRole == "viewer" {
				r.Header.Set("userId", userId)
				r.Header.Set("username", username)
				r.Header.Set("userRole", userRole)
				next(w, r)
				return
			}
			utils.ErrorJSON(w, errors.New("admin or viewer permission denied"), "authToken", http.StatusForbidden)
			return

		} else {
			utils.ErrorJSON(w, errors.New("corrupt token"), "authToken", http.StatusUnauthorized)
			return

		}
	})
}

func getAccessToken(r *http.Request) (*jwt.Token, error) {
	// Cookie
	cookie, err := r.Cookie("authToken")
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		// validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_ACCESS_TOKEN_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
