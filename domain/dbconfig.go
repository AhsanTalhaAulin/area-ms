package domain

import (
	"os"
)

var DbName = "area_db"
var DbUserName = "area_user"
var DbPass = os.Args[1]

// func PrintDbPass() {
// 	log.Println("DbPass : ", DbPass)
// }

var DbPort = "3308"

var DbMaxIdleConns = 10
var DbMaxOpenConns = 10
var DbConnMaxLifetime = 1

var DbTable = "areas"
