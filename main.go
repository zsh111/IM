package main

import (
	"IMsystem/router"
	"IMsystem/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	
	r := router.Router()
	r.Run(":8081")
}