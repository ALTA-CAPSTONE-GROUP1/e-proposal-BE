package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/admin/position"
	"github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type positionController struct {
	service position.UseCase
}

func New(pu position.UseCase) position.Handler {
	return &positionController{
		service: pu,
	}
}

func (pc *positionController) AddPositionHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(AddPositionRequest)
		userID := helper.DecodeToken(c)
		if userID != "admin" {
			log.Error("user are not admin try to acces add position")
			return c.JSON(helper.ResponseFormat(http.StatusUnauthorized, "you are not admin", nil))
		}

		if err := c.Bind(req); err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "invalid input", nil))
		}

		if err := c.Validate(req); err != nil {
			log.Error("errror in validate input" + err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "bad request, invalid input", nil))
		}
		newPosition := position.Core{
			Name: req.Position,
			Tag:  req.Tag,
		}

		if err := pc.service.AddPositionLogic(newPosition); err != nil {
			log.Errorf("error occurs on calling Position logic with data %v, %v", newPosition.Name, newPosition.Tag)
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "internal server error", nil))
		}

		return c.JSON(helper.ResponseFormat(http.StatusCreated, "succes to create position", nil))
	}
}

func (pc *positionController) GetAllPositionHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := helper.DecodeToken(c)
		if userID != "admin" {
			log.Error("user are not admin try to acces add position")
			return c.JSON(helper.ResponseFormat(http.StatusUnauthorized, "you are not admin", nil))
		}

		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		search := c.QueryParam("search")

		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			log.Errorf("limit are not a number %v", limit)
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "Server Error, limit are NaN", nil))
		}
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			log.Errorf("offset are not a number %v", offset)
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "Server Error, offset are NaN", nil))
		}
		if limitInt < 0 || offsetInt < 0 {
			c.Logger().Error("error occurs because limit/offset are negatif")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "limit and offset cannot negative", nil))
		}

		positions, err := pc.service.GetPositionsLogic(limitInt, offsetInt, search)
		if err != nil {
			c.Logger().Error("error occurs when calling GetPositionsLogic")
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "Server Error", nil))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "succes to get positions data", positions))
	}
}

func (pc *positionController) DeletePositionHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := helper.DecodeToken(c)
		if userID != "admin" {
			log.Error("user are not admin try to acces add position")
			return c.JSON(helper.ResponseFormat(http.StatusUnauthorized, "you are not admin", nil))
		}
		position := c.QueryParam("position")
		tag := c.QueryParam("tag")

		if position == "" || tag == "" {
			log.Error("position or tag are empty string")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "data to delete are empty", nil))
		}

		if err := pc.service.DeletePositionLogic(position, tag); err != nil {
			if strings.Contains(err.Error(), "count position query error") {
				log.Error("errors occurs when counting the datas for delete")
				return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "server error", nil))
			}
			if strings.Contains(err.Error(), "no data found for deletion") {
				log.Error("no position data found for deletion")
				return c.JSON(helper.ResponseFormat(http.StatusNotFound, "position data not found", nil))
			}
			log.Error("unexpected error")
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "unexpected server error", nil))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "succes to delete position data", nil))
	}
}