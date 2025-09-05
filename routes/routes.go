package routes

import (
	"fmt"
	"go-webserver/routes/api"
)

func RegisterMainRoutes() {
	fmt.Println("Registering main routes...")
	RegisterIndex()
	RegisterPage()
}

func RegisterApiRoutes() {
	fmt.Println("Registering API routes...")
	api.RegisterTest()
}