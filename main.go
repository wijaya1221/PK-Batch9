package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"

)

type BaseResponse struct {
	Status bool   		`json:"status"`
	Message string 		`json:"message"`
	Data interface{}	`json:"data"`
}

type News struct {
	Title string 		`json:"title"`
	Content string 		`json:"content"`
}

type User struct {
	Id int
	PhotoUrl string
	UserName string
	FullName string
}

func main(){
	initDatabase()
	e := echo.New()
	defer DB.Close()
	e.GET("/users", GetUsersController)
	e.GET("/news", GetNewsController)
	e.POST("/news", AddNewsController)
	e.GET("/news/:id", GetDetailNewsController)
	e.Start(":8000")
}

var DB *sql.DB

func initDatabase() {
	var err error
	DB, err = sql.Open("mysql", "root:123ABC4d.@tcp(localhost:3306)/prakerja9")
	if err != nil {
		panic("Gagal konek ke database")
	}
	// defer DB.Close()
}

func GetUsersController(c echo.Context) error {
	var users []User

	// database
	query := "SELECT * FROM user"
	rows, err := DB.Query(query)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.PhotoUrl, &user.UserName, &user.FullName); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: "Berhasil get data",
		Data: users,
	})
}

func AddNewsController(c echo.Context) error {
	var newsRequest News
	c.Bind(&newsRequest)

	return c.JSON(http.StatusCreated, BaseResponse{
		Status: true,
		Message: "success",
		Data: newsRequest,
	})
}

func GetDetailNewsController(c echo.Context) error {
	var id = c.Param("id")
	var data = News{"A", "A"}

	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: id,
		Data: data,
	})
}

func GetNewsController(c echo.Context) error {
	country := c.QueryParam("country")
	
	
	var data []News

	// select * from news
	// orm


	// dummy
	data = append(data, News{"A", "A"})
	data = append(data, News{"B", "B"})

	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: country,
		Data: data,
	})
}





