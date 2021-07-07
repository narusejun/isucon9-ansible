package main

import (
	"echo/handler"
	"echo/interceptor"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	cache map[string]int
)

func init() {
	cache = map[string]int{}
}

func whenEver(c echo.Context) error {
	whenever := c.Param("whenever")

	jsonMap := map[string]string{
		"whenever": whenever,
	}

	return c.JSON(http.StatusOK, jsonMap)
}

func postSomething(c echo.Context) error {
	key := c.FormValue("key")
	value := c.FormValue("value")

	valueInt64, err := strconv.Atoi(value)

	if err != nil {
		log.Println(err)
		return echo.ErrForbidden
	}

	cache[key] = valueInt64

	return c.NoContent(204)
}

func showSomething(c echo.Context) error {
	jsonMap := map[string]string{}

	for key, value := range cache {
		jsonMap[key] = strconv.Itoa(value)
	}

	return c.JSON(http.StatusOK, jsonMap)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.Use(interceptor.BasicAuth())

	e.GET("/hello", handler.MainPage(), interceptor.BasicAuth())
	e.GET("/hello/:username", handler.SubPage())

	e.GET("/json/what/:whatever", handler.JsonPage())
	e.GET("/json/when/:whenever", whenEver)

	e.POST("/something", postSomething)
	e.GET("/show", showSomething)

	e.Start(":1323")
}
