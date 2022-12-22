package svc

import (
	"area-ms/conn"
	"area-ms/domain"
	"log"
	"math"
	"net/http"

	"github.com/labstack/echo"
)

func UpdateDistance(context echo.Context) error {
	for key := range domain.DistAreas {
		var presentList []bool
		conn.Client.Db.Model(&domain.Area{}).Pluck(key, &presentList)
		log.Println(presentList)
		// domain.DistAreas[key] = nil
		// domain.DistAreas[key] = append(domain.DistAreas[key], createDistList(presentList)...)
		domain.DistAreas[key] = createDistList(presentList)
	}

	for key, element := range domain.DistAreas {
		log.Println(key, " ", element)
	}

	return context.String(http.StatusAccepted, "okk")
}

func createDistList(presentList []bool) []int {
	distList := []int{}
	for index := 0; index < len(presentList); index++ {
		if presentList[index] == true {
			distList = append(distList, 0)

		} else {
			if index == 0 {
				distList = append(distList, math.MaxInt)
			} else {
				if distList[index-1] == math.MaxInt {
					distList = append(distList, math.MaxInt)
				} else {
					distList = append(distList, distList[index-1]+1)
				}

			}
		}
	}

	for index := len(presentList) - 2; index >= 0; index-- {
		if distList[index] > distList[index+1]+1 {
			distList[index] = distList[index+1] + 1
		}
	}

	return distList
}
