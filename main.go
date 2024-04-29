package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func staticFunc(url string) string {
	return "/static/" + url
}

func newTemplate() *Templates {
	funcMap := template.FuncMap{
		"static": staticFunc,
	}

	tmpls, err := template.New("").Funcs(funcMap).ParseGlob("views/*.html")
	if err != nil {
		panic(err)
	}
	return &Templates{
		templates: template.Must(tmpls, err),
	}
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()

	// Serve static files
	fs := http.FileServer(http.Dir("./static/"))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.POST("/clear-cache", clearCache)

	e.Logger.Fatal(e.Start(":8080"))
}

func clearCache(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("Image cache for `%s` cleared successfully!", c.FormValue("url")))
}
