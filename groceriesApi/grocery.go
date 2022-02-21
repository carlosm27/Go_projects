package main

import "gorm.io/gorm"

type Grocery struct {
	gorm.Model
	Name     string `json: "name"`
	Quantity int    `json: "quantity"`
}
