package handler

import (
	"net/http"
	"strings"

	"github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/auth"
	"github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/helper"
	"github.com/labstack/echo/v4"
)

type authController struct {
	service auth.UseCase
}

func New(us auth.UseCase) auth.Handler {
	return &authController{
		service: us,
	}
}

// LoginHandler implements auth.Handler
func (uc *authController) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginInput
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("error on bind login input", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "invalid input", nil))
		}

		res, err := uc.service.LogInLogic(input.ID, input.Password)
		if err != nil {
			c.Logger().Error("error on calling Login Logic", err.Error())

			if strings.Contains(err.Error(), "not exist") {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "username does not exist, please sign up", nil))
			} else if strings.Contains(err.Error(), "wrong") {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "password is wrong please try again", nil))
			} else if strings.Contains(err.Error(), "blank") {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "username is blank please try again", nil))
			}
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "internal server error", nil))
		}
		token, err := helper.GenerateToken(res.ID)
		if err != nil {
			c.Logger().Error("error on generation token", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}
		var data = new(LoginResponse)
		data.Token = token

		if input.Password == "admin" {
			data.Role = "admin"
		} else {
			data.Role = "user"
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "succes login!", data))
	}
}

func (uc *authController) SignValidationLogic() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			req SignVaidation
		)

		if err := c.Bind(&req); err != nil {
			c.Logger().Error("error on bind login input", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "invalid input", nil))
		}

		result, err := uc.service.SignVallidationLogic(req.SignID)
		if err != nil {
			c.Logger().Error("error on finding sign id data", err.Error())

			if strings.Contains(err.Error(), "not found") {
				return c.JSON(helper.ReponseFormatWithMeta(
					http.StatusNotFound,
					"sign data not found",
					nil, nil))
			}

			return c.JSON(helper.ReponseFormatWithMeta(http.StatusInternalServerError,
				"server error",
				nil, nil))
		}
		response := SignResponse{
			SubmissionTitle:  result.Title,
			OfficialName:     result.Officialname,
			OfficialPosition: result.Officialposition,
			Date:             result.Date,
		}

		return c.JSON(helper.ReponseFormatWithMeta(http.StatusOK,
			"succes to get sign data, data generated by system",
			response, nil))
	}
}
