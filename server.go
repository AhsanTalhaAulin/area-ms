package main

import (
	"github.com/labstack/echo"

	"area-ms/conn"
	"area-ms/svc"
)

// main function
func main() {

	// domain.PrintDbPass()
	conn.ConnectDB()
	e := echo.New()

	e.POST("/area", svc.SaveArea)
	e.GET("/area", svc.GetArea)
	e.GET("/test", svc.UpdateDistance)

	e.Logger.Fatal(e.Start(":8080"))

}
