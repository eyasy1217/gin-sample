package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", gin.H{"e": false})
	})

	router.POST("/login", func(ctx *gin.Context) {
		// form取得
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")

		// sqlite3データベースの読み込み
		db, _ := sql.Open("sqlite3", "sample.db")

		// 問い合わせ結果を入れる箱を用意
		var fullname string

		// データベース問い合わせ
		row := db.QueryRow("SELECT fullname FROM sample_user WHERE email = ? and password = ?;", email, password)

		defer db.Close()

		err := row.Scan(&fullname)

		if err == nil {
			ctx.HTML(200, "home.html", gin.H{"fullname": fullname})
		} else {
			// データがなかったら
			if err == sql.ErrNoRows {
				ctx.HTML(200, "login.html", gin.H{"e": true})
				// それ以外のエラー
			} else {
				log.Println(err)
			}
		}
	})

	router.Run()
}
