package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"personal-web/connection"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Blog struct {
	ID			int
	Subject     string
	StartDate   string
	EndDate     string
	Month       string
	Description string
	Duration    string
	checkbox1   string
	checkbox2   string
	checkbox3   string
	checkbox4   string
	Icon1       string
	Icon2       string
	Icon3       string
	Icon4       string
	Myicon1     string
	Myicon2     string
	Myicon3     string
	Myicon4     string
	Image		string
}

var dataBlog = []Blog{
	{
		Subject:     "Kucing Lucu",
		StartDate:   "2023-03-17",
		EndDate:     "2023-04-18",
		Duration:    "2 weeks",
		Description: "Alangkah Indahnya Hari ini",
		Icon1:       `fa-brands fa-node-js fa-xl me-2`,
		Icon2:       `fa-brands fa-react fa-xl me-2`,
		Icon3:       `fa-brands fa-jsfiddle fa-xl me-2`,
		Icon4:       `fa-brands fa-java fa-xl me-2`,
		Myicon1:     `Node Js`,
		Myicon2:     `Next Js`,
		Myicon3:     `React Js`,
		Myicon4:     `TypeScript`,
	},
	{
		Subject:     "Kucing Comel",
		StartDate:   "2023-06-04",
		EndDate:     "2023-08-01",
		Duration:    "2 months",
		Description: "Makan Dulu aja... Lagi laper,,",
		Icon1:       `fa-brands fa-node-js fa-xl me-2`,
		Icon2:       `fa-brands fa-react fa-xl me-2`,
		Icon3:       `fa-brands fa-jsfiddle fa-xl me-2`,
		Icon4:       `fa-brands fa-java fa-xl me-2`,
		Myicon1:     `Node Js`,
		Myicon2:     `Next Js`,
		Myicon3:     `React Js`,
		Myicon4:     `TypeScript`,
	},
}

func main() {
	e := echo.New()

	connection.DatabaseConnect()

	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/project", project)
	e.GET("/testimonial", testimonial)
	e.GET("/blog", blog)
	e.GET("/detail/:id", detail)

	e.POST("/add-blog", addBlog)
	e.POST("/blog-delete/:id", deleteBlog)
	e.POST("/update-blog/:id", editBlog)

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

func blog(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, name, startdate, enddate, description, checkbox1, checkbox2, checkbox3, checkbox4, duration FROM tb_projects")

	var result []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Subject, &each.StartDate, &each.EndDate, &each.Description, &each.checkbox1, &each.checkbox2, &each.checkbox3, &each.checkbox4, &each.Duration)
		
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.Image = ""
		each.Icon1 = MyIcon(each.checkbox1)
		each.Icon2 = MyIcon(each.checkbox2)
		each.Icon3 = MyIcon(each.checkbox3)
		each.Icon4 = MyIcon(each.checkbox4)
		each.Myicon1 = MyLabel(each.checkbox1)
		each.Myicon2 = MyLabel(each.checkbox2)
		each.Myicon3 = MyLabel(each.checkbox3)
		each.Myicon4 = MyLabel(each.checkbox4)

		result = append(result, each)
	}
	
	tmpl, err := template.ParseFiles("views/blog.html")
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	blogs := map[string]interface{}{
		"Blogs": result,
	}

	return tmpl.Execute(c.Response(), blogs)
}

func detail(c echo.Context) error {
	id := c.Param("id") 

	tmpl, err := template.ParseFiles("views/detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	blogDetail := Blog{}

	for index, data := range dataBlog {
		if index == idToInt {
			blogDetail = Blog{
				Subject:     data.Subject,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Duration:    data.Duration,
				Description: data.Description,
				Icon1:       data.Icon1,
				Icon2:       data.Icon2,
				Icon3:       data.Icon3,
				Icon4:       data.Icon4,
				Myicon1:     data.Myicon1,
				Myicon2:     data.Myicon2,
				Myicon3:     data.Myicon3,
				Myicon4:     data.Myicon4,
			}
		}
	}

	data := map[string]interface{}{
		"Id":   id,
		"Blog": blogDetail,
	}

	return tmpl.Execute(c.Response(), data)
}

func MyIcon(Valu string) string {
	if Valu == "checkbox1" {
		return `fa-brands fa-node-js fa-xl me-2`
	} else if Valu == "checkbox2" {
		return `fa-brands fa-react fa-xl me-2`
	} else if Valu == "checkbox3" {
		return `fa-brands fa-jsfiddle fa-xl me-2`
	} else if Valu == "checkbox4" {
		return `fa-brands fa-java fa-xl me-2`
	} else {
		return ""
	}
}

func MyLabel(Valu string) string {
	if Valu == "checkbox1" {
		return `Node Js`
	} else if Valu == "checkbox2" {
		return `Next Js`
	} else if Valu == "checkbox3" {
		return `React Js`
	} else if Valu == "checkbox4" {
		return `TypeScript`
	} else {
		return ""
	}
}

func addBlog(c echo.Context) error {
	project := c.FormValue("input-project")
	startDate := c.FormValue("input-startdate")
	endDate := c.FormValue("input-enddate")
	iconA := c.FormValue("checkbox1")
	iconB := c.FormValue("checkbox2")
	iconC := c.FormValue("checkbox3")
	iconD := c.FormValue("checkbox4")
	icon1 := MyIcon(iconA)
	icon2 := MyIcon(iconB)
	icon3 := MyIcon(iconC)
	icon4 := MyIcon(iconD)
	label1 := MyLabel(iconA)
	label2 := MyLabel(iconB)
	label3 := MyLabel(iconC)
	label4 := MyLabel(iconD)
	description := c.FormValue("input-description")

	newBlog := Blog{
		Subject:     project,
		StartDate:   startDate,
		EndDate:     endDate,
		Duration:    getDuration(endDate, startDate),
		Description: description,
		Icon1:       icon1,
		Icon2:       icon2,
		Icon3:       icon3,
		Icon4:       icon4,
		Myicon1:     label1,
		Myicon2:     label2,
		Myicon3:     label3,
		Myicon4:     label4,
	}


	dataBlog = append(dataBlog, newBlog)

	fmt.Println(dataBlog)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func getDuration(endDate, startDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " year"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " months"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " month"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " weeks"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " week"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " days"
				} else {
					duration = strconv.Itoa(durationDays) + " day"
				}
			}
		}
	}

	return duration
}

func editBlog(edit echo.Context) error {
	id, _ := strconv.Atoi(edit.Param("id"))
	fmt.Println("index : ", id)

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)
	return edit.Redirect(http.StatusMovedPermanently, "/project")
}


func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("index : ", id)

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}