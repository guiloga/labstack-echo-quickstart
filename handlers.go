package echo_quickstart

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/labstack/echo/v4"
)

// StaticPath holds the path of served static files
var StaticPath string = GetStaticDir()

// GetStaticDir resolves the static files path
func GetStaticDir() string {
	curDir, err := os.Getwd()
	if err != nil {
		return "/static"
	}
	staticPath := path.Dir(curDir) + "/static"
	return staticPath
}

func ListUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, *MockedUsersList)
}

// CreateUser [User: CREATE]
func CreateUser(c echo.Context) error {
	// Bind json, xml, form or query payload into Go struct based on Content-Type request header
	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	r := c.Request()
	if r.Header.Get("Accept") == "application/xml" {
		return c.XML(http.StatusCreated, u)
	}

	return c.JSON(http.StatusCreated, u)
}

// RetrieveUser [User: RETRIEVE]
func RetrieveUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// UpdateUser [User: UPDATE]
func UpdateUser(c echo.Context) error {
	// User ID from path `users/:id`
	return c.String(http.StatusOK, "Updated")
}

// DeleteUser [User: DELETE]
func DeleteUser(c echo.Context) error {
	// User ID from path `users/:id`
	return c.String(http.StatusOK, "Deleted")
}

// ShowUser [Other: show User]
func ShowUser(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, fmt.Sprintf("team: %v member: %v", team, member))
}

// SaveUserOld [Other: save User (old)]
func SaveUserOld(c echo.Context) error {
	// Get team and member from form-data
	team := c.FormValue("team")
	member := c.FormValue("member")
	valid := true
	if team == "" || member == "" {
		valid = false
	}
	if !valid {
		return c.String(http.StatusBadRequest,
			"'team' and 'member' field values are required.")
	}
	return c.String(http.StatusOK,
		fmt.Sprintf("team: %v member: %v", team, member))
}

// SaveUser [Other: save User]
func SaveUser(c echo.Context) error {
	// Get name
	name := c.FormValue("name")

	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(StaticPath + "/" + avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")

}
