package controllers

import (
	"net/http"
	"project/config"
	"project/models"

	"github.com/labstack/echo/v4"
)

func GetSaveControllers(c echo.Context) error {
	DB := config.Connect()
	var saveResep []models.SaveResep

	if err := DB.Find(&saveResep).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, saveResep)
}

// by id
func GetSaveResepController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")
	saveResep := models.SaveResep{}

	DB.Where("ID = ?", id).First(&saveResep)

	if saveResep.ID == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, saveResep)
}

func CreateSaveResepController(c echo.Context) error {
	DB := config.Connect()
	saveResep := models.SaveResep{}
	c.Bind(&saveResep)

	if err := DB.Save(&saveResep).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, saveResep)
}

func DeleteSaveResepController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")

	DB.Delete(&models.SaveResep{}, id)

	return c.JSON(http.StatusOK, nil)
}

func UpdateSaveResepController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")
	saveResep := models.SaveResep{}

	DB.Where("ID = ?", id).First(&saveResep)

	if saveResep.ID == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}

	payload := models.SaveResep{}
	c.Bind(&payload)

	saveResep.ID = payload.ID
	saveResep.IdUser = payload.IdUser
	saveResep.IdResep = payload.IdResep
	DB.Save(&saveResep)

	return c.JSON(http.StatusOK, saveResep)
}

// func SaveResep(c echo.Context) error {
// 	DB := config.Connect()
// 	saveResep := models.SaveResep{}
// 	c.Bind(&saveResep)

// 	if err := DB.Save(&saveResep).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success create new user",
// 		"user":    saveResep,
// 	})
// }

