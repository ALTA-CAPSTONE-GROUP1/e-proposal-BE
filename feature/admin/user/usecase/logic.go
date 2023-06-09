package usecase

import (
	"errors"
	"strings"

	"github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/admin/user"
	"github.com/labstack/gommon/log"
)

type userLogic struct {
	u user.Repository
}

func New(u user.Repository) user.UseCase {
	return &userLogic{
		u: u,
	}
}

// DeleteUser implements user.UseCase
func (ul *userLogic) DeleteUser(id string) error {
	err := ul.u.DeleteUser(id)
	if err != nil {
		log.Error("failed on calling deleteuser query")
		if strings.Contains(err.Error(), "finding user") {
			log.Error("error on finding user (not found)")
			return errors.New("bad request, user not found")
		} else if strings.Contains(err.Error(), "cannot delete") {
			log.Error("error on delete user")
			return errors.New("internal server error, cannot delete user")
		}
		log.Error("error in delete user (else)")
		return err
	}
	return nil
}

// UpdateUser implements user.UseCase
func (ul *userLogic) UpdateUser(id string, updateUser user.Core) error {
	if err := ul.u.UpdateUser(id, updateUser); err != nil {
		log.Error("failed on calling updateprofile query")
		if strings.Contains(err.Error(), "hashing password") {
			log.Error("hashing password error")
			return errors.New("is invalid")
		} else if strings.Contains(err.Error(), "affected") {
			log.Error("no rows affected on update user")
			return errors.New("data is up to date")
		}
		return err
	}
	return nil
}

// GetUserById implements user.UseCase
func (ul *userLogic) GetUserById(id string) (user.Core, error) {
	result, err := ul.u.GetUserById(id)
	if err != nil {
		log.Error("failed to find user", err.Error())
		return user.Core{}, errors.New("internal server error")
	}

	return result, nil
}

// GetAllUser implements user.UseCase
func (ul *userLogic) GetAllUser(limit, offset int, name string) ([]user.Core, int, error) {
	result, totaldata, err := ul.u.SelectAllUser(limit, offset, name)
	if err != nil {
		log.Error("failed to find all user", err.Error())
		return []user.Core{}, totaldata, errors.New("internal server error")
	}

	return result, totaldata, nil
}

// RegisterUser implements user.UseCase
func (ul *userLogic) RegisterUser(newUser user.Core) error {
	if err := ul.u.InsertUser(newUser); err != nil {
		log.Error("error on calling register insert user query", err.Error())
		if strings.Contains(err.Error(), "column") {
			return errors.New("server error")
		} else if strings.Contains(err.Error(), "value") {
			return errors.New("invalid value")
		} else if strings.Contains(err.Error(), "too short") {
			return errors.New("invalid password length")
		}
		return errors.New("server error")
	}
	return nil
}
