package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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
	// URL Validation
	url := c.FormValue("url")
	url = strings.Replace(url, " ", "", -1)
	if url == "" {
		return jsonResponse(c, http.StatusBadRequest, "URL is required")
	}

	// Check if binaries exist
	if _, err := os.Stat("binaries"); os.IsNotExist(err) {
		return jsonResponse(c, http.StatusInternalServerError, "Binaries not found.")
	}

	executable := "clear-badge-cache"
	if strings.Contains(strings.ToLower(os.Getenv("GOOS")), "windows") {
		executable += ".exe"
	} else {
		executable += ".sh"
	}
	cmd := exec.Command(executable, url)
	cmd.Dir = "binaries"

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()
	if err != nil {
		return jsonResponse(c, http.StatusInternalServerError, fmt.Sprintf("Something went wrong, check your URL.<br>%v<br>%v", errBuf.String(), err.Error()))
	}

	return jsonResponse(c, http.StatusOK, fmt.Sprintf("Image cache for `%s` cleared successfully!", c.FormValue("url")))
}

func jsonResponse(c echo.Context, status int, message string) error {
	var resp Response
	if status < 300 {
		resp = Response{
			Status:  "Success",
			Message: message,
		}
	} else {
		resp = Response{
			Status:  "Error",
			Message: message,
		}
	}
	return c.JSON(200, resp)
}
