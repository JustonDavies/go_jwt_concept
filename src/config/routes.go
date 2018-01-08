//-- Package Declaration -----------------------------------------------------------------------------------------------
package config

//-- Imports -----------------------------------------------------------------------------------------------------------
import (
	"github.com/labstack/echo"

	"github.com/JustonDavies/go_jwt_concept/src/controllers/authorization"
	"github.com/JustonDavies/go_jwt_concept/src/controllers/examples"
)

//-- Constants ---------------------------------------------------------------------------------------------------------
var (
	routes = []*route{

		//-- Examples ----------
		{`GET`, `/`, examples.Get, nil},
		{`GET`, `/restricted`, examples.Get, []echo.MiddlewareFunc{authorization.AuthenticationMiddleware}},

		//-- Sessions ----------
		{`POST`, `/auth/login`, authorization.Create, nil},
		{`POST`, `/auth/renew`, authorization.Renew, []echo.MiddlewareFunc{authorization.AuthenticationMiddleware}},
	}
)

//-- Structs -----------------------------------------------------------------------------------------------------------
type route struct {
	Verb       string
	Path       string
	Handler    echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
}

//-- Exported Functions ------------------------------------------------------------------------------------------------
func InjectRoutes(server *echo.Echo) {
	for _, route := range routes {
		server.Add(route.Verb, route.Path, route.Handler, route.Middleware...)
	}
}

//-- Internal Functions ------------------------------------------------------------------------------------------------
