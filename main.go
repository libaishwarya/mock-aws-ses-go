package main

import (
	router "github.com/libaishwarya/mock-aws-ses-go/internal/router"
)

func main() {
	r := router.NewRouter()

	r.Run(":8080")
}
