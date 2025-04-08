// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// User -.
type User struct {
	Name  string `json:"source" example:"Andrey"`
	Email string `json:"email" example:"test@test.com"`
	Phone string `json:"phone" example:"+79999999999"`
}
