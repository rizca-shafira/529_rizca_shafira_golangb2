package controllers

import (
	"net/http"
	"strconv"

	"ass3/models"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Water int `json:"water"`
	Wind  int `json: "wind"`
}

func (idb *InDB) CreateData(c *gin.Context) {
	var (
		reload models.Reload
		result gin.H
	)

	water := c.PostForm("water")
	wind := c.PostForm("wind")

	intWater, _ := strconv.Atoi(water)
	intWind, _ := strconv.Atoi(wind)

	reload.Water = intWater
	reload.Wind = intWind

	dataReload := Data{
		Water: reload.Water,
		Wind:  reload.Wind,
	}

	err := idb.DB.Create(&reload).Error

	if err != nil {
		result = gin.H{
			"result": "Data cannot be inserted",
		}
		c.JSON(http.StatusNoContent, result)
	} else {
		result = gin.H{
			"status": dataReload,
		}
		c.JSON(http.StatusOK, result)
	}
}
func (idb *InDB) GetLatestData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

	var (
		reload    models.Reload
		reloadAll []models.Reload
		result    gin.H
	)

	idb.DB.Find(&reloadAll).Take(&reloadAll).Last(&reloadAll)

	if len(reloadAll) <= 0 {
		result = gin.H{
			"message": "Data not found!",
		}
		c.JSON(http.StatusNotFound, result)
	} else {
		idb.DB.Find(&reload).Take(&reload).Last(&reload)

		var waterStatus, windStatus string

		dataReload := Data{
			Water: reload.Water,
			Wind:  reload.Wind,
		}

		if dataReload.Water <= 5 {
			waterStatus = "Aman"
		} else if dataReload.Water >= 6 && dataReload.Water <= 8 {
			waterStatus = "Siaga"
		} else if dataReload.Water > 8 {
			waterStatus = "Bahaya"
		}

		if dataReload.Wind <= 6 {
			windStatus = "Aman"
		} else if dataReload.Wind >= 7 && dataReload.Wind <= 15 {
			windStatus = "Siaga"
		} else if dataReload.Wind > 15 {
			windStatus = "Bahaya"
		}

		result = gin.H{
			"status":       dataReload,
			"waterMessage": waterStatus,
			"windMessage":  windStatus,
		}

		c.JSON(http.StatusOK, result)

	}

}
