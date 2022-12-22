package svc

import (
	"log"
	"net/http"

	"area-ms/conn"
	"area-ms/domain"

	"github.com/labstack/echo"
)

func SaveArea(context echo.Context) error {

	var areaList []domain.Area

	err := context.Bind(&areaList)
	if err != nil {
		return context.String(http.StatusBadRequest, "bad request")
	}
	log.Println(areaList)
	RowsAffected := 0
	for index := range areaList {
		log.Println(areaList[index])

		result := conn.Client.Db.Create(&areaList[index])
		if result.Error != nil {
			return context.String(http.StatusBadRequest, "could not save data")
		}
		RowsAffected += int(result.RowsAffected)
	}
	log.Printf("Rows affected %v\n", RowsAffected)

	return context.String(http.StatusOK, "okkk")
}
