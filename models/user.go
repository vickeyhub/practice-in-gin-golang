package models

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name" binding:"required,min=8"`
	Email   string  `json:"email" binding:"required,email"`
	Age     int     `json:"age" binding:"required,gte=18"`
	Role    string  `json:"role" binding:"required,oneof=admin user manager"`
	Address Address `json:"address" binding:"required"`
}

type Address struct {
	City string `json:"city" binding:"required"`
}
