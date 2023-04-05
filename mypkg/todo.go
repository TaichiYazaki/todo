package mypkg

import "gorm.io/gorm"

type Todo struct{
	gorm.Model
	Text string
	Status string
}