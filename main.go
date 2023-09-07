package main

import "IMsystem/router"

func main() {
	r := router.Router()
	r.Run(":8081")
}