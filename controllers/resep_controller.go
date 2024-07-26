package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"project/config"
	"project/constants"
	"project/models"

	"github.com/labstack/echo/v4"
)

func CallApi(c echo.Context) error {

	url := constants.BaseUrl+"/recipes/findByIngredients"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-api-key", "43ab617e49b24debb876ea1043a2b3f7")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return c.JSON(http.StatusOK,"")
}



func GetResepControllers(c echo.Context) error {
	DB := config.Connect()
	var resep []models.Resep

	if err := DB.Find(&resep).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resep)
}

//bahan by id
func GetResepController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")
	resep := models.Resep{}

	DB.Where("ID = ?", id).First(&resep)

	if resep.ID == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, resep)
}

func CreateResepController(c echo.Context) error {
	DB := config.Connect()
	resep := models.Resep{}
	c.Bind(&resep)

	if err := DB.Save(&resep).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resep)
}

func DeleteResepController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")

	DB.Delete(&models.Resep{}, id)

	return c.JSON(http.StatusOK, nil)
}
