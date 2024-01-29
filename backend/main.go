package main

import (
	"e-combomb/bootstrap"
	"e-combomb/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var (
	store *sessions.CookieStore
)

func init() {
	store = sessions.NewCookieStore([]byte("some-secret-key"))
}

func main() {
	router := gin.New()

	// Enable CORS for all origins
	router.Use(cors.Default())

	apiRoutes := router.Group("/api")

	app := bootstrap.App()
	env := app.Env
	gin.SetMode(env.GinMode)

	routes.SetupRoutes(apiRoutes, store, &app)

	router.Run()
}
