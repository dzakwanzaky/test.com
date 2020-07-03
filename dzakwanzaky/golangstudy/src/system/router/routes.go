package router

import (
	"net/http"

	"log"

	"test.com/dzakwanzaky/golangstudy/pkg/types/routes"
	AuthHandler "test.com/dzakwanzaky/golangstudy/src/controllers/auth"
	HomeHandler "test.com/dzakwanzaky/golangstudy/src/controllers/home"
	"xorm.io/xorm"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Global Middleware Reached!")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) routes.Routes {

	AuthHandler.Init(db)
	HomeHandler.Init(db)

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
		routes.Route{"AuthStore", "POST", "/auth/login", AuthHandler.Login},
		routes.Route{"AuthCheck", "GET", "/auth/check", AuthHandler.Check},
	}
}
