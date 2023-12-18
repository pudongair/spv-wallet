package xpubs

import (
	"github.com/BuxOrg/bux-server/actions"
	"github.com/BuxOrg/bux-server/config"
	apirouter "github.com/mrz1836/go-api-router"
)

// Action is an extension of actions.Action for this package
type Action struct {
	actions.Action
}

// RegisterRoutes register all the package specific routes
func RegisterRoutes(router *apirouter.Router, appConfig *config.AppConfig, services *config.AppServices) {

	// Use the authentication middleware wrapper
	a, requireAdmin := actions.NewStack(appConfig, services)
	requireAdmin.Use(a.RequireAdminAuthentication)

	// Use the authentication middleware wrapper - this will only check for a valid xPub
	aBasic, requireBasic := actions.NewStack(appConfig, services)
	requireBasic.Use(aBasic.RequireBasicAuthentication)

	// Load the actions and set the services
	action := &Action{actions.Action{AppConfig: a.AppConfig, Services: a.Services}}

	// V1 Requests
	router.HTTPRouter.GET("/"+config.APIVersion+"/xpub", action.Request(router, requireBasic.Wrap(action.get)))
	router.HTTPRouter.POST("/"+config.APIVersion+"/xpub", action.Request(router, requireAdmin.Wrap(action.create)))
	router.HTTPRouter.PATCH("/"+config.APIVersion+"/xpub", action.Request(router, requireAdmin.Wrap(action.update)))
}
