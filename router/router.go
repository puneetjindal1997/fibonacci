package router

import (
	"fibonacci/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
}

type routes struct {
	router *gin.Engine
}
type Routes []Route

var FibonacciRoutes = Routes{
	//routes Fibonacci
	Route{"FibonacciPrint", "GET", "/", controllers.FibonacciPrint},
}

/*
 *	Function for grouping log routes
 */
func (r routes) FibonacciRoutesMethod(rg *gin.RouterGroup) {
	orderRouteGrouping := rg.Group("/")
	for _, route := range FibonacciRoutes {
		switch route.Method {
		case "GET":
			orderRouteGrouping.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			orderRouteGrouping.POST(route.Pattern, route.HandlerFunc)
		case "OPTIONS":
			orderRouteGrouping.OPTIONS(route.Pattern, route.HandlerFunc)
		case "PUT":
			orderRouteGrouping.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			orderRouteGrouping.DELETE(route.Pattern, route.HandlerFunc)
		default:
			orderRouteGrouping.GET(route.Pattern, func(c *gin.Context) {
				c.JSON(200, gin.H{
					"result": "Specify a valid http method with this route.",
				})
			})
		}
	}
}

// append routes with versions
func ClientRoutes() {
	r := routes{
		router: gin.Default(),
	}
	v1 := r.router.Group(os.Getenv("API_VERSION"))
	v1.Use(CORSMiddleware())
	r.FibonacciRoutesMethod(v1)
	if err := r.router.Run(":" + os.Getenv("CLIENT_PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// Middlewares
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
		}
	}
}
