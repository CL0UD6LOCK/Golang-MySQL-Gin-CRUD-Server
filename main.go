package main

import (
	"database/sql"
	"fmt"
	"studentmanager/handlers"
	"studentmanager/storage"

	"github.com/joho/godotenv"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //只注册驱动
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system environment")
	}
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN environment variable is not set")
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	h := &handlers.StudentHandler{Store: storage.NewMySQLStore(db)}
	router := gin.Default()
	router.GET("/students", h.ListStudents)
	router.GET("/students/:id", h.GetStudent)
	router.POST("/students", h.CreateStudent)
	router.DELETE("/students/:id", h.DeleteStudent)
	router.PUT("/students/:id", h.UpdateStudent)
	router.PATCH("/students/:id", h.UpdateStudent)

	fmt.Println("Starting server at http://localhost:8080/students (Ctrl+C to stop)")
	err = router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
