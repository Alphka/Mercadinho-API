package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHeadRoutes mirrors every GET route as a HEAD route (Gin does not do
// it on its own), so uptime monitors that only send HEAD requests work on any
// endpoint. Routes that already have a HEAD handler are left untouched.
// It must be called after all other routes are registered.
func RegisterHeadRoutes(router *gin.Engine) {
	routes := router.Routes()

	registered := make(map[string]bool)
	for _, route := range routes {
		if route.Method == http.MethodHead {
			registered[route.Path] = true
		}
	}

	for _, route := range routes {
		if route.Method != http.MethodGet || registered[route.Path] {
			continue
		}

		handler := route.HandlerFunc

		router.HEAD(route.Path, func(context *gin.Context) {
			// Some handlers (e.g. Swagger UI) reject anything but GET, so they
			// get a GET copy of the request. The original request must stay
			// untouched: net/http checks it to discard the body of HEAD responses.
			getRequest := *context.Request
			getRequest.Method = http.MethodGet
			context.Request = &getRequest

			handler(context)
		})
	}
}
