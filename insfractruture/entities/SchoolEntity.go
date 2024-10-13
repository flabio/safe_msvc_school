package entities

import "time"

type School struct {
	Id             uint       `gorm:"primary_key:auto_increment"`
	Url            string     `gorm:"type:text; null"`
	Name           string     `gorm:"type:varchar(150);not null"`
	ProviderNumber string     `gorm:"type:varchar(150);not null"`
	Email          string     `gorm:"type:varchar(150);unique_index;not null"`
	Address        string     `gorm:"type:varchar(150)"`
	Phone          string     `gorm:"type:varchar(150);"`
	ZipCode        string     `gorm:"type:varchar(50);"`
	StateId        uint       `gorm:"type:integer"`
	Active         bool       `gorm:"type:boolean"`
	CreatedAt      time.Time  `gorm:"<-:created_at"`
	UpdatedAt      *time.Time `gorm:"type:TIMESTAMP(6)"`
}
