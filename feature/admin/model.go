package admin

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID          string         `gorm:"primaryKey;size:50"`
	OfficeID    int            `gorm:"foreignKey:OfficeID"`
	PositionID  int            `gorm:"foreignKey:PositionID"`
	Name        string         `gorm:"size:50;not null"`
	Email       string         `gorm:"size:50"`
	PhoneNumber string         `gorm:"size:50"`
	Password    string         `gorm:"size:50"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Office struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"size:50;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Position struct {
	ID        int            `gorm:"primaryKey"`
	Name      string         `gorm:"size:50;not null"`
	Tag       string         `gorm:"size:50;not null"`
	Types     []Type         `gorm:"many2many:position_has_type"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Type struct {
	ID          int            `gorm:"primaryKey;autoIncrement"`
	Name        string         `gorm:"size:50;not null"`
	Requirement string         `gorm:"size:255;not null"`
	Positions   []Position     `gorm:"many2many:position_has_type"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type PositionHasType struct {
	ID         int            `gorm:"primaryKey;autoIncrement"`
	PositionID int            `gorm:"primaryKey"`
	TypeID     int            `gorm:"primaryKey"`
	As         string         `gorm:"size:10;not null"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Position   Position       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Type       Type           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}