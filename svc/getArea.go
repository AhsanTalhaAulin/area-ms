package svc

import (
	"area-ms/domain"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func GetArea(context echo.Context) error {

	var area domain.Area

	err := context.Bind(&area)
	if err != nil {
		return context.String(http.StatusBadRequest, "bad request")
	}
	log.Println(area)

	log.Println(len(domain.DistAreas))

	printRequiredLists(area)
	log.Println("---")
	log.Println("---")

	maxDistanceArray := calculateMaxDistances(area)

	// not the proper way!!!!
	locations := findAppropriateLocations(maxDistanceArray)

	result, _ := json.Marshal(locations)

	return context.String(http.StatusOK, string(result))

}

func findAppropriateLocations(maxDistanceArray []int) []int {
	var locations []int
	minDistance := calculateMinDistance(maxDistanceArray)

	for i, element := range maxDistanceArray {
		if element == minDistance {
			locations = append(locations, i+1)
		}
	}
	return locations
}

func printRequiredLists(area domain.Area) {
	if area.Gym {
		log.Println("gym    ", domain.DistAreas["gym"])
	}
	if area.School {
		log.Println("school ", domain.DistAreas["school"])
	}
	if area.Store {
		log.Println("store  ", domain.DistAreas["store"])
	}
	if area.Lab {
		log.Println("lab    ", domain.DistAreas["lab"])
	}
	if area.Office {
		log.Println("office ", domain.DistAreas["office"])
	}
}

func calculateMaxDistances(area domain.Area) []int {
	var maxDistanceArray []int
	for i := 0; i < len(domain.DistAreas); i++ {
		maxDistanceArray = append(maxDistanceArray, 0)
		// log.Println(key, element)

		if area.Gym {
			// log.Println(domain.DistAreas["gym"][i])
			if maxDistanceArray[i] < domain.DistAreas["gym"][i] {
				maxDistanceArray[i] = domain.DistAreas["gym"][i]
			}
		}
		if area.School {
			// log.Println(domain.DistAreas["school"][i])
			if maxDistanceArray[i] < domain.DistAreas["school"][i] {
				maxDistanceArray[i] = domain.DistAreas["school"][i]
			}

		}
		if area.Store {
			// log.Println(domain.DistAreas["store"][i])
			if maxDistanceArray[i] < domain.DistAreas["store"][i] {
				maxDistanceArray[i] = domain.DistAreas["store"][i]
			}
		}
		if area.Lab {
			// log.Println(domain.DistAreas["lab"][i])
			if maxDistanceArray[i] < domain.DistAreas["lab"][i] {
				maxDistanceArray[i] = domain.DistAreas["lab"][i]
			}
		}
		if area.Office {
			// log.Println(domain.DistAreas["office"][i])
			if maxDistanceArray[i] < domain.DistAreas["office"][i] {
				maxDistanceArray[i] = domain.DistAreas["office"][i]
			}

		}
		// log.Println("---")
		// log.Println("---")
	}
	log.Println(maxDistanceArray)
	return maxDistanceArray
}

func calculateMinDistance(maxDistanceArray []int) int {
	var minDistance int
	for i, element := range maxDistanceArray {
		if i == 0 {
			minDistance = element
		} else {
			if minDistance > element {
				minDistance = element
			}
		}
	}

	log.Println(minDistance)
	return minDistance
}
