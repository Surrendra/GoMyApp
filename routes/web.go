package routes

import (
	"fmt"
	"myapp/config"
	"myapp/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

func InitRoute() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to go | gin")
	})
	router.GET("/get_users", func(c *gin.Context) {
		db, err := config.InitDatabase()
		if err != nil {
			fmt.Println(err)
		}

		// users := map[string]interface{}{}
		// db.Model(&model.User{}).First(users)
		var users []model.User
		errors := db.Model(&model.User{}).Preload("Role").Find(&users).Error

		if errors != nil {
			fmt.Println(errors)
			c.JSON(http.StatusNotFound, gin.H{
				"status":  404,
				"message": "Your data is empty",
				"data":    nil,
			})
		}
		fmt.Println(users)
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "Your data",
			"data":    users,
		})
	})
	router.Run(":1212")
}

func GetALlUser(c echo.Context) error {
	db, err := config.InitDatabase()
	if err != nil {
		fmt.Println(err)
	}

	result := db.Model(&model.User{})
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, gin.H{
			"status":  502,
			"message": "server StatusBadRequest",
			"data":    nil,
		})
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status":  400,
		"message": "Succesfully get data users",
		"data":    nil,
	})
	print(result.Rows())
	return nil

}
