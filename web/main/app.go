package main

import "web"

func main() {
	router := web.GetRouter()
	router.Run(":8080")
}