package user

import "time"

type User struct {
	Id          int        `json:"id,omitempty"`
	Username    string     `json:"username,omitempty"`
	Password    string     `json:"password,omitempty"`
	DisplayName string     `json:"displayName,omitempty"`
	UserRole    string     `json:"userRole,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
}

type GeneratePasswordRequest struct {
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
