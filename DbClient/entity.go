package main

import "github.com/jinzhu/gorm"

type Entity struct {
	gorm.Model
	Text string
}
