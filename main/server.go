package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	eq "echo_quickstart"
)

// AdminBasicAuthValidator validates admin user authentication
func adminBasicAuthValidator(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "adminSecret" {
		return true, nil
	}
	return false, nil
}

func logUserRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uri := c.Request().URL.RequestURI()
		fmt.Println(fmt.Sprintf(
			"[info][middleware pre-processing request] User request to uri: %v", uri))
		err := next(c)
		fmt.Println(fmt.Sprintf(
			"[info][middleware post-processing request] User request to uri: %v", uri))
		return err
	}
}

func main() {
	e := echo.New()

	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Admin Group Routes (with Group level middleware)
	adminGroup := e.Group("admin")
	adminGroup.Use(middleware.BasicAuth(adminBasicAuthValidator))
	// Users List
	adminGroup.GET("/users-list", eq.ListUsers)
	// Statis files
	adminGroup.Static("/static", eq.GetStaticDir())

	// User Routes (with Route level middleware)
	e.POST("/users", eq.CreateUser, logUserRequest)
	e.GET("/users/:id", eq.RetrieveUser, logUserRequest)
	e.PUT("/users/:id", eq.UpdateUser, logUserRequest)
	e.DELETE("/users/:id", eq.DeleteUser, logUserRequest)
	/*
		applying the logUserRequest middlewareFunc is equivalent to do
		CreateUserMiddleware := logUserRequest(CreateUser)
		e.POST("/users", CreateUserMiddleware)
	*/

	// Other Routes
	e.GET("/show", eq.ShowUser)
	e.POST("/save", eq.SaveUser)
	e.POST("/save-old", eq.SaveUserOld)

	e.Logger.Fatal(e.Start(":1323"))
}
