package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
	connection, err := sqlx.Connect("mysql", "sealteam:sckshuhari@(store-tearup:3306)/toy")
	if err != nil {
		log.Fatalln("cannot connect to tearup", err)
	}

	route := gin.Default()
	route.GET("/api/v1/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": GetUserNameFromDB(connection),
		})
	})
	log.Fatal(route.Run(":8000"))
}

func GetUserNameFromDB(connection *sqlx.DB) User {
	user := User{}
	err := connection.Get(&user, "SELECT id,name FROM user WHERE ID=1")
	if err != nil {
		fmt.Printf("Get user name from tearup get error : %s", err.Error())
		return User{}
	}
	return user
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
