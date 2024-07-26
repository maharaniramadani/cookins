package controllers

import (
	"net/http"
	"project/config"
	"project/models"

	"github.com/labstack/echo/v4"
)

func GetBahanControllers(c echo.Context) error {
	DB := config.Connect()
	var bahan []models.Bahan

	if err := DB.Find(&bahan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, bahan)
}

//bahan by id
func GetBahanController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")
	bahan := models.Bahan{}

	DB.Where("ID = ?", id).First(&bahan)

	if bahan.ID == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, bahan)
}

func CreateBookController(c echo.Context) error {
	DB := config.Connect()
	bahan := models.Bahan{}
	c.Bind(&bahan)

	if err := DB.Save(&bahan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, bahan)
}

func DeleteBookController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")

	DB.Delete(&models.Bahan{}, id)

	return c.JSON(http.StatusOK, nil)
}
