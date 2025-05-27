package models

import "time"

type Role string
type Gender string

const (
	ADMIN Role = "ADMIN"
	USER  Role = "USER"
)

const (
	MALE    Gender = "MALE"
	FEMALE  Gender = "FEMALE"
	OTHER   Gender = "OTHER"
	UNKNOWN Gender = "UNKNOWN"
)

type User struct {
	BaseModel

	Email         string     `json:"email" gorm:"unique"`
	Password      string     `json:"password"`
	DisplayName   string     `json:"display_name"`
	Role          Role       `json:"role" gorm:"default:USER"`
	Gender        Gender     `json:"gender" gorm:"default:UNKNOWN"`
	Username      *string    `json:"username" gorm:"unique"`
	RefreshToken  *string    `json:"refresh_token"`
	Bio           *string    `json:"bio"`
	PhotoURL      *string    `json:"photo_url"`
	BackgroundURL *string    `json:"background_url"`
	Location      *string    `json:"location"`
	PhoneNumber   *string    `json:"phone_number"`
	Birthday      *time.Time `json:"birthday"`
}
