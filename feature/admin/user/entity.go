package position

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Core struct {
	ID          string
	OfficeID    int
	PositionID  int
	Name        string
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type Handler interface {
	RegisterHandler() echo.HandlerFunc
}

type UseCase interface {
	RegisterUser(newUser Core) error
}

type Repository interface {
	InsertUser(newUser Core) error
}
