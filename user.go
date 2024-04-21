package main

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID            int64   `json:"id,omitempty"`
	Username      string  `json:"username"`
	IsProvider    bool    `json:"isProvider"`
	Password      string  `json:"password"`
	ProfilePicURL string  `json:"profilePicURL"`
	PCSpecs       string  `json:"pcSpecs"`
	Description   string  `json:"description"`
	CloudService  string  `json:"cloudService"`
	CreatedAt     string  `json:"createdAt,omitempty"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}

type Claims struct {
	ID       int64  `json:"id,omitempty"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type MessageDto struct {
	ID        string
	Message   string
	From      string
	Timestamp time.Time
}
