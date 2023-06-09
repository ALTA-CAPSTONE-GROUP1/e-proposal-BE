package usecase

import (
	"errors"
	"strings"

	"github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/user/cc"
	"github.com/labstack/gommon/log"
)

type ccLogic struct {
	cl cc.Repository
}

func New(cr cc.Repository) cc.UseCase {
	return &ccLogic{
		cl: cr,
	}
}

func (cl *ccLogic) GetAllCcLogic(userID string) ([]cc.CcCore, error) {
	result, err := cl.cl.GetAllCc(userID)
	if err != nil {
		if strings.Contains(err.Error(), "record") {
			log.Errorf("error not found ccs", err)
			return []cc.CcCore{}, errors.New("record not found")
		}
		return []cc.CcCore{}, err
	}
	return result, nil
}
