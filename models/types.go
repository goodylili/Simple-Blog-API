package models

import "gorm.io/gorm"

type Message struct {
	Status      string
	Description string
}

type Blogs struct {
	// id is the assigned number after the link
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Reads  int    `json:"reads"`
}
