// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// User -.
type User struct {
	Id    int    `json:"id"  example:"1"`
	Name  string `json:"name"  example:"Andrey"`
	Email string `json:"email" example:"test@test.com"`
	Phone string `json:"phone" example:"+79999999999"`
}

func NewUser(name, email, phone string) *User {
	return &User{
		Name:  name,
		Email: email,
		Phone: phone,
	}
}
