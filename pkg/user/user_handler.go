package user

import (
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/poomipat-k/kmir-backend/pkg/common"
	"github.com/poomipat-k/kmir-backend/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

const accessExpireDurationMinute = 30
const refreshExpireDurationHour = 1440 // 60 days

type UserStore interface {
	GetUserByUsername(username string) (User, error)
}

type UserHandler struct {
	store UserStore
}

func NewUserHandler(s UserStore) *UserHandler {
	return &UserHandler{
		store: s,
	}
}

func (h *UserHandler) GenerateHashedPassword(w http.ResponseWriter, r *http.Request) {
	var payload GeneratePasswordRequest
	err := utils.ReadJSON(w, r, &payload)
	if err != nil {
		utils.ErrorJSON(w, err, "payload", http.StatusBadRequest)
		return
	}

	err = validatePassword(payload.Password)
	if err != nil {
		utils.ErrorJSON(w, err, "password", http.StatusBadRequest)
		return
	}

	hashedPassword, err := generateHashedAndSaltedPassword(payload.Password, 8, "_")
	if err != nil {
		utils.ErrorJSON(w, err, "generate hash", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, hashedPassword)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload LoginRequest
	err := utils.ReadJSON(w, r, &payload)
	if err != nil {
		utils.ErrorJSON(w, err, "payload", http.StatusBadRequest)
		return
	}
	name, err := validateLoginPayload(payload)
	if err != nil {
		utils.ErrorJSON(w, err, name, http.StatusBadRequest)
		return
	}

	user, err := h.store.GetUserByUsername(payload.Username)
	if err != nil {
		utils.ErrorJSON(w, InvalidCredentialError{}, "credential", http.StatusBadRequest)
		return
	}

	err = comparePassword(payload.Password, user.Password)
	if err != nil {
		utils.ErrorJSON(w, InvalidCredentialError{}, "credential", http.StatusBadRequest)
		return
	}

	accessExpiredAtUnix := time.Now().Add(accessExpireDurationMinute * time.Minute).Unix()
	accessToken, err := generateAccessToken(user.Id, user.Username, user.UserRole, accessExpiredAtUnix)
	if err != nil {
		utils.ErrorJSON(w, err, "authToken", http.StatusInternalServerError)
		return
	}
	accessTokenCookie := http.Cookie{
		Name:     "authToken",
		Value:    accessToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/api",
		Expires:  time.Unix(accessExpiredAtUnix, 0),
	}

	refreshExpiredAtUnix := time.Now().Add(refreshExpireDurationHour * time.Hour).Unix()
	refreshToken, err := generateRefreshToken(user, refreshExpiredAtUnix)
	if err != nil {
		utils.ErrorJSON(w, err, "refreshToken", http.StatusInternalServerError)
		return
	}
	refreshTokenCookie := http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/api/v1/auth",
		Expires:  time.Unix(refreshExpiredAtUnix, 0),
	}

	http.SetCookie(w, &accessTokenCookie)
	http.SetCookie(w, &refreshTokenCookie)

	utils.WriteJSON(w, http.StatusOK, common.CommonSuccessResponse{Success: true, Message: "log in successfully"})
}

func (h *UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUsernameFromRequestHeader(r)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "username", http.StatusUnauthorized)
		return
	}
	userRole, err := utils.GetUserRoleFromRequestHeader(r)
	if userRole == "" {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "username", http.StatusUnauthorized)
		return
	}
	user, err := h.store.GetUserByUsername(userId)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrorJSON(w, err, "", http.StatusNotFound)
		return
	}
	utils.WriteJSON(w, http.StatusOK, User{
		Id:       user.Id,
		Username: user.Username,
		UserRole: user.UserRole,
	})
}

// Private methods
func generateHashedAndSaltedPassword(password string, saltLen int, delim string) (string, error) {
	salt := utils.RandAlphaNum(saltLen)

	toHash := strings.Join([]string{password, salt}, "")
	hashed, err := hashPassword(toHash)
	if err != nil {
		return "", err
	}

	passwordToStore := strings.Join([]string{hashed, salt}, delim)
	return passwordToStore, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func generateAccessToken(userId int, username string, userRole string, expiredAtUnix int64) (string, error) {
	accessSecretKey := []byte(os.Getenv("JWT_ACCESS_TOKEN_SECRET_KEY"))

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   userId,
		"username": username,
		"userRole": userRole,
		"iat":      time.Now().Unix(),
		"exp":      expiredAtUnix,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := t.SignedString(accessSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func generateRefreshToken(user User, expiredAtUnix int64) (string, error) {
	refreshTokenSecretKey := []byte(os.Getenv("JWT_REFRESH_TOKEN_SECRET_KEY"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.Id,
		"username": user.Username,
		"userRole": user.UserRole,
		"iat":      time.Now().Unix(),
		"exp":      expiredAtUnix,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := t.SignedString(refreshTokenSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
