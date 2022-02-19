package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// User
type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}
type UserDTO struct {
	Name    string
	Email   string
	IsAdmin bool
}

func main() {
	e := echo.New()
	e.POST("/users", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
