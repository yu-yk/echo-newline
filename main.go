package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.JSON(http.StatusOK, User{"Jon Snow", "jon@labstack.com"})
	fmt.Printf("%#v", rec.Body.String())
}
