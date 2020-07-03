package router

import (
	"net/http"

	"xorm.io/xorm"

	"log"

	"test.com/dzakwanzaky/golangstudy/pkg/types/routes"
	StatusHandler "test.com/dzakwanzaky/golangstudy/src/controllers/v1/status"
)

var db *xorm.Engine

func Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")
		if len(token) < 1 {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		log.Println("Inside V1 Middleware")

		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) (SubRoute map[string]routes.SubRoutePackage) {

	StatusHandler.Init(db)

	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": routes.SubRoutePackage{
			Routes: routes.Routes{
				routes.Route{"Status", "GET", "/status", StatusHandler.Index},
			},
			Middleware: Middleware,
		},
	}
	return
}
