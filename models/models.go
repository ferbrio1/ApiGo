//models.go

package main

import "gorm.io/gorm"

type IceCream struct {
	gorm.Model

	Sabor string `gorm:"not null;unique_index" json:"sabor"`
	Tamano string `gorm:"not null" json:"tamano"`
	Precio string `gorm:"not null" json:"precio"`
	Stock string `gorm:"not null;" json:"stock"`

}
