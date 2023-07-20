package main

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/project", project)
	e.GET("/testimonial", testimonial)
	e.GET("/detail/:id", detail)

	e.POST("/add-blog", addBlog)

	e.GET("/coba", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func project(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func testimonial(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/testimonial.html")
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return tmpl.Execute(c.Response(), nil)
}


func detail(c echo.Context) error {
	id := c.Param("id")

	tmpl, err := template.ParseFiles("views/detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	blogDetail := map[string]interface{}{
		"Id":			id,
		"Project":		"Suasana Hari ini",
		"StartDate":	"20 07 2023",
		"EndDate":		"30 07 2023",
		"Description":	"Suasananya sejuk enak ",
	}

	return tmpl.Execute(c.Response(), blogDetail)
}

func addBlog(c echo.Context) error {
	project := c.FormValue("input-project")
	startDate := c.FormValue("input-startdate")
	endDate := c.FormValue("input-enddate")
	checkbox1 := c.FormValue("checkbox1")
	checkbox2 := c.FormValue("checkbox2")
	checkbox3 := c.FormValue("checkbox3")
	checkbox4 := c.FormValue("checkbox4")
	description := c.FormValue("input-description")

	println("Project : " + project)
	println("Start Date : " + startDate)
	println("End Date : " + endDate)
	println("Technologies : " + checkbox1 +" " + checkbox2 + " " + checkbox3 + " " + checkbox4)
	println("Description : " + description)

	return c.Redirect(http.StatusMovedPermanently, "/project")
}