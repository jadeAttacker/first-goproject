package main

import (
	// "database/sql"
	// "fmt"
	// "log"
	// "net/http"

	// "github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
	"go.study/world"
)

// func response() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Next()
// 		if c.Writer.Written() {
// 			return
// 		}

// 		params := c.Keys
// 		if len(params) == 0 {
// 			return
// 		}
// 		c.JSON(http.StatusOK, params)
// 	}
// }

func main() {

	// db, err := sql.Open("mysql", "root:2019NULIfendou@tcp(127.0.0.1:3306)/day01?parseTime=true")
	// defer db.Close()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// db.SetMaxIdleConns(20)
	// db.SetMaxOpenConns(20)

	// if err := db.Ping(); err != nil {
	// 	log.Fatalln(err)
	// }
	// 同包的函数
	world.PkgFunc()
	// router := gin.Default()
	// // router.Use(response())
	// router.GET("/", func(c *gin.Context) {
	// 	name := c.Query("name")
	// 	score := c.Query("score")
	// 	rs, err := db.Exec("INSERT INTO first(name, score) VALUES (?, ?)", name, score)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	id, err := rs.LastInsertId()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Println("insert person Id {}", id)
	// 	c.String(http.StatusOK, "Hello World"+","+name+","+score)
	// 	fmt.Println("hello world============")
	// })

	// router.Run(":8000")
}
