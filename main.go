package main

import (
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

type myForm struct {
	User_id string `form:"user_id"`
	Pass    string `form:"pass"`
}

func main() {
	r := gin.Default()

	// redis-server でサーバーを起動
	ci, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = ci.Do("FLUSHALL")
	if err != nil {
		fmt.Println("no")

	}
	defer ci.Close()
	//redisに接続

	r.LoadHTMLGlob("views/*")
	r.GET("/login", indexHandler)
	r.POST("/user", formHandler)

	r.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm)
	//fmt.Println(fakeForm.Pass)
	//fmt.Println(fakeForm.User_id)
	ci, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(1)
	}
	ci.Do("HSET", fakeForm.User_id, "pass", fakeForm.Pass)

	//fmt.Println(ri) // OK
	aii, err := redis.String(ci.Do("HGET", fakeForm.User_id, "pass"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(aii)

	c.JSON(200, gin.H{"text": fakeForm.User_id})

	// 値の読み出し

}
