package approve

import (
	"github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/admin"
	"github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/user"
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        int
	UserID    string
	TypeID    int
	Title     string
	Message   string
	Status    string
	Is_Opened bool
	Type      admin.Type
	Files     []user.File
	Tos       []user.To
	Ccs       []user.Cc
	Signs     []user.Sign
}

type Handler interface {
	GetSubmissionAprroveHandler() echo.HandlerFunc
}

type UseCase interface {
	GetSubmissionAprrove(userID string, limit, offset int, search string) ([]Core, error)
}

type Repository interface {
	SelectSubmissionAprrove(userID string, limit, offset int, search string) ([]Core, error)
}
