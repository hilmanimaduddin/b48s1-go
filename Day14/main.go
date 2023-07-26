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
	Checkbox1   string
	Checkbox2   string
	Checkbox3   string
	Checkbox4   string
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
	e.GET("/update/:id", update)

	e.POST("/add-blog", addBlog)
	e.POST("/blog-delete/:id", deleteBlog)
	e.POST("/update-blog", editBlog)

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

		err := data.Scan(&each.ID, &each.Subject, &each.StartDate, &each.EndDate, &each.Description, &each.Checkbox1, &each.Checkbox2, &each.Checkbox3, &each.Checkbox4, &each.Duration)
		
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.Image = ""
		each.Icon1 = MyIcon(each.Checkbox1)
		each.Icon2 = MyIcon(each.Checkbox2)
		each.Icon3 = MyIcon(each.Checkbox3)
		each.Icon4 = MyIcon(each.Checkbox4)
		each.Myicon1 = MyLabel(each.Checkbox1)
		each.Myicon2 = MyLabel(each.Checkbox2)
		each.Myicon3 = MyLabel(each.Checkbox3)
		each.Myicon4 = MyLabel(each.Checkbox4)

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

	idToInt, _ := strconv.Atoi(id)

	tmpl, err := template.ParseFiles("views/detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	BlogDetail := Blog{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT id, name, description, checkbox1, checkbox2, checkbox3, checkbox4, startdate, enddate, duration FROM tb_projects WHERE id=$1", idToInt).Scan(&BlogDetail.ID, &BlogDetail.Subject, &BlogDetail.Description, &BlogDetail.Checkbox1, &BlogDetail.Checkbox2, &BlogDetail.Checkbox3, &BlogDetail.Checkbox4, &BlogDetail.StartDate, &BlogDetail.EndDate, &BlogDetail.Duration)

	BlogDetail.Icon1 = MyIcon(BlogDetail.Checkbox1)
	BlogDetail.Icon2 = MyIcon(BlogDetail.Checkbox2)
	BlogDetail.Icon3 = MyIcon(BlogDetail.Checkbox3)
	BlogDetail.Icon4 = MyIcon(BlogDetail.Checkbox4)
	BlogDetail.Myicon1 = MyLabel(BlogDetail.Checkbox1)
	BlogDetail.Myicon2 = MyLabel(BlogDetail.Checkbox2)
	BlogDetail.Myicon3 = MyLabel(BlogDetail.Checkbox3)
	BlogDetail.Myicon4 = MyLabel(BlogDetail.Checkbox4)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}


	data := map[string]interface{}{
		"Id":   id,
		"Blog": BlogDetail,
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
	description := c.FormValue("input-description")
	duration := getDuration(endDate, startDate)


	// newBlog := Blog{
	// 	Subject:     project,
	// 	StartDate:   startDate,
	// 	EndDate:     endDate,
	// 	Duration:    getDuration(endDate, startDate),
	// 	Description: description,
	// 	Icon1:       icon1,
	// 	Icon2:       icon2,
	// 	Icon3:       icon3,
	// 	Icon4:       icon4,
	// 	Myicon1:     label1,
	// 	Myicon2:     label2,
	// 	Myicon3:     label3,
	// 	Myicon4:     label4,
	// }


	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_projects (name, description, checkbox1, checkbox2, checkbox3, checkbox4, startdate, enddate, duration) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", project, description, iconA, iconB, iconC, iconD, startDate, endDate, duration)


	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

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


func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// fmt.Println("index : ", id)

	// dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	connection.Conn.Exec(context.Background(), "DELETE from tb_projects WHERE id=$1", id)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func update(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	tmpl, err := template.ParseFiles("views/update.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	Update := Blog{}

	errUpdate := connection.Conn.QueryRow(context.Background(), "SELECT id, name, description, checkbox1, checkbox2, checkbox3, checkbox4, startdate, enddate FROM tb_projects WHERE id=$1", idToInt).Scan(&Update.ID, &Update.Subject, &Update.Description, &Update.Checkbox1, &Update.Checkbox2, &Update.Checkbox3, &Update.Checkbox4, &Update.StartDate, &Update.EndDate)

	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"Id":   id,
		"Blog": Update,
	}

	return tmpl.Execute(c.Response(), data)
}

func editBlog(c echo.Context) error {
	id := c.FormValue("id")
	project := c.FormValue("input-project")
	startDate := c.FormValue("input-startdate")
	endDate := c.FormValue("input-enddate")
	iconA := c.FormValue("checkbox1")
	iconB := c.FormValue("checkbox2")
	iconC := c.FormValue("checkbox3")
	iconD := c.FormValue("checkbox4")
	description := c.FormValue("input-description")
	duration := getDuration(endDate, startDate)

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_projects SET name=$1, description=$2, checkbox1=$3, checkbox2=$4, checkbox3=$5, checkbox4=$6, startdate=$7, enddate=$8, duration=$9 WHERE id=$10", project, description, iconA, iconB, iconC, iconD, startDate, endDate, duration, id)


	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}