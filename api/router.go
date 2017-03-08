package api

import (
	"gin-template/config"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Router defined api namespace router initialize
func (api *API) Router() *gin.Engine {
	// initialize gin config
	appSecret := config.Get("app", "secret").Value()
	appName := config.Get("app", "name").Value()
	sessionStore := sessions.NewCookieStore([]byte(appSecret))

	// defined router
	router := gin.Default()
	router.Use(sessions.Sessions("_"+appName+"_session", sessionStore))
	router.GET("/", api.HomeIndex)
	router.NoRoute(HandleNoRoute)
	router.Use(HandleRecovery())
	return router
}