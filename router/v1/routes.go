package routesv1

import (
	"net/http"

	user "github.com/giftedunicorn/kusalt/kusalt/v1/user"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Auth        bool
	Verify      bool
}

type RouteList []Route

var Routes = RouteList{
	Route{
		"HealthCheck",
		"GET",
		"/api/v1/healthcheck",
		handlers.HealthCheck,
		false,
		false,
	},
	// ACTION
	Route{
		"Signin",
		"POST",
		"/api/v1/actions",
		handlers.CreateAction,
		false,
		true,
	},
	// AUTH
	Route{
		"Signin",
		"POST",
		"/api/v1/auth/signin",
		auth.Signin,
		false,
		true,
	},
	Route{
		"Signin",
		"POST",
		"/api/v1/auth/signin/mfa",
		auth.SigninMfa,
		false,
		true,
	},
	Route{
		"Signin",
		"POST",
		"/api/v1/auth/authorize",
		auth.CheckLoginUser,
		true,
		true,
	},
	Route{
		"Signin",
		"POST",
		"/api/v1/auth/token",
		auth.GetIdToken,
		false,
		true,
	},
	Route{
		"Signin",
		"POST",
		"/api/v1/auth/logout",
		auth.Signout,
		false,
		true,
	},
	Route{
		"SignupOne",
		"POST",
		"/api/v1/auth/signup/one",
		auth.SignupOne,
		false,
		true,
	},
	Route{
		"SignupTwo",
		"POST",
		"/api/v1/auth/signup/two",
		auth.SignupTwo,
		false,
		true,
	},
	// USERS
	Route{
		"GetUser",
		"POST",
		"/api/v1/user",
		user.GetUser,
		true,
		true,
	},
}
