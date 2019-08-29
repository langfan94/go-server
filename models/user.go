package models

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}
